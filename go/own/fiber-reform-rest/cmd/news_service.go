package main

import (
	"fiber-reform-rest/internal/api/bootstrap"
	"fiber-reform-rest/internal/infra/env"
	"os"
	"os/signal"
	"syscall"

	"fmt"
	"log"
)

func main() {
	app := bootstrap.NewApp()
	
	go func() {
		listenAddr := fmt.Sprintf("%s:%s",
			env.GetEnv("APP_HOST", "localhost"),
			env.GetEnv("APP_PORT", "4000"))
		if err := app.Listen(listenAddr); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	fmt.Println("Fiber was successful shutdown.")
}
