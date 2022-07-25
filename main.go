package main

import (
	"database/sql"
	"embed"
	"log"
	"net/http"
	"os"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sabinlehaci/go-web-app/db"
	"github.com/sabinlehaci/go-web-app/handler"
	"github.com/sabinlehaci/go-web-app/tmdbApi"
)


func main() {
	// For example: POSTGRES_URL="postgres://postgres:mysecretpassword@localhost:5432/postgres"
	database, err := sql.Open("pgx", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("oops, db connection failed", err)
	}

	err = validateSchema(database)
	if err != nil {
		log.Fatal("oops, db migration failed", err)
	}

	log.Println("Serving on 0.0.0.0:9090")

	http.ListenAndServe(":9090", &handler.Handlers{
		MovieGetter: &tmdbApi.Client{
			APIKey: os.Getenv("TMDB"),
			
		},
		DB: db.New(database),
	})
}

//go:embed db/migrations/*.sql
var fs embed.FS

// Migrate migrates the Postgres schema to the current version.

func validateSchema(db *sql.DB) (retErr error) {
	sourceInstance, err := iofs.New(fs, "db/migrations")
	if err != nil {
		return err
	}
	defer func() {
		err := sourceInstance.Close()
		if retErr == nil {
			retErr = err
		}
	}()
	driverInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, "postgres", driverInstance)
	if err != nil {
		return err
	}
	err = m.Up() // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
