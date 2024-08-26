package main

import "fmt"

// HelloWorld greets the world.
func HelloWorld() {
	var p *int
	var a int = 1
	p = &a
	var b int = *p
	*p = *p + 2
	println(fmt.Sprintf("a: %d, b: %b", a, b))
}

func main() {
	HelloWorld()
}
