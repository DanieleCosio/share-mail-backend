package urlgen

import (
	"math"
	"testing"
)

func TestGenerateUniqueStrings(t *testing.T) {
	length := 2
	printables := GeneratePrintableArray()
	strings := GenerateUniqueStrings(length)
	count := math.Pow(float64(len(printables)), float64(length))

	if len(strings) != int(count) {
		t.Errorf("Expected %d strings, got %d", int(count), len(strings))
	}
}
