package solutions

func twoSum(nums []int, target int) []int {
	complementIndex := make(map[int]int, len(nums))

	for i, n := range nums {
		c := target - n
		if ci, ok := complementIndex[c]; ok {
			return []int{ci, i}
		}

		complementIndex[n] = i
	}

	return []int{-1, -1}
}
