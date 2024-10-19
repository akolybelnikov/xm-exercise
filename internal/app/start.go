package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akolybelnikov/xm-exercise/internal/config"
)

func Run(cfg *config.Config) error {
	// Create router
	r := NewRouter()

	// Start server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.App.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.App.Timeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.App.Idle) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	return gracefulShutdown(srv, cfg.App.Wait)
}

func gracefulShutdown(srv *http.Server, waitTimeout int) error {
	quit := make(chan os.Signal, 1)

	// Listen for interrupt signals
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitTimeout)*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server gracefully stopped")
	return nil
}
