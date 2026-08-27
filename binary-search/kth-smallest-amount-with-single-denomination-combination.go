package binary_search

// https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}

	return (a / gcd(a, b)) * b
}

type subsetInfo struct {
	lcm  int64
	sign int64
}

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)

	totalSubsets := 1 << n
	subsets := make([]subsetInfo, totalSubsets-1)

	mask := 1
	for mask < totalSubsets {
		var currentLCM int64 = 1
		size := 0

		for i := range n {
			if (mask>>i)&1 == 1 {
				currentLCM = lcm(currentLCM, int64(coins[i]))
				size++
			}
		}

		s := subsetInfo{currentLCM, 1}
		if size%2 == 0 {
			s.sign = -1
		}
		subsets[mask-1] = s
		mask++
	}

	countLessEqual := func(x int64) int64 {
		var total int64
		for _, s := range subsets {
			total += s.sign * (x / s.lcm)
		}
		return total
	}

	minCoin := coins[0]
	for i := range coins {
		minCoin = min(minCoin, coins[i])
	}

	l := int64(1)
	r := int64(minCoin * k)
	ans := r
	for l <= r {
		mid := l + (r-l)/2

		if countLessEqual(mid) >= int64(k) {
			ans = mid
			r = mid - 1
			continue
		}

		l = mid + 1
	}
	return ans
}
