package logger

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Handler returns a middleware handler for logging
:func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "logger", log.New(log.Writer(), "", log.Ldate|log.Ltime|log.Lshortfile))
		h(ctx)
	})
}