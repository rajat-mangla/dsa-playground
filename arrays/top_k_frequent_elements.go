package arrays

//------------------- Bucket Sort Approach ------------------//

// TopKFrequentUsingBucketSort - Bucketing sort using frequency map
// Time Complexity: O(N)
// Space Complexity: O(N)
func TopKFrequentUsingBucketSort(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range frequencyMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}

	if len(result) > k {
		return result[:k]
	}
	return result
}

// ------------------- Heap Sort Approach ------------------//
type element struct {
	num  int
	freq int
}

// TopKFrequentUsingMinHeap - Min-Heap sort using frequency map
// Time Complexity: O(N log N)
// Space Complexity: O(N)
func TopKFrequentUsingMinHeap(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	heapArr := make([]element, 0, len(frequencyMap))
	for num, freq := range frequencyMap {
		heapArr = append(heapArr, element{num, freq})
	}

	ans := heapSort(heapArr)
	return ans[0:k]
}

func heapSort(data []element) []int {
	length := len(data)
	for i := length/2 - 1; i >= 0; i-- {
		minHeapify(data, i, length)
	}

	for i := length - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		minHeapify(data, 0, i)
	}

	result := make([]int, length)
	for i, elem := range data {
		result[i] = elem.num
	}
	return result
}

func minHeapify(data []element, i, length int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && data[left].freq < data[smallest].freq {
		smallest = left
	}

	if right < length && data[right].freq < data[smallest].freq {
		smallest = right
	}

	if smallest != i {
		data[i], data[smallest] = data[smallest], data[i]
		minHeapify(data, smallest, length)
	}
}

// TopKFrequentUsingMinHeapV2 - Min-Heap sort using frequency map (optimized)
// Time Complexity: O(N log k)
// Space Complexity: O(N)
func TopKFrequentUsingMinHeapV2(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	heapArr := make([]element, 0, k)
	for num, freq := range frequencyMap {
		if len(heapArr) < k {
			heapArr = append(heapArr, element{num, freq})
			if len(heapArr) == k {
				// build min-heap when we have k elements
				for i := k/2 - 1; i >= 0; i-- {
					minHeapify(heapArr, i, k)
				}
			}
			continue
		}

		if freq > heapArr[0].freq {
			heapArr[0] = element{num, freq}
			minHeapify(heapArr, 0, len(heapArr)) // heapify from root
		}
	}

	ans := make([]int, k)
	for i, elem := range heapArr {
		ans[i] = elem.num
	}
	return ans
}

// ------------------- Quick Select Approach ------------------//
// TopKFrequentUsingQuickSelect - Quick Select using frequency map
// Time Complexity: O(N) on average, O(N^2) in the worst case
// Space Complexity: O(N)
func TopKFrequentUsingQuickSelect(nums []int, k int) {
}
