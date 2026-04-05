package recursion

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/special-binary-string/description
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}

	result := []string{}
	count := 0
	start := 0

	for i := range s {
		if s[i] == '1' {
			count += 1
		}

		if s[i] == '0' {
			count -= 1
		}

		if count == 0 {
			middle := makeLargestSpecial(s[start+1 : i])
			result = append(result, "1"+middle+"0")
			start = i + 1
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return strings.Join(result, "")
}
