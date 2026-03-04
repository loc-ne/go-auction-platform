package main

import (
    "context"
    "log"
    "os"
    
    "github.com/loc-ne/go-auction/services/auth/migrations" 
    "github.com/loc-ne/go-auction/services/auth/internal/repository/postgres" 
    "github.com/loc-ne/go-auction/services/auth/internal/usecase" 
    "github.com/loc-ne/go-auction/services/auth/internal/delivery/http" 
    "github.com/loc-ne/go-auction/services/auth/internal/delivery/http/middleware" 
    "github.com/gin-gonic/gin"
    // "github.com/joho/godotenv"
)

func main() {
    // log.Println("Starting Auth Service...")
    // if err := godotenv.Load(); err != nil {
	// 	log.Println(".env not found")
	// }
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
    authGroup := router.Group("/api/v1/auth")
    {
        authGroup.POST("/login", handler.Login)
        authGroup.POST("/register", handler.Register)
        
        protected := authGroup.Use(middleware.AuthMiddleware(secret))
        {
            protected.GET("/me", handler.Me)
           // protected.POST("/logout", handler.Logout)
        }
    }

    port := os.Getenv("AUTH_PORT")
    log.Printf("Auth Service listening on port %s...\n", port)
    router.Run(":" + port)
    
}