package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service...")
	var err error

	connection, err := connectToDB()
	if connection == nil {
		log.Println("\nCouldn't connect to Postgres!")
		panic(err)
		return
	}

	log.Print("Connected to Postgres!!\n\n")

	// Set up config
	app := Config{
		DB:     connection,
		Models: data.New(connection),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() (*sql.DB, error) {
	godotenv.Load(".env")
	dsn := os.Getenv("DSN")
	counts := 0

	for {
		connection, err := openDB(dsn)
		if connection != nil {
			return connection, nil
		}
		counts++

		if counts > dbMaxAttempts {
			log.Println(err)
			return nil, err
		}

		log.Println("Retrying connect in two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
