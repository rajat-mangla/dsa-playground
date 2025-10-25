package two_pointers

// Key Idea — “Exactly K” = “At Most K” − “At Most (K−1)”
// This is the same conceptual trick
//we used for the “nice subarrays” and “binary subarrays with sum” problems, but generalized.

func subarraysWithKDistinct(nums []int, k int) int {
	return cntAtMostKDistinctSubArrays(nums, k) - cntAtMostKDistinctSubArrays(nums, k-1)
}

func cntAtMostKDistinctSubArrays(nums []int, goal int) int {
	length := len(nums)
	r := 0
	l := 0

	mapStore := map[int]int{}
	numSubArrays := 0
	for r < length {
		mapStore[nums[r]]++

		for len(mapStore) > goal && l < length {
			mapStore[nums[l]]--
			if mapStore[nums[l]] == 0 {
				delete(mapStore, nums[l])
			}
			l++
		}

		numSubArrays += r - l + 1
		r++
	}

	return numSubArrays
}
