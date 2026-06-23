package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

//go.embed migrations/*.sql
    var embedMigrations embed.FS

func main() {
    ctx := context.Background()

    // Create a connection pool to the db
    pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
    if err != nil {
        fmt.Print("Unable to create connection pool: ", err)
    }
    if err := pool.Ping(ctx); err != nil { 
        fmt.Println("Unable to reach the database: ", err)
    }
    defer pool.Close()

    db := stdlib.OpenDBFromPool(pool)
    defer db.Close()

    
    goose.SetBaseFS(embedMigrations)
    if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }

    // Perform db migrations using goose
    if err := goose.Up(db, "migrations"); err != nil {
        panic(err)
    }
}