package main

import (
	"errors"
	"fmt"
)

//type error1 interface {
//	error
//}
//
//type error1 struct {
//	As(interface{}) bool
//}

func main() {
	var err = errors.New("no such key")
	x, ok := err.(interface{ As(interface{}) bool })
	if ok {
		if x.As(err) {
			fmt.Println("ok and x.As")
		} else {
			fmt.Println("ok and not x.As")
		}
	} else {
		fmt.Println("not ok")
	}
}
