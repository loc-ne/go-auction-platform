package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/bidding/internal/entity"
	"github.com/loc-ne/go-auction/services/bidding/internal/usecase"
)

type BidHandler struct {
	usecase usecase.BidUsecase
}

func NewBidHandler(router *gin.Engine, uc usecase.BidUsecase, authMiddleware gin.HandlerFunc) {
	handler := &BidHandler{
		usecase: uc,
	}

	bidGroup := router.Group("/api/v1/bids")
	bidGroup.Use(authMiddleware)
	{
		bidGroup.POST("", handler.CreateBid)
	}
}

func (h *BidHandler) CreateBid(c *gin.Context) {
	type CreateBidRequest struct {
		ProductID uuid.UUID `json:"product_id"`
		Amount    int64     `json:"amount"`
	}

	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
		return
	}

	var req CreateBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update later (race condition)
	if err := h.usecase.ValidateBid(c.Request.Context(), req.ProductID.String(), req.Amount, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bid := entity.Bid{
		ProductID: req.ProductID,
		UserID:    userID,
		Amount:    req.Amount,
	}
	
	fullNameVal, _ := c.Get("full_name")
	bidderName, _ := fullNameVal.(string)
	if bidderName == "" {
		bidderName = "Anonymous"
	}
	if err := h.usecase.CreateBid(c.Request.Context(), &bid, bidderName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bid created successfully", "data": bid})
}
