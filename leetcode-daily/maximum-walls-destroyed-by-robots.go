package leetcode_daily

import (
	"sort"
)

// date: 2026-04-03
// https://leetcode.com/problems/maximum-walls-destroyed-by-robots
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type robot struct {
		dis, pos int
	}

	robotArr := make([]robot, n)
	for i := range n {
		robotArr[i] = robot{distance[i], robots[i]}
	}

	sort.Slice(robotArr, func(i, j int) bool {
		return robotArr[i].pos < robotArr[j].pos
	})

	sort.Slice(walls, func(i, j int) bool {
		return walls[i] < walls[j]
	})

	leftCnt := make([]int, n)
	rightCnt := make([]int, n)

	for i := range n {
		l := robotArr[i].pos - robotArr[i].dis
		r := robotArr[i].pos + robotArr[i].dis

		if i > 0 {
			l = max(l, robotArr[i-1].pos+1)
		}

		if i < n-1 {
			r = min(r, robotArr[i+1].pos-1)
		}

		leftCnt[i] = count(walls, l, robotArr[i].pos)
		rightCnt[i] = count(walls, robotArr[i].pos, r)
	}

	dp := make([][2]int, n)
	// left hit
	dp[0][0] = leftCnt[0]
	// right hit
	dp[0][1] = rightCnt[0]

	for i := 1; i < n; i++ {
		prevR := min(robotArr[i-1].pos+robotArr[i-1].dis, robotArr[i].pos-1)
		currL := max(robotArr[i].pos-robotArr[i].dis, robotArr[i-1].pos+1)

		overlap := 0
		if prevR >= currL {
			overlap = count(walls, currL, prevR)
		}

		// left hit can overlap with previous right hit, so we need to subtract the overlap
		dp[i][0] = max(dp[i-1][0]+leftCnt[i], dp[i-1][1]+leftCnt[i]-overlap)
		// right hit will not overlap with previous robot
		dp[i][1] = max(dp[i-1][1]+rightCnt[i], dp[i-1][0]+rightCnt[i])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func count(walls []int, l, r int) int {
	return upperBound(walls, r) - lowerBound(walls, l)
}

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
