// Created by Anh Cao on 27.08.2024.

package main

import (
	"log"
	"net/http"

	"github.com/AnhCaooo/electric-push-notifications/internal/api/handlers"
	"github.com/AnhCaooo/electric-push-notifications/internal/api/middleware"
	"github.com/AnhCaooo/electric-push-notifications/internal/api/routes"
	"github.com/AnhCaooo/electric-push-notifications/internal/cache"
	"github.com/AnhCaooo/electric-push-notifications/internal/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize cache
	cache.NewCache()

	// Initial new router
	r := mux.NewRouter()
	for _, endpoint := range routes.Endpoints {
		r.HandleFunc(endpoint.Path, endpoint.Handler).Methods(endpoint.Method)
	}
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.NotAllowed)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	// Middleware
	r.Use(middleware.Logger)

	// Start server
	logger.Logger.Info("Server started on :8002")
	log.Fatal(http.ListenAndServe(":8002", r))
}
