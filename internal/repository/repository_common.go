package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

const wagerTableName = "wagers"
const databaseTestName = "wager_test"
const driverName = "mysql"
const purchaseTableName = "purchases"

var mockDatabaseUrl = fmt.Sprintf("user:password@tcp(localhost:3307)/%s?parseTime=true", databaseTestName)

var db *sql.DB

func getLastIndexId(db *sql.DB, tableName string) int {
	rows, err := db.Query("SELECT AUTO_INCREMENT FROM information_schema.tables WHERE table_name = ?", tableName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query increment value")
		return -1
	}
	defer rows.Close()

	if !rows.Next() {
		return -1
	}
	var latestIndex int
	err = rows.Scan(&latestIndex)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query increment value")
		return -1
	}
	return latestIndex
}

func initDb() {
	database, err := sql.Open(driverName, mockDatabaseUrl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connection database")
	}
	db = database
}

func tearDown() {
	db.Close()
}
