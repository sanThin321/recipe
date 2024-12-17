package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-cp2pr5kf7o1s73bl8mog-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "root"
	postgres_password = "MRjiBRpfZvYhLCobB9ny9HBV5KMEAzRe"
	postgres_dbname   = "backendprj"
)

var Db *sql.DB

func init() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	Db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Test the connection by pinging the database
	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Database connection successful")
}
