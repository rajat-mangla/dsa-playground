package recursion

// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	}

	if k < mid {
		return findKthBit(n-1, k)
	}

	idx := 1<<n - k
	return invert(findKthBit(n-1, idx))
}

func invert(x byte) byte {
	if x == '0' {
		return '1'
	}

	return '0'
}
