package main

import (
	"log"
	"healthprobe-for-app/cmd"
)

func main() {
	log.Println("Hello, healthprobe-for-app!")
	cmd.Execute()
}

// go run . -s 5 -r 10 -e 15 -p E:/dev/tmp
// go run . -s 5 -r 10 -e 15 -p /tmp