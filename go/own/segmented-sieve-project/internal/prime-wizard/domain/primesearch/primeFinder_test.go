package primesearch

import (
	"reflect"
	"segmented-sieve-project/internal/prime-wizard/domain/types"
	"testing"
)

func TestPrimeFinder_NewPrimeFinder(t *testing.T) {
	actual := NewPrimeFinder(10)
	expected := &PrimeFinder{
		value:          0,
		numberStatuses: []uint8{2, 1, 1, 2, 0, 2, 0, 2, 2, 0},
		primeNumbers:   []uint32{2, 3},
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
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func TestPrimeFinder_DetectPrimes(t *testing.T) {
	testData := []struct {
		areaSize uint32
		minMax   types.Range
		expected []uint32
	}{
		{
			areaSize: 100,
			minMax:   types.Range{Min: 11, Max: 30},
			expected: []uint32{11, 13, 17, 19, 23, 29},
		},
		{
			areaSize: 100,
			minMax:   types.Range{Min: 100, Max: 130},
			expected: []uint32{101, 103, 107, 109, 113, 127},
		},
		{
			areaSize: 100,
			minMax:   types.Range{Min: 1000, Max: 1030},
			expected: []uint32{1009, 1013, 1019, 1021},
		},
	}
	for _, v := range testData {
		pf := NewPrimeFinder(v.areaSize)
		actual := pf.DetectPrimes(v.minMax)
		if !reflect.DeepEqual(actual, v.expected) {
			t.Errorf("actual: %v, expected: %v", actual, v.expected)
		}
	}
}
