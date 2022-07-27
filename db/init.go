package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func ConnectPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println("Connecting to Postgres...")
	log.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Panicln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Successfully connected!")

	return db
}
