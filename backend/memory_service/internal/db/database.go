package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() *Database {
	// Get environment variables
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatalf("failed to parse DB_MAX_OPEN_CONNS: %v", err)
	}
	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("failed to parse DB_MAX_IDLE_CONNS: %v", err)
	}
	maxConnLifetime, err := time.ParseDuration(os.Getenv("DB_MAX_CONN_LIFETIME_MIN"))
	if err != nil {
		log.Fatalf("failed to parse DB_MAX_CONN_LIFETIME: %v", err)
	}

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	// Set database connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxConnLifetime * time.Minute)

	// Health check
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("connected to database")
	return &Database{Conn: db}
}

func (d *Database) Close() error {
	// Close database connection
	log.Println("closing database connection")
	return d.Conn.Close()
}


