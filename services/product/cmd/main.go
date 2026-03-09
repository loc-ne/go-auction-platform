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
    "github.com/loc-ne/go-auction/services/product/internal/repository/redis"
	"github.com/loc-ne/go-auction/services/product/internal/worker"
    "github.com/loc-ne/go-auction/shared/pkg"
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
	statRepo := postgres.NewProductStatRepository(db.Pool)

    redisClient, err := redis.NewRedisRepository()
	if err != nil {
		log.Fatal("Failed to connect to redis:", err)
	}
	defer redisClient.Close()

    productUsecase := usecase.NewProductUsecase(repo, redisClient)
	productStatUsecase := usecase.NewProductStatUsecase(statRepo, redisClient)

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudAPISecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudinaryUsecase := pkg.NewCloudinaryUsecase(cloudName, cloudAPIKey, cloudAPISecret)

	productStatWorker := worker.NewProductStatWorker(redisClient, productStatUsecase, productUsecase)
	go productStatWorker.Start(ctx)

    router := gin.Default()
	router.MaxMultipartMemory = 20 << 20 

    productHttp.NewProductHandler(router, productUsecase, productStatUsecase, jwtSecret)
	productHttp.NewMediaHandler(router, cloudinaryUsecase, jwtSecret)
    
	port := os.Getenv("PRODUCT_PORT")
    log.Printf("Product Service listening on port %s...\n", port)
    router.Run(":" + port)
    
}