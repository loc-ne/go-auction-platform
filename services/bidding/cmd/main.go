package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loc-ne/go-auction/services/bidding/internal/delivery/http"
	"github.com/loc-ne/go-auction/services/bidding/internal/repository/postgres"
	"github.com/loc-ne/go-auction/services/bidding/internal/usecase"
	"github.com/loc-ne/go-auction/services/bidding/migrations"
	"github.com/loc-ne/go-auction/shared/middleware"
)

func main() {

	ctx := context.Background()
	migrations.RunMigrations()

	db, err := postgres.NewPostgresDB(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Pool.Close()

	bidRepo := postgres.NewBidRepository(db.Pool)
	bidUsecase := usecase.NewBidUsecase(bidRepo)

	router := gin.Default()
	
	jwtSecret := os.Getenv("JWT_SECRET")
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	
	http.NewBidHandler(router, bidUsecase, authMiddleware)

	port := os.Getenv("BIDDING_PORT")

	log.Printf("Bidding Service listening on port %s...\n", port)
	router.Run(":" + port)

}