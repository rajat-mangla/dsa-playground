package two_pointers

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	targetCount := map[byte]int{}
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	left, right := 0, 0
	minLeft, minRight := 0, len(s)+1
	windowCount := map[byte]int{}
	required := len(targetCount)
	formed := 0

	for right < len(s) {
		char := s[right]
		windowCount[char]++

		if count, exists := targetCount[char]; exists && windowCount[char] == count {
			formed++
		}

		for left <= right && formed == required {
			char = s[left]

			if right-left+1 < minRight-minLeft {
				minLeft = left
				minRight = right + 1
			}

			windowCount[char]--
			if count, exists := targetCount[char]; exists && windowCount[char] < count {
				formed--
			}
			left++
		}
		right++
	}

	if minRight == len(s)+1 {
		return ""
	}
	return s[minLeft:minRight]
}
