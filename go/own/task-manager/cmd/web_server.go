package main

import (
	"task_manager/internal/controller"
)

func main() {
	r := controller.SetupRouter()
	r.Run(":8080")
}