package recursive-sum

import "fmt"

func sum(n int) int {
	if n == 1 { 
		return 1
	}
	return n + sum(n-1)
}

func main() {
	println(fmt.Sprintf("Result: %d", sum(100)))
}
