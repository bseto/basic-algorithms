package verifier

import (
	"testing"
)

var TestVerifySortedProvider = []struct {
	Name     string
	Slice    []int64
	IsSorted bool
}{
	{
		Name:     "Easy Sorted",
		Slice:    []int64{1, 2, 3, 4, 5},
		IsSorted: true,
	},
	{
		Name:     "Easy Not Sorted",
		Slice:    []int64{5, 2, 3, 4, 1},
		IsSorted: false,
	},
	{
		Name:     "Duplicated Sorted",
		Slice:    []int64{1, 1, 1, 5, 5},
		IsSorted: true,
	},
}

func TestVerifySorted(t *testing.T) {
	for _, testData := range TestVerifySortedProvider {
		t.Run(testData.Name, func(t *testing.T) {
			isSorted := VerifySorted(testData.Slice)
			if testData.IsSorted != isSorted {
				t.Errorf("expected: %v, got: %v", testData.IsSorted, isSorted)
			}
		})
	}
}
