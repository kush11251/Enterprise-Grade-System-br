package http

import (
	"example.com/enterprise-grade-system/pkg/controllers"
	"example.com/enterprise-grade-system/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

// StartServer starts the HTTP server
:func StartServer(cfg *config.Config) {
		router := mux.NewRouter()
		router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
		router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
		handler := logger.Handler(router)
		http.ListenAndServe("":" + strconv.Itoa(cfg.Server.Port), handler)
}