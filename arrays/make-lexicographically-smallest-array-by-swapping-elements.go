package arrays

import (
	"sort"
)

// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements
type group struct {
	numArr []int
	idxArr []int
}

type pair struct {
	val int
	idx int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	pairArr := make([]pair, n)
	for i := range n {
		pairArr[i] = pair{nums[i], i}
	}

	sort.Slice(pairArr, func(i, j int) bool {
		return pairArr[i].val < pairArr[j].val
	})

	var groups []group
	currG := group{[]int{pairArr[0].val}, []int{pairArr[0].idx}}
	for i := 1; i < n; i++ {
		diff := pairArr[i].val - pairArr[i-1].val
		if diff <= limit {
			currG.numArr = append(currG.numArr, pairArr[i].val)
			currG.idxArr = append(currG.idxArr, pairArr[i].idx)
			continue
		}

		groups = append(groups, currG)
		currG = group{[]int{pairArr[i].val}, []int{pairArr[i].idx}}
	}
	groups = append(groups, currG)

	for _, g := range groups {
		sort.Ints(g.idxArr)
		sort.Ints(g.numArr)

		for i := range len(g.idxArr) {
			nums[g.idxArr[i]] = g.numArr[i]
		}
	}

	return nums
}
