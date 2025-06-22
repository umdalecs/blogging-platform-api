package main

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgresDB(cfg *pgxpool.Config) *pgxpool.Pool {
	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
