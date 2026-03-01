package http

import (
	"errors"
    "net/http"
    "fmt"
    "time"
    "github.com/gin-gonic/gin" 
    "github.com/loc-ne/go-auction/services/auth/internal/usecase"
	"github.com/loc-ne/go-auction/services/auth/internal/entity"
)

type UserHandler struct {
    userUsecase usecase.UserUsecase 
    sessionUsecase usecase.SessionUsecase
}

func NewUserHandler(u usecase.UserUsecase, s usecase.SessionUsecase) *UserHandler {
    return &UserHandler{userUsecase: u, sessionUsecase: s}
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

    err := h.userUsecase.Register(ctx, req.FullName, req.Email, req.Password)
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

    user, accessToken, refreshToken, err := h.userUsecase.Login(ctx, req.Email, req.Password)
    if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
        return
    }
    session := &entity.Session{
        UserID: user.ID,
        RefreshToken: refreshToken,
        UserAgent: c.Request.UserAgent(),
        ClientIP: c.ClientIP(),
        ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
    }   
    err = h.sessionUsecase.CreateSession(ctx, session)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
        return
    }

    c.SetCookie("access_token", accessToken, 60*60*24*7, "/", "", false, true)
    c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}