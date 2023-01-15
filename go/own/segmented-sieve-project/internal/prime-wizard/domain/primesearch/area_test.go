package primesearch

import (
	"reflect"
	"segmented-sieve-project/internal/prime-wizard/domain/types"
	"testing"
)

func TestArea_NewArea(t *testing.T) {
	minMax := types.Range{Min: 1, Max: 10}
	area := NewArea(minMax, 10)
	if area.MinMax != minMax {
		t.Errorf("MinMax is not set correctly")
	}
	if len(area.Statuses) != 10 {
		t.Errorf("Statuses slice is not set correctly")
	}
	expected := []uint32{0, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(area.Statuses, expected) {
		t.Errorf("Statuses slice is not set correctly. Actual: %v, expected: %v", area.Statuses, expected)
	}
}

func TestArea_IsPrimeOverPossible(t *testing.T) {
	minMax := types.Range{Min: 1, Max: 5}
	area := NewArea(minMax, 5)
	if !area.IsPrimeOverPossible(3) {
		t.Errorf("IsPrimeOverPossible is not working correctly")
	}
	if area.IsPrimeOverPossible(2) {
		t.Errorf("IsPrimeOverPossible is not working correctly")
	}
}

func TestArea_Sieve(t *testing.T) {
	testData := []struct {
		minMax   types.Range
		areaSize uint32
		expected []uint32
	}{
		{
			minMax:   types.Range{Min: 11, Max: 20},
			areaSize: 10,
			expected: []uint32{11, 0, 13, 0, 15, 0, 17, 0, 19, 0},
		},
		{
			minMax:   types.Range{Min: 1, Max: 5},
			areaSize: 5,
			expected: []uint32{0, 2, 3, 0, 5},
		},
		{
			minMax:   types.Range{Min: 101, Max: 120},
			areaSize: 20,
			expected: []uint32{101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0},
		},
	}

	for _, v := range testData {
		area := NewArea(v.minMax, v.areaSize)
		area.Sieve(2)
		if !reflect.DeepEqual(area.Statuses, v.expected) {
			t.Errorf("Statuses slice is not set correctly. Actual: %v, expected: %v", area.Statuses, v.expected)
		}
	}
}
