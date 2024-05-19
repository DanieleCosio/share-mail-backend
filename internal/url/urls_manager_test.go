package url

import (
	"slices"
	"testing"
)

func TestMapDiff(t *testing.T) {
	m1 := []string{"b", "c", "d", "e"}
	m2 := []string{"a", "b", "c"}
	result1 := []string{"d", "e"}
	result2 := []string{"a"}

	diff1 := mapDiff(&m1, &m2)
	diff2 := mapDiff(&m2, &m1)

	for _, v := range result1 {
		if !slices.Contains(*diff1, v) {
			t.Errorf("Expected %v, got %v", result1, diff1)
		}
	}

	for _, v := range result2 {
		if !slices.Contains(*diff2, v) {
			t.Errorf("Expected %v, got %v", result2, diff2)
		}
	}
}
