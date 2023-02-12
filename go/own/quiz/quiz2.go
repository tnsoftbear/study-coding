package main

var (
	_ = f("w", x) // 3rd
	x = f("x", z) // 2nd
	y = f("y", x) // 4th
	z = f("z")    // 1st
)

func f(s string, deps ...int) int {
	print(s)
	return 0
}

func main() {
	f("\n")
}

// zxwy
