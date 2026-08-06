package solutions

func lengthOfLongestSubstring(s string) int {
	var last [128]int // stores index+1; 0 = unseen
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if last[c] > l {
			l = last[c]
		}
		last[c] = r + 1
		best = max(best, r-l+1)
	}
	return best
}
