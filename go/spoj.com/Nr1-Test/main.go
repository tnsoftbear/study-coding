package main

import "fmt"

func main() {
	var v string
	for ;; {
		fmt.Scan(&v)
		if v == "42" {
			return
		}
		fmt.Println(v)
	}
}
