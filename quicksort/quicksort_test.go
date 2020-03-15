package quicksort

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/bseto/basic-algorithms/verifier"
)

func TestQuickSort(t *testing.T) {
	b, err := ioutil.ReadFile(filepath.Join("..", "testSlices", "100.txt"))
	if err != nil {
		t.Fatalf("unable to read the file: %v", err)
	}

	var randomizedSlice []int64
	err = json.Unmarshal(b, &randomizedSlice)
	if err != nil {
		t.Fatalf("unable to unmarshal the randomized slice: %v", err)
	}

	high := len(randomizedSlice) - 1

	sortedSlice := Quicksort(randomizedSlice, 0, int64(high))
	isSorted := verifier.VerifySorted(sortedSlice)
	if isSorted != true {
		t.Fatal("Quicksort did not sort the slice")
	}
}

func TestSwap(t *testing.T) {
	testSlice := []int64{11, 22}
	swap(&testSlice[0], &testSlice[1])
	t.Errorf("is this right: %v", testSlice)
}
