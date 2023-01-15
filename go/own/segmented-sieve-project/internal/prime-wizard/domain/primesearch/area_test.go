package primesearch

import (
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
	for i, v := range area.Statuses {
		if v != expected[i] {
			t.Errorf("Statuses[%d] is not set correctly", i)
		}
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

func TestArea_Sieve1(t *testing.T) {
	minMax := types.Range{Min: 1, Max: 5}
	area := NewArea(minMax, 5)
	area.Sieve(2)
	expected := []uint32{0, 2, 3, 0, 5}
	for i, v := range area.Statuses {
		if v != expected[i] {
			t.Errorf("Statuses[%d] is not set correctly", i)
		}
	}
}

func TestArea_Sieve2(t *testing.T) {
	minMax := types.Range{Min: 11, Max: 20}
	area := NewArea(minMax, 10)
	area.Sieve(2)
	expected := []uint32{11, 0, 13, 0, 15, 0, 17, 0, 19, 0}
	for i, v := range area.Statuses {
		if v != expected[i] {
			t.Errorf("Statuses[%d] is not set correctly", i)
		}
	}
}
