package main

func main() {
	a := []int{1, 2}
	b := []int{2, 1}
	c := canCompleteCircuit2(a, b)
	println(c)
}

func canCompleteCircuit2(gas []int, dist []int) int {
	n := len(gas)
	startIndex := 0
	tankSoFar := 0
	for i := 0; i < n*2; i++ {
		g := gas[i%n]
		d := dist[i%n]
		tankSoFar += g - d
		if tankSoFar < 0 {
			startIndex = i + 1
			tankSoFar = 0
		} else if startIndex+n == i {
			return startIndex
		}
	}
	return -1
}
