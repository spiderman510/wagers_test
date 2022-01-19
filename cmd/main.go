package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"wagers/pkg/api"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	databaseConnection := os.Getenv("DATABASE_CONNECTION")
	log.Info().Msg("Database connection: " + databaseConnection)
	if databaseConnection == "" {
		databaseConnection = "user:password@tcp(db:3306)/wager?parseTime=true"
	}
	db, err := sql.Open("mysql", databaseConnection)

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	defer db.Close()

	engine, err := api.CreateAPIEngine(db)

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	port := os.Getenv("LISTEN")
	if port == "" {
		port = ":8080"
	}
	log.Info().Msg("Starting HTTP server on port " + port)
	log.Fatal().Err(engine.Run(port)).Msg("Fail to listen and serve")
}
