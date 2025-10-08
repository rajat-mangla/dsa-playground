package stack_and_queue

func subArrayRanges(nums []int) int64 {
	return sumSubArrayMax(nums) - sumSubarrayMins(nums)
}
