package bubblesort

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/bseto/basic-algorithms/verifier"
)

func TestBubbleSort(t *testing.T) {
	b, err := ioutil.ReadFile(filepath.Join("..", "randomslice.txt"))
	if err != nil {
		t.Fatalf("unable to read the file: %v", err)
	}

	var randomizedSlice []int64
	err = json.Unmarshal(b, &randomizedSlice)
	if err != nil {
		t.Fatalf("unable to unmarshal the randomized slice: %v", err)
	}

	sortedSlice := BubbleSort(randomizedSlice)
	isSorted := verifier.VerifySorted(sortedSlice)
	if isSorted != true {
		t.Fatal("BubbleSort did not sort the slice")
	}
}
