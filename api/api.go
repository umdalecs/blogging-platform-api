package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/umdalecs/blogging-platform-api/post"
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
	v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	postRepository := post.NewPostRepository(s.Db)
	postHandler := post.NewPostHandler(postRepository)
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
