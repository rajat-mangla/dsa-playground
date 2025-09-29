package sorting

// QuickSort sorts an array of integers using the quick sort algorithm.
func QuickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}

func qs(arr []int, low, high int) {
	if low < high {

		pIndex := partition(arr, low, high)
		qs(arr, low, pIndex-1)
		qs(arr, pIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}
