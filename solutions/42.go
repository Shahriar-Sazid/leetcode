package solutions

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	total := 0

	for l < r {
		if leftMax <= rightMax {
			l++
			if height[l] > leftMax {
				leftMax = height[l]
			} else {
				total += leftMax - height[l]
			}
		} else {
			r--
			if height[r] > rightMax {
				rightMax = height[r]
			} else {
				total += rightMax - height[r]
			}
		}
	}

	return total
}
