package bubblesort

// BubbleSort should implement bubblesort to sort the input []int64 slice
func BubbleSort(unsortedSlice []int64) []int64 {

	unsortedSliceLength := len(unsortedSlice)
	for {
		swapped := false
		for index, element := range unsortedSlice {
			if index+1 == unsortedSliceLength {
				break
			}

			if element > unsortedSlice[index+1] {
				swapped = true
				unsortedSlice[index] = unsortedSlice[index+1]
				unsortedSlice[index+1] = element
			}
		}
		if swapped == false {
			// no swaps occurred, we are done sorting
			break
		}
	}
	return unsortedSlice
}
