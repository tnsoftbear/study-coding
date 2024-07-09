package main

import (
	"task_manager/internal/config"
	"task_manager/internal/controller/route"
)

func main() {
	r := route.SetupRouter()
	r.Run(config.GetStrEnv("APP_HOST", ":8080")) // todo: port
}