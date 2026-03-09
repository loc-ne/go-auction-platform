package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	"github.com/loc-ne/go-auction/services/product/internal/usecase"
	"github.com/loc-ne/go-auction/shared/middleware"
)

type ProductHandler struct {
	usecase     usecase.ProductUsecase
	statUsecase usecase.ProductStatUsecase
}

func NewProductHandler(r *gin.Engine, u usecase.ProductUsecase, statU usecase.ProductStatUsecase, jwtSecret string) {
	handler := &ProductHandler{
		usecase:     u,
		statUsecase: statU,
	}

	productGroup := r.Group("/api/v1/products")
	{
		productGroup.GET("/active", handler.GetActiveAuctions)
		productGroup.GET("/trending", handler.GetTrendingProducts)
		productGroup.GET("/:id", handler.GetProductByID)

		protected := productGroup.Group("")
		protected.Use(middleware.AuthMiddleware(jwtSecret))
		{
			protected.POST("", handler.CreateProduct)
			protected.POST("/:id/favorite", handler.HandleFavorite)
		}
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Ép kiểu userID thành UUID
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	product.SellerID = parsedUserID

	if err := h.usecase.Create(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "data": product})
}

func (h *ProductHandler) GetActiveAuctions(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	products, err := h.usecase.GetActiveAuctions(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active products: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")

	product, err := h.usecase.GetByID(c.Request.Context(), idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	h.statUsecase.QueueView(idStr)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandler) HandleFavorite(c *gin.Context) {
	idStr := c.Param("id")
	productID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := uuid.Parse(userIDRaw.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
		return
	}

	isFavorite, err := h.usecase.HandleFavorite(c.Request.Context(), userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle favorite action: " + err.Error()})
		return
	}

	go h.statUsecase.RefreshHotRankingByID(context.Background(), idStr)

	c.JSON(http.StatusOK, gin.H{"is_favorite": isFavorite})
}

func (h *ProductHandler) GetTrendingProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}

	products, err := h.usecase.GetTrendingProducts(c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trending products: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
