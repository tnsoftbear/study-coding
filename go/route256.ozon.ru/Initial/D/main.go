package main

import (
	"bufio"
	"fmt"
	"os"
)

const YES = "YES"
const NO = "NO"

func check(pw string, verified *map[string]int) bool {
	var pwLen = len(pw)
	if pwLen < 2 || pwLen > 24 || pw[0] == '-' {
		return false
	}

	var normalizedB = []byte(pw)
	for i, c := range pw {
		if c == 45 || c == 95 || (c >= 48 && c <= 57) || (c >= 97 && c <= 122) {
			continue
		} else if c >= 65 && c <= 90 {
			normalizedB[i] = byte(c + 32)
			continue
		} else {
			return false
		}
	}

	var normalized = string(normalizedB)
	if _, ok := (*verified)[normalized]; ok {
		(*verified)[normalized]++
		return false
	} else {
		(*verified)[normalized] = 0
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		var passwordsCount int
		fmt.Fscanf(in, "%d\n", &passwordsCount)

		var passwords = make([]string, passwordsCount)
		for i := 0; i < passwordsCount; i++ {
			fmt.Fscanf(in, "%s\n", &passwords[i])
		}

		var verified = make(map[string]int, passwordsCount)
		var j = 0
		for i := 0; i < passwordsCount; i++ {
			if check(passwords[i], &verified) {
				j++
				fmt.Fprintln(out, YES)
			} else {
				fmt.Fprintln(out, NO)
			}
		}

		fmt.Fprintln(out)
	}
}
