package main

import (
	"fiber-reform-rest/internal/app/bootstrap"
	"fiber-reform-rest/internal/infra/env"

	"fmt"
	"log"
)

func main() {
	app := bootstrap.NewApp()
	listenAddr := fmt.Sprintf("%s:%s",
		env.GetEnv("APP_HOST", "localhost"),
		env.GetEnv("APP_PORT", "4000"))
	log.Fatal(app.Listen(listenAddr))
}
