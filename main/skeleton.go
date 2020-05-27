package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/scharissis/go-server-skeleton/skeleton"
)

func main() {
	apiPort := ":" + skeleton.GetOrDefault("API_PORT", "8000")
	apiPrefix := skeleton.GetOrDefault("API_PREFIX", "/api")

	var wg sync.WaitGroup
	defer wg.Wait() // application will not exit until WaitGroup empty

	srv := skeleton.NewServer(apiPrefix)
	s := &http.Server{
		Addr:         apiPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      srv,
	}

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		sig := <-sigs
		signal.Stop(sigs) // to allow force-exit on double ^C
		wg.Add(1)
		defer wg.Done()
		log.Printf("received %s, shutting down", sig)
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("error whilst shutting down HTTP server: %v\n", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("error from ListenAndServe: %v\n", err)
	}
	log.Printf("Server running (port %s)...\n", apiPort)
}
