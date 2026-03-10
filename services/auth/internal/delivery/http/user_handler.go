package http

import (
	"errors"
    "net/http"
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

    c.SetCookie("access_token", accessToken, 15*60, "/", "", false, true)
    c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", false, true)
    resUser := gin.H{
        "email":     user.Email,
        "fullName":  user.FullName,
        "role":      user.Role,
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login successfully", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken, "user": resUser}})
}

func (h *UserHandler) Me(c *gin.Context) {
    email, exists := c.Get("email")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized"})
        return
    }

    ctx := c.Request.Context()
    user, err := h.userUsecase.GetUser(ctx, email.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to retrieve user information"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "user": gin.H{
            "email":     user.Email,
            "fullName":  user.FullName,
            "role":      user.Role,
        },
    })
}

func (h *UserHandler) Logout(c *gin.Context) {
    refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token not found"})
        return
    }

    ctx := c.Request.Context()
    
    session, err := h.sessionUsecase.GetSessionByRefreshToken(ctx, refreshToken)
    if err == nil && session != nil {
        _ = h.sessionUsecase.DeleteSession(ctx, session.ID)
    }

    c.SetCookie("access_token", "", -1, "/", "", false, true)
    c.SetCookie("refresh_token", "", -1, "/", "", false, true)

    c.JSON(http.StatusOK, gin.H{
        "success": true, 
        "message": "Logout successfully",
    })
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
    refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
         c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
         return
    }

    ctx := c.Request.Context()

    session, err := h.sessionUsecase.GetSessionByRefreshToken(ctx, refreshToken)
    if err != nil || session == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Session is invalid or expired"})
        return
    }

    user, err := h.userUsecase.GetUserByID(ctx, session.UserID)
    if err != nil || user == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
        return
    }

    newAccessToken, newRefreshToken, err := h.userUsecase.RefreshTokens(ctx, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new token"})
        return
    }

    _ = h.sessionUsecase.DeleteSession(ctx, session.ID)
    
    newSession := &entity.Session{
        UserID: user.ID,
        RefreshToken: newRefreshToken,
        UserAgent: c.Request.UserAgent(),
        ClientIP: c.ClientIP(),
        ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
    }   
    _ = h.sessionUsecase.CreateSession(ctx, newSession)

    c.SetCookie("access_token", newAccessToken, 15*60, "/", "", false, true)
    c.SetCookie("refresh_token", newRefreshToken, 60*60*24*30, "/", "", false, true)

    resUser := gin.H{
        "email":     user.Email,
        "fullName":  user.FullName,
        "role":      user.Role,
    }
    c.JSON(http.StatusOK, gin.H{
        "success": true, 
        "data": gin.H{
             "access_token": newAccessToken, 
             "refresh_token": newRefreshToken, 
             "user": resUser,
        },
    })
}