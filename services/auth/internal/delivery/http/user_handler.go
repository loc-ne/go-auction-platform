package http

import (
	"errors"
    "net/http"
    "github.com/gin-gonic/gin" 
    "github.com/loc-ne/go-auction/services/auth/internal/usecase"
)

type UserHandler struct {
    usecase usecase.UserUsecase 
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
    return &UserHandler{usecase: u}
}

func (h *UserHandler) Register(c *gin.Context) {
    var req struct {
		FullName     string    `json:"full_name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }
    ctx := c.Request.Context()
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data is invalid"})
        return
    }

    err := h.usecase.Register(ctx, req.FullName, req.Email, req.Password)
    if err != nil {
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
        c.JSON(http.StatusConflict, gin.H{"error": "Email has already been used"})
        return
    }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Register successfully!"})
}

func (h *UserHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }
    ctx := c.Request.Context()
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data is invalid"})
        return
    }

    user, err := h.usecase.Login(ctx, req.Email, req.Password)
    if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successfully!", "data": user})
}