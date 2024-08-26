package main

type str string

func (s str) log() {
	println(s)
}

func main() {
	var s str = "hello"
	s.log()
}
