package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db := InitMysqlDB(mysql.Config{
		User:                 Envs.DbUser,
		Passwd:               Envs.DbPassw,
		Addr:                 Envs.DbAddr,
		DBName:               Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	s := NewApiServer(":8080", db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
