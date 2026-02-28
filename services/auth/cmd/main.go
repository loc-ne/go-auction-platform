package main

import (
    "context"

)

func main() {
    log.Println("Starting Auth Service...")
    ctx := context.Background()
    migrations.RunMigrations()

    db, err := postgres.NewPostgresDB(ctx)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Pool.Close()


    
}