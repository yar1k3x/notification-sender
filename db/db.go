package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(user, password, host, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:36885)/%s?parseTime=true", user, password, host, dbname)
	log.Println("Connecting to database with DSN:", dsn)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open error: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("DB.Ping error: %w", err)
	}

	return nil
}
