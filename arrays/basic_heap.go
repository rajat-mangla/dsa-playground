package arrays

func HeapSort(data []int) []int {
	length := len(data)
	BuildMaxHeap(data)

	for i := length - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapify(data, i, 0)
	}

	return data
}

func BuildMaxHeap(data []int) []int {
	length := len(data)
	for i := length/2 - 1; i >= 0; i-- {
		heapify(data, length, i)
	}

	return data
}

func heapify(data []int, length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && data[left] > data[largest] {
		largest = left
	}

	if right < length && data[right] > data[largest] {
		largest = right
	}

	if largest != i {
		data[i], data[largest] = data[largest], data[i]
		heapify(data, length, largest)
	}
}
