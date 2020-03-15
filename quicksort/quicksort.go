package quicksort

// Swap will swap the two numbers
func swap(a, b *int64) {
	temp := *a
	*a = *b
	*b = temp
}

// Quicksort should implement quicksort to sort the input []int64 slice
func Quicksort(unsortedSlice []int64, low, high int64) []int64 {
	if low < high {
		partitionIndex := Partition(unsortedSlice, low, high)
		Quicksort(unsortedSlice, low, partitionIndex-1)
		Quicksort(unsortedSlice, partitionIndex+1, high)
	}

	return unsortedSlice
}

// Partition will follow the original implementation of quicksort
// where the pivot starts at the last (rightmost) element
func Partition(unsortedSlice []int64, low, high int64) int64 {
	pivot := unsortedSlice[high]
	largerNumberIndex := low

	for cursor := low; cursor <= high; cursor++ {
		if unsortedSlice[cursor] <= pivot {
			swap(&unsortedSlice[cursor], &unsortedSlice[largerNumberIndex])
			largerNumberIndex++
		}
	}
	swap(&unsortedSlice[largerNumberIndex], &unsortedSlice[high])

	return largerNumberIndex
}
