package sorting

// InsertionSort sorts an array of integers using the insertion sort algorithm.
func InsertionSort(arr []int) []int {
	for i, elem := range arr {
		j := i - 1
		for j >= 0 && arr[j] > elem {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
		arr[j+1] = elem
	}
	return arr
}
