package main

import (
    "context"
    "log"
    "github.com/loc-ne/go-auction/services/auth/migrations" 
    "github.com/loc-ne/go-auction/services/auth/internal/repository/postgres" 
    "github.com/loc-ne/go-auction/services/auth/internal/usecase" 
    "github.com/loc-ne/go-auction/services/auth/internal/delivery/http" 
    "github.com/gin-gonic/gin"
    "os"
    )

func main() {
    log.Println("Starting Auth Service...")
    secret := os.Getenv("JWT_SECRET") 
	
    ctx := context.Background()
    migrations.RunMigrations()

    db, err := postgres.NewPostgresDB(ctx)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Pool.Close()

    userRepo := postgres.NewUserRepository(db.Pool)
    userUC := usecase.NewUserUsecase(userRepo, secret)
	sessionRepo := postgres.NewSessionRepository(db.Pool)
	sessionUC := usecase.NewSessionUsecase(sessionRepo)
    handler := http.NewUserHandler(userUC, sessionUC)

    router := gin.Default()
    router.POST("/register", handler.Register)
    router.POST("/login", handler.Login)

    router.Run(":8080")
    
}