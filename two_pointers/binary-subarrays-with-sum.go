package two_pointers

// Intuition:
// We can think of this problem like this:
// Each subarray sum is the difference between two prefix sums.
// If we know the number of prefix sums that give us (current_sum - goal),
// we can count how many subarrays end at the current index that have the required sum.
func numSubarraysWithSum(nums []int, goal int) int {
	prefixCount := map[int]int{0: 1}
	prefixSum, result := 0, 0

	for _, num := range nums {
		prefixSum += num
		result += prefixCount[prefixSum-goal]
		prefixCount[prefixSum]++
	}

	return result
}

//func numSubarraysWithSum(nums []int, goal int) int {
//	return cntAtMostSubarraysFor(nums, goal) - cntAtMostSubarraysFor(nums, goal-1)
//}

func cntAtMostSubarraysFor(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}
	length := len(nums)
	r := 0
	l := 0

	cnt := 0
	numSubArrays := 0
	for r < length {
		cnt += nums[r]

		for cnt > goal && l < length {
			cnt -= nums[l]
			l++
		}

		numSubArrays += r - l + 1
		r++
	}

	return numSubArrays
}
