package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Pass     string
	Database string
	SSLMode  string
}

func (cfg Config) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Pass:     "root",
		Database: "lenseCode",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	_, err = db.Exec(`

		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL
			);
		create table if not exists orders (
			id serial PRIMARY KEY,
			user_id integer NOT NULL,
			amount integer NOT NULL,
			description text NOT NULL
		);

`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("created table")
}
