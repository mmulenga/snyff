package main

import (
	"fmt"
    "database/sql"
    "embed"
    "github.com/pressly/goose/v3"
)

    //go.embed migrations/*.sql
    var embedMigrations embed.FS

func main() {
    var db *sql.DB
    goose.SetBaseFS(embedMigrations)

    if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }

    if err := goose.Up(db, "migrations"); err != nil {
        panic(err)
    }
}