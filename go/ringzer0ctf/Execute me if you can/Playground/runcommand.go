package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("./testshell", "0>&1").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Output is %s\n", out)
}