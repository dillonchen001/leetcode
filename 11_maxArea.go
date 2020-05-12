package main

func maxArea(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}

	start := 0
	end := length - 1
	maxArea := 0
	ll := 0
	hi := 0
	for start < end {
		ll = end - start
		hi = height[start]

		if hi > height[end] {
			hi = height[end]
			end = end - 1
		} else {
			start = start + 1
		}

		if ll*hi > maxArea {
			maxArea = ll * hi
		}
	}
	return maxArea
}
