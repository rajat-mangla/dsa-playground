package sorting

// SelectionSort sorts an array of integers using the selection sort algorithm.
func SelectionSort(arr []int) []int {
	for i := range arr {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr
}
