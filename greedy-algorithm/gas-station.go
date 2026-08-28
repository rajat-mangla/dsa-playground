package greedy_algorithm

// https://leetcode.com/problems/gas-station/description/
func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	curr := 0

	n := len(gas)

	ans := 0

	for i := range n {
		diff := gas[i] - cost[i]

		total += diff
		curr += diff
		if curr < 0 {
			ans = i + 1
			curr = 0
		}
	}

	if total < 0 {
		return -1
	}

	return ans
}

// [1,2,3,4,5],
// [3,4,5,1,2]

// [-2,-2,-2,3,3]

// [2,3,4],
// [3,4,3]

// [-1, -1, 1, -1, 1, 1]
