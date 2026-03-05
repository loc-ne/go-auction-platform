package main

import (
    "context"
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/loc-ne/go-auction/services/product/migrations"
    "github.com/loc-ne/go-auction/services/product/internal/repository/postgres"
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

    port := os.Getenv("PRODUCT_PORT")
    log.Printf("Product Service listening on port %s...\n", port)
    router.Run(":" + port)
    
}