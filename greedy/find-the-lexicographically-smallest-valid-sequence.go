package greedy

// https://leetcode.com/problems/find-the-lexicographically-smallest-valid-sequence
func validSequence(word1 string, word2 string) []int {
	n1 := len(word1)
	n2 := len(word2)

	last := make([]int, n2)
	for i := range n2 {
		last[i] = -1
	}

	i := n1 - 1
	j := n2 - 1
	for i >= 0 && j >= 0 {
		if word1[i] == word2[j] {
			last[j] = i
			j--
		}

		i--
	}

	ans := make([]int, n2)
	canSkip := true

	i = 0
	j = 0
	for i < n1 && j < n2 {
		if word1[i] == word2[j] {
			ans[j] = i
			j++
			i++
			continue
		}

		if canSkip && (j == n2-1 || i < last[j+1]) {
			canSkip = false
			ans[j] = i
			j++
		}
		i++
	}

	if j == n2 {
		return ans
	}

	return []int{}
}
