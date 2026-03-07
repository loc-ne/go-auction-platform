package main

import (
    "context"
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/loc-ne/go-auction/services/bidding/migrations"
    "github.com/loc-ne/go-auction/services/bidding/internal/repository/postgres"
)

func main() {
 
    ctx := context.Background()
    migrations.RunMigrations()

    db, err := postgres.NewPostgresDB(ctx)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Pool.Close()

    router := gin.Default()

    port := os.Getenv("BIDDING_PORT")
    log.Printf("Bidding Service listening on port %s...\n", port)
    router.Run(":" + port)
    
}