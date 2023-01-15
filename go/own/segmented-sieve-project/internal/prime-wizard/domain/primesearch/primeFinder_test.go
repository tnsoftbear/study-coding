package primesearch

import (
	"segmented-sieve-project/internal/prime-wizard/domain/types"
	"testing"
)

func TestPrimeFinder_NewPrimeFinder(t *testing.T) {
	actual := NewPrimeFinder(10, false)
	expected := &PrimeFinder{
		value:          0,
		numberStatuses: []uint8{2, 1, 1, 2, 0, 2, 0, 2, 2, 0},
		primeNumbers:   []uint32{2, 3},
		isResultOutput: false,
		areaSize:       10,
	}
	if actual.value != expected.value {
		t.Errorf("actual value: %d, expected value: %d", actual.value, expected.value)
	}
	if len(actual.numberStatuses) != len(expected.numberStatuses) {
		t.Errorf("actual numberStatuses: %v, expected numberStatuses: %v", actual.numberStatuses, expected.numberStatuses)
	}
	for i, v := range actual.numberStatuses {
		if v != expected.numberStatuses[i] {
			t.Errorf("actual numberStatuses: %v, expected numberStatuses: %v", actual.numberStatuses, expected.numberStatuses)
		}
	}
	if len(actual.primeNumbers) != len(expected.primeNumbers) {
		t.Errorf("actual primeNumbers: %v, expected primeNumbers: %v", actual.primeNumbers, expected.primeNumbers)
	}
	for i, v := range actual.primeNumbers {
		if v != expected.primeNumbers[i] {
			t.Errorf("actual primeNumbers: %v, expected primeNumbers: %v", actual.primeNumbers, expected.primeNumbers)
		}
	}
	if actual.isResultOutput != expected.isResultOutput {
		t.Errorf("actual isResultOutput: %v, expected isResultOutput: %v", actual.isResultOutput, expected.isResultOutput)
	}
}

func TestPrimeFinder_DetectPrimes(t *testing.T) {
	pf := NewPrimeFinder(100, false)
	actual := pf.DetectPrimes(types.Range{Min: 11, Max: 30})
	expected := []uint32{11, 13, 17, 19, 23, 29}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}
	}
}
