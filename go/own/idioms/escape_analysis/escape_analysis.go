package main

func f() *int {
	v := 0x100*0x100 - 1
	return &v
}

func main() {
	var p = f()
	println(*p)
}
