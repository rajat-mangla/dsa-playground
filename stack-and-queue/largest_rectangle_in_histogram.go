package stack_and_queue

func largestRectangleArea(heights []int) int {
	maxArea := 0

	pse := prevSmallerElemIdxArr(heights)
	nse := nextSmallerElemIdxArr(heights)

	for i := range heights {
		length := nse[i] - pse[i] - 1
		maxArea = max(maxArea, heights[i]*length)
	}

	return maxArea
}

func prevSmallerElemIdxArr(h []int) []int {
	stack := []int{}
	pse := make([]int, len(h))
	for i := range h {

		for len(stack) > 0 && h[stack[len(stack)-1]] >= h[i] {
			stack = stack[:len(stack)-1]
		}

		idx := -1
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}
		stack = append(stack, i)
		pse[i] = idx
	}

	return pse
}

func nextSmallerElemIdxArr(h []int) []int {
	l := len(h)
	stack := []int{}
	nse := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && h[stack[len(stack)-1]] >= h[i] {
			stack = stack[:len(stack)-1]
		}

		idx := l
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}
		stack = append(stack, i)
		nse[i] = idx
	}

	return nse
}
