package two_pointers

// the “Count Number of Nice Subarrays” problem is indeed a close cousin of “Binary Subarrays With Sum.”
//
//Key Observation:
//
//This is identical in structure to the “Binary Subarrays With Sum” problem —
//if we map each element to whether it’s odd (1) or even (0).

// Now, the problem becomes:
//
//Count the number of subarrays in arr whose sum = k.
//
//…and that’s exactly what we did for “Binary Subarrays With Sum.”

func numberOfSubarrays(nums []int, goal int) int {
	prefixCount := map[int]int{0: 1}
	prefixSum, result := 0, 0

	for _, num := range nums {
		if num%2 != 0 {
			prefixSum += 1
		}
		result += prefixCount[prefixSum-goal]
		prefixCount[prefixSum]++
	}

	return result
}

//func numberOfSubarrays(nums []int, k int) int {
//	return cntAtMostSubArrays(nums, k) - cntAtMostSubArrays(nums, k-1)
//}

func cntAtMostSubArrays(nums []int, k int) int {
	length := len(nums)
	var l, r int

	oddCnt := 0
	cntSubArrs := 0
	for r < length {
		if nums[r]%2 != 0 {
			oddCnt++
		}

		for oddCnt > k && l < length {
			if nums[l]%2 != 0 {
				oddCnt--
			}
			l++
		}

		cntSubArrs += r - l + 1
		r++
	}

	return cntSubArrs
}
