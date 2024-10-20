package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akolybelnikov/xm-exercise/internal/kafka"
	"github.com/akolybelnikov/xm-exercise/internal/repository"

	"github.com/akolybelnikov/xm-exercise/internal/config"
)

func Run(cfg *config.Config) error {
	var err error
	// Initialize database
	repo, err := repository.NewPostgresCompanyRepository(&cfg.DB)
	if err != nil {
		return err
	}

	// initialize Kafka producer
	producer, err := kafka.NewMutationProducer(&cfg.Kafka)
	if err != nil {
		return err
	}
	producer.Start()

	// Create Kafka topic
	if err = kafka.CreateTopic(cfg.Kafka.Brokers, cfg.Kafka.Topic); err != nil {
		return err
	}

	routerCfg := &RouterConfig{
		Producer: producer,
		Repo:     repo,
		Topic:    cfg.Kafka.Topic,
		Secret:   cfg.App.Secret,
		Exp:      cfg.App.TokenExp,
	}

	// Create router
	r := NewRouter(routerCfg)

	// Start server
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.App.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.App.Timeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.App.IdleTimeout) * time.Second,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("Server started on port " + cfg.App.Port)

	// Graceful shutdown
	return gracefulShutdown(srv, producer, cfg.App.WaitTimeout, cfg.Kafka.FlushTimeout)
}

func gracefulShutdown(srv *http.Server, kp *kafka.Producer, waitTimeout int, timeout int) error {
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

	// Close Kafka producer
	kp.Close(timeout)
	log.Println("Kafka producer closed")

	return nil
}
