package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ts22082/logging-service/handlers/api"
	"github.com/ts22082/logging-service/handlers/views"
	mongodb_client "github.com/ts22082/logging-service/utils/mongodb"
	redis_client "github.com/ts22082/logging-service/utils/redis"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env: %v: ", err)
	}

	if err := redis_client.Init(nil); err != nil {
		log.Fatalf("Could not connect to Redis Server: %v", err)
	}
	redis_client.SetupGracefulShutdown()

	if err := mongodb_client.Init(nil); err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	mongodb_client.SetupGracefulShutdown()

	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	)

	router.HandleFunc("/", views.Home).Methods("GET")
	router.HandleFunc("/docs/{subject}", views.Docs).Methods("GET")
	router.HandleFunc("/login", views.Login).Methods("GET")
	router.HandleFunc("/logout", views.Logout).Methods("GET")
	router.HandleFunc("/dashboard", views.Dashboard).Methods("GET")
	router.HandleFunc("/dashboard/logs/{projectId}", views.ProjectLogs).Methods("GET")
	router.HandleFunc("/email/login_validate/{login_token}", views.LoginValidate).Methods("GET")
	router.HandleFunc("/admin", views.Admin).Methods("GET")

	apiRoute := router.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/log", api.Log).Methods("POST")
	apiRoute.HandleFunc("/send_login_link", api.SendLoginLink).Methods("POST")
	apiRoute.HandleFunc("/project", api.Project).Methods("POST")
	apiRoute.HandleFunc("/api-key", healthCheck).Methods("POST")
	apiRoute.HandleFunc("/dashboard/logs/{projectId}/stream", api.ProjectLogsStream).Methods("GET")
	apiRoute.HandleFunc("/health", healthCheck).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		fmt.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("\nShutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	} else {
		fmt.Println("Server shut down successfully")
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, router *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, router)
		log.Printf("%s %s %s %v", router.Method, router.RequestURI, router.RemoteAddr, time.Since(start))
	})
}

func healthCheck(w http.ResponseWriter, router *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
