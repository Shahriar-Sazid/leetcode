package solutions

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)
	mi := (lo + hi) / 2
	for lo < hi {
		if target == nums[mi] {
			return mi
		}
		if target > nums[mi] {
			lo = mi + 1
			mi = (lo + hi) / 2
		} else {
			hi = mi
			mi = (lo + hi) / 2
		}
	}

	return -1
}

func search(nums []int, target int) int {
	pivotIndex := -1
	for i := 0; i < len(nums)-1; i++ {
		prev := nums[i]
		cur := nums[i+1]

		if prev <= cur {
			continue
		} else {
			pivotIndex = i + 1
		}
	}

	if pivotIndex == -1 {
		return binarySearch(nums, target)
	}

	idx := binarySearch(nums[:pivotIndex], target)
	if idx != -1 {
		return idx
	}

	idx = binarySearch(nums[pivotIndex:], target)
	if idx != -1 {
		return pivotIndex + idx
	}

	return -1
}
