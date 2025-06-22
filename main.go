package main

import (
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	cfg.ConnConfig.Host = Envs.DbAddr
	cfg.ConnConfig.Port = Envs.DbPort
	cfg.ConnConfig.User = Envs.DbUser
	cfg.ConnConfig.Password = Envs.DbPass
	cfg.ConnConfig.Database = Envs.DbName
	cfg.ConnConfig.TLSConfig = nil
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour

	db := InitPostgresDB(cfg)

	s := NewApiServer(":8080", db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
