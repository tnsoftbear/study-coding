package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	
	log.Println("Space station is ready for spaceship arrival!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
