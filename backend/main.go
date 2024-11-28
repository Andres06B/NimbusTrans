package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	database "proyecto/Database"
	Cors "proyecto/Middleware"
	routes "proyecto/Routes"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "SISTEMA: ", log.LstdFlags)
	router := mux.NewRouter()

	database.Init()

	routes.ConductoRoutes(logger, router)
	routes.BusRoutes(logger, router)
	routes.RutaRoutes(logger, router)
	routes.AsignacionRoutes(logger, router)

	cors := Cors.CorsMiddleware(router)

	// Iniciar el servidor
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      cors,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Iniciar el servidor en un goroutine
	go func() {
		logger.Println("Starting server in http://localhost:8080")
		err := srv.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Iniciar la funcion de shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	logger.Println("Received terminated, graceful shutdown", sig)
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	// Cerrar el servidor
	if err != nil {
		logger.Println(err)
	}
	srv.Shutdown(tc)

}
