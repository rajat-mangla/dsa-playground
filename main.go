package main

import (
	"fmt"

	"github.com/rajat-mangla/dsa-playground/sorting"
)

func main() {

	arr := []int{-2, 45, 0, 11, -9, 88, -97, -202, 747}
	//arr = []int{12, 11, 13, 5, 6, 7}
	fmt.Println(sorting.QuickSort(arr))
	//fmt.Println(arrays.BuildMaxHeap(arr))
	//fmt.Println(arrays.HeapSort(arr))
}

// [9 6 4 5 5 3 2 1 1 3 5]
// Output: [9 6 5 5 5 4 2 1 3 1 3]
// Explanation: The output is a max-heap representation of the input array.
