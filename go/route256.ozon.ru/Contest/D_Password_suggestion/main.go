package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkVovel(b byte) bool {
	// e, u, i, o, a, y
	return b == 97 || b == 101 || b == 105 || b == 111 || b == 117 || b == 121
}

func checkAndAppend(pw string) string {
	var isUpper, isLower, isVovel, isConsonantal, isDigit bool
	var b byte
	for _, c := range pw {
		b = byte(c)
		if b >= 65 && b <= 90 {
			isUpper = true
			if checkVovel(b + 32) {
				isVovel = true
			} else {
				isConsonantal = true
			}
		} else if b >= 97 && b <= 122 {
			isLower = true
			if checkVovel(b) {
				isVovel = true
			} else {
				isConsonantal = true
			}
		} else if b >= 48 && b <= 57 {
			isDigit = true
		}
		if isUpper && isLower && isVovel && isConsonantal && isDigit {
			return pw
		}
	}
	if !isUpper {
		if !isVovel {
			pw = pw + "A"
			isVovel = true
		} else if !isConsonantal {
			pw = pw + "B"
			isConsonantal = true
		} else {
			pw = pw + "C"
		}
	}
	if !isLower {
		if !isVovel {
			pw = pw + "a"
			isVovel = true
		} else if !isConsonantal {
			pw = pw + "b"
			isConsonantal = true
		} else {
			pw = pw + "c"
		}
	}

	if !isVovel {
		pw = pw + "a"
		isVovel = true
	} else if !isConsonantal {
		pw = pw + "b"
		isConsonantal = true
	}

	if !isDigit {
		pw = pw + "0"
	}
	return pw
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	var pw string
	for testIdx := 0; testIdx < testCount; testIdx++ {
		fmt.Fscanf(in, "%s\n", &pw)
		pw = checkAndAppend(pw)
		fmt.Fprintln(out, pw)
	}
}
