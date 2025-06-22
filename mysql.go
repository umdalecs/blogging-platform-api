package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitMysqlDB(cfg mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	testConnection(db)

	return db
}

func testConnection(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
