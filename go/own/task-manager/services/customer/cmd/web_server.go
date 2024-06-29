package main

import (
	"task_manager/internal/config"
	"task_manager/internal/controller"
)

func main() {
	r := controller.SetupRouter()
	r.Run(config.GetStrEnv("APP_HOST", ":8080")) # todo: port
}