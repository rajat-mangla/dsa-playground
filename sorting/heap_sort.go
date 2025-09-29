package sorting

// HeapSort sorts an array of integers using the heap sort algorithm.
func HeapSort(arr []int) []int {
	l := len(arr)

	// Build max heap
	for i := l/2 - 1; i >= 0; i-- {
		heapify(arr, l, i)
	}

	// One by one extract elements from heap
	for i := l - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, l, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < l && arr[largest] < arr[left] {
		largest = left
	}

	if right < l && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, l, largest)
	}
}
