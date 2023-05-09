package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func newDatabase() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DB")
	sslMode := os.Getenv("DB_SSLMODE")

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("pgx", config)
	if err != nil {
		log.Fatalf("could not create a connection to the database, err: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
