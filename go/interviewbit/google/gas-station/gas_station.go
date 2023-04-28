package main;

func main() {
	a := []int{1, 2}
	b := []int{2, 1}
	c := canCompleteCircuit(a, b)
	println(c)
}

/**
 * @input A : Integer array
 * @input B : Integer array
 * 
 * @Output Integer
 */
func canCompleteCircuit(A []int , B []int )  (int) {
	// compare array size
	if len(A) != len(B) {
		return -1
	}

	// find start point
	for i := 0; i < len(A); i++ {
		// check if start point is valid
		if A[i] < B[i] {
			continue
		}

		if checkCircuit(A, B, i) {
			return i
		}
	}

	return -1
}

func checkCircuit(A []int, B []int, start int) (bool) {
	currentGas := 0
	for i := start; i < len(A); i++ {
		currentGas = currentGas + A[i] - B[i]
		if currentGas < 0 {
			return false
		}
	}

	for i := 0; i < start; i++ {
		currentGas = currentGas + A[i] - B[i]
		if currentGas < 0 {
			return false
		}
	}

	return true
}