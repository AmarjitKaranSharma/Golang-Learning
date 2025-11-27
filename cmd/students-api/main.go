package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AmarjitKaranSharma/golang-student-api/internal/config"
	"github.com/AmarjitKaranSharma/golang-student-api/internal/http/handlers/student"
	"github.com/AmarjitKaranSharma/golang-student-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("Hello, Students API!")

	// Load configuration
	cfg := config.MustLoad()

	// database connection setup
	_, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	// start server
	// fmt.Printf("Server Started %s", cfg.HttpServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Server Started", slog.String("address", cfg.HttpServer.Address))
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Failed to start server:", err)
		}
	}()

	<-done

	slog.Info("Shutting Down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
