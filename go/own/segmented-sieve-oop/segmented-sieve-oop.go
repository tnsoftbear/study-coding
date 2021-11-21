package main

import (
	"segmented-sieve-oop/application"
)

func main() {
	var app *application.App = &application.App{}
	app.Run()
}

// 2146483647 2147483647
// 9999000000 10000000000
// 999900000 1000000000
// 999800000 999900000
// 999700000 999800000
