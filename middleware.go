package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next
	}
}

type httpWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *httpWriterWrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &httpWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.statusCode, r.URL.Path, r.Method, time.Since(start))
	})
}
