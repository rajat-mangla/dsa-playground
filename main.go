package main

import (
	"fmt"

	"github.com/rajat-mangla/dsa-playground/arrays"
)

func main() {

	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println(arrays.BuildMaxHeap(arr))

	fmt.Println(arrays.HeapSort(arr))
}

// [9 6 4 5 5 3 2 1 1 3 5]
// Output: [9 6 5 5 5 4 2 1 3 1 3]
// Explanation: The output is a max-heap representation of the input array.
