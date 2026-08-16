package greedy_algorithm

import (
	"math"
)

// https://leetcode.com/problems/stone-game-ix
func stoneGameIX(stones []int) bool {
	cnt := [3]int{}
	n := len(stones)
	for i := range n {
		cnt[stones[i]%3] += 1
	}

	if cnt[0]%2 == 0 {
		return cnt[1] >= 1 && cnt[2] >= 1
	}

	return math.Abs(float64(cnt[2]-cnt[1])) > 2
}
