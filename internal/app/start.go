package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akolybelnikov/xm-exercise/internal/repository"

	"github.com/akolybelnikov/xm-exercise/internal/config"
)

func Run(cfg *config.Config) error {
	// Initialize database
	repo, dbErr := repository.NewPostgresCompanyRepository(&cfg.DB)
	if dbErr != nil {
		return dbErr
	}

	// Create router
	r := NewRouter(repo)

	// Start server
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.App.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.App.Timeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.App.IdleTimeout) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("Server started on port " + cfg.App.Port)

	// Graceful shutdown
	return gracefulShutdown(srv, cfg.App.WaitTimeout)
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
