package solutions

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sc := make(map[rune]int, len(s))
	tc := make(map[rune]int, len(t))

	for _, ch := range s {
		sc[ch]++
	}

	for _, ch := range t {
		tc[ch]++
	}

	for k1, v1 := range sc {
		if v2 := tc[k1]; v2 != v1 {
			return false
		}
	}

	return true
}
