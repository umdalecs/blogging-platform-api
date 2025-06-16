package db

import (
	"database/sql"
	"log"

	mysqlDrv "github.com/go-sql-driver/mysql"
)

func InitMysqlDB(cfg mysqlDrv.Config) *sql.DB {
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
