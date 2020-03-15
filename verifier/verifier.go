package verifier

// VerifySorted will check that the input sorted array is what it says it is -
// that is is indeed sorted.
func VerifySorted(sortedArray []int64) bool {
	sortedArrayLen := len(sortedArray)
	sorted := true

	for index, element := range sortedArray {
		if index+1 == sortedArrayLen {
			break
		}

		if element > sortedArray[index+1] {
			sorted = false
			break
		}
	}

	return sorted
}
