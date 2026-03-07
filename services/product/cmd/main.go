package main

import (
    "context"
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/loc-ne/go-auction/services/product/migrations"
    "github.com/loc-ne/go-auction/services/product/internal/repository/postgres"
    "github.com/loc-ne/go-auction/services/product/internal/usecase"
    productHttp "github.com/loc-ne/go-auction/services/product/internal/delivery/http"
)   

func main() {
 
    ctx := context.Background()
    migrations.RunMigrations()

    db, err := postgres.NewPostgresDB(ctx)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Pool.Close()

    jwtSecret := os.Getenv("JWT_SECRET") 
    repo := postgres.NewProductRepository(db.Pool)
    productUsecase := usecase.NewProductUsecase(repo)
    router := gin.Default()
    productHttp.NewProductHandler(router, productUsecase, jwtSecret)
    port := os.Getenv("PRODUCT_PORT")
    log.Printf("Product Service listening on port %s...\n", port)
    router.Run(":" + port)
    
}