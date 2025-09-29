package sorting

// MergeSort sorts an array of integers using the merge sort algorithm.
func MergeSort(arr []int) []int {
	mergeSort(arr, 0, len(arr)-1)
	return arr
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := (right + left) / 2

		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	l := left
	r := mid + 1
	tempArr := make([]int, right-left+1)
	k := 0

	for l <= mid && r <= right {
		if arr[l] <= arr[r] {
			tempArr[k] = arr[l]
			l++
		} else {
			tempArr[k] = arr[r]
			r++
		}
		k++
	}

	for l <= mid {
		tempArr[k] = arr[l]
		l++
		k++
	}

	for r <= right {
		tempArr[k] = arr[r]
		r++
		k++
	}

	copy(arr[left:right+1], tempArr)
}
