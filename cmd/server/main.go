package main

import (
	"context"
	"log"
	"os"

    "snyff/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
    ctx := context.Background()

    // Create a connection pool to the db
    pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v", err)
    }
    if err := pool.Ping(ctx); err != nil { 
        log.Fatalf("Unable to reach the database: %v", err)
    }
    defer pool.Close()

    db := stdlib.OpenDBFromPool(pool)
    defer db.Close()

    goose.SetBaseFS(migrations.FS)
    if err := goose.SetDialect("postgres"); err != nil {
        log.Fatalf("Failed to set goose dialect: %v", err)
    }

    // Perform db migrations using goose
    if err := goose.Up(db, "."); err != nil {
       log.Fatalf("Failed to run migrations: %v", err)
    }
}