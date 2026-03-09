package http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/loc-ne/go-auction/shared/pkg"
	"github.com/loc-ne/go-auction/shared/middleware"
)

type MediaHandler struct {
	cloudinaryUsecase pkg.CloudinaryUsecase
}

func NewMediaHandler(r *gin.Engine, cloudinaryUsecase pkg.CloudinaryUsecase, jwtSecret string) {
	handler := &MediaHandler{
		cloudinaryUsecase: cloudinaryUsecase,
	}

	mediaGroup := r.Group("/api/v1/media")
	mediaGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		mediaGroup.POST("/upload", handler.UploadMedia)
	}
}

func (h *MediaHandler) UploadMedia(c *gin.Context) {
	err := c.Request.ParseMultipartForm(20 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form: " + err.Error()})
		return
	}

	files := c.Request.MultipartForm.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files provided"})
		return
	}

	if a := len(files); a > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 5 files allowed"})
		return
	}

	var imageUrls []string
	var errs []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := range files {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			file := files[idx]
			url, err := h.cloudinaryUsecase.UploadImage(c.Request.Context(), file)

			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errs = append(errs, fmt.Sprintf("Failed to upload file %s: %v", file.Filename, err))
			} else {
				imageUrls = append(imageUrls, url)
			}
		}(i)
	}

	wg.Wait()

	if len(errs) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Some files failed to upload", "details": errs})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"data":    imageUrls,
	})
}
