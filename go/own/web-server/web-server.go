package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-server/internal/config"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	cfg := config.Init()

	http.HandleFunc("/", hello)
	
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{
			"pong": true,
		})
	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		pid := os.Getpid()
		process, _ := os.FindProcess(pid)
		if err := process.Signal(syscall.SIGTERM); err != nil {
			log.Printf("Cannot send SIGTERM signal to the current process (PID: %d), because of error: %s", pid, err)
		}
	})

	srv := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: http.DefaultServeMux,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", srv.Addr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped
}

// https://pkg.go.dev/net/http	
