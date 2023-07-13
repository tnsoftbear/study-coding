package main

import (
	"fmt"
	"runtime"
)

func main() {
	someFunc()
	Class{}.someMethod()
}

func someFunc() {
	fmt.Println(currentFunctionName()) // main.someFunc
}

type Class struct {
}

func (obj Class) someMethod() {
	fmt.Println(currentFunctionName()) // main.Class.someMethod
}

func currentFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fn.Name()
}
