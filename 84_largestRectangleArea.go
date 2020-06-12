package main

func largestRectangleArea(heights []int) int {
	lh := len(heights)

	if lh == 0 {
		return 0
	}

	result := 0
	var stack []int

	stack = append(stack, heights[0])

	for i := 1; i < lh; i++ {
		ls := len(stack)

		if heights[i] < stack[ls-1] {
			for j := ls - 1; j >= 0; j-- {
				if stack[j] <= heights[i] {
					break
				}

				if stack[j]*(ls-j) > result {
					result = stack[j] * (ls - j)
				}

				stack[j] = heights[i]
			}
		}

		stack = append(stack, heights[i])
	}

	lsr := len(stack)
	for j := 1; j <= lsr; j++ {
		if stack[lsr-j]*j > result {
			result = stack[lsr-j] * j
		}
	}
	return result
}
