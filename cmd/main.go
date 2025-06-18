package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/umdalecs/blogging-platform-api/api"
	"github.com/umdalecs/blogging-platform-api/config"
	"github.com/umdalecs/blogging-platform-api/db"
)

func main() {
	db := db.InitMysqlDB(mysql.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DbPassw,
		Addr:                 config.Envs.DbAddr,
		DBName:               config.Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	s := api.NewApiServer(":8080", db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
