package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/loc-ne/go-auction/services/bidding/internal/usecase"
)

type WebsocketBid struct {
	usecase usecase.BidUsecase
	hub     *Hub
}

func NewWebsocketBid(usecase usecase.BidUsecase, h *Hub) *WebsocketBid {
	return &WebsocketBid{
		usecase: usecase,
		hub:     h,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func (wsb *WebsocketBid) ServeWs(c *gin.Context) {
	productID := c.Query("productId")
	
	ok, err := wsb.usecase.CheckRoomActive(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room is not active"})
		return
	}
	
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	userIDStr := fmt.Sprintf("%v", userIDVal)

	if productID == "" || userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing productId or userId"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		return
	}

	client := NewClient(wsb.hub, productID, userIDStr, conn)
	
	wsb.hub.Register(client)

	go client.WritePump()
	go client.ReadPump()
}
