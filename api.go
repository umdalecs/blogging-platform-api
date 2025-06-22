package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewApiServer(addr string, db *pgxpool.Pool) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	e := gin.Default()

	v1 := e.Group("/api/v1")

	postRepository := NewPostRepository(s.db)
	postHandler := NewPostHandler(postRepository)
	postHandler.RegisterRoutes(v1)

	return e.Run(s.addr)
}
