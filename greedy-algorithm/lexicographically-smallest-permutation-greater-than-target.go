package greedy_algorithm

// https://leetcode.com/problems/lexicographically-smallest-permutation-greater-than-target
func lexGreaterPermutation(s string, target string) string {
	cntStore := [26]int{}

	for i := range s {
		cntStore[int(s[i]-'a')]++
	}

	n := len(target)
	ans := []byte{}

	for i := range n {
		char := int(target[i] - 'a')

		if cntStore[char] > 0 {
			cntStore[char]--
			if canFormGreaterString(cntStore, target[i+1:]) {
				ans = append(ans, target[i])
				continue
			}
			cntStore[char]++
		}

		char++
		for char < 26 {
			if cntStore[char] > 0 {
				cntStore[char]--
				ans = append(ans, byte(char+'a'))
				ans = append(ans, getMinString(cntStore)...)
				return string(ans)
			}

			char++
		}

		return ""
	}

	return string(ans)
}

func canFormGreaterString(cnt [26]int, target string) bool {
	return getMaxString(cnt) > target
}

func getMaxString(cnt [26]int) string {
	byteArr := []byte{}
	for i := 25; i >= 0; i-- {
		c := cnt[i]
		for _ = range c {
			byteArr = append(byteArr, byte(i+'a'))
		}
	}

	return string(byteArr)
}

func getMinString(cnt [26]int) []byte {
	byteArr := []byte{}
	for i := range 26 {
		c := cnt[i]
		for _ = range c {
			byteArr = append(byteArr, byte(i+'a'))
		}
	}

	return byteArr
}
