package stack_and_queue

func nextGreaterElement(n int) int {
	nums := []int{}
	for n/10 != 0 || n%10 != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}

	idx := findPivot(nums)
	if idx == -1 {
		return -1
	}

	ans := 0
	multiplier := 1
	for i := range nums {
		ans += nums[i] * multiplier
		multiplier = multiplier * 10
	}

	if ans > 2147483647 {
		return -1
	}
	return ans
}

func findPivot(arr []int) int {
	l := len(arr)
	pivot := -1
	for i := 0; i < l-1; i++ {
		if arr[i] > arr[i+1] {
			pivot = i + 1
			break
		}
	}

	if pivot > 0 {
		for i := range pivot {
			if arr[i] > arr[pivot] {
				arr[i], arr[pivot] = arr[pivot], arr[i]
				break
			}
		}

		tempArr := append([]int{}, arr[:pivot]...)
		for i := range tempArr {
			arr[i] = tempArr[pivot-1-i]
		}
	}

	return pivot
}
