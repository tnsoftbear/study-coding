package main

var x = 0

func f() int {
	x = 3
	return x
}

func main() {
	x = 0
	a, _ := x, f()

	x = 0
	var b, _ = x, f() // так b=0
	// f(); b := x // так b=3
	// f(); var b = x // так b=3
	// b, _ := x, f() // так b=3
	// var b int; b, _ = x, f() // и так b=3
	println(a, b) // 3 0
}
