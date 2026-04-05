package binary_search

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) >> 1

		if nums[mid] < target {
			left = mid + 1
			continue
		}

		right = mid
	}

	return left
}

func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) >> 1

		if nums[mid] <= target {
			left = mid + 1
			continue
		}

		right = mid
	}

	return left
}

func searchRange(nums []int, target int) []int {
	lower := lowerBound(nums, target)
	upper := upperBound(nums, target)

	if lower == upper {
		return []int{-1, -1}
	}

	return []int{lower, upper - 1}
}
