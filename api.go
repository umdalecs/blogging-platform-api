package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Db:   db,
	}
}

func (s *ApiServer) Run() error {
	v1 := http.NewServeMux()

	postRepository := NewPostRepository(s.Db)
	postHandler := NewPostHandler(postRepository)
	postHandler.RegisterRoutes(v1)

	r := http.NewServeMux()
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	mc := MiddlewareChain(
		LoggerMiddleware,
	)

	server := &http.Server{
		Addr:    s.Addr,
		Handler: mc(r),
	}

	log.Printf("Server listening port %s", s.Addr)
	return server.ListenAndServe()
}

func WriteJson(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func WriteJsonErr(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func WriteEmpty(w http.ResponseWriter, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}
