package solutions

func isValid(s string) bool {
	pairs := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	runeStack := make([]rune, 0, len([]rune(s)))
	for _, r := range s {
		if p, ok := pairs[r]; ok {
			var last rune
			if len(runeStack) > 0 {
				last = runeStack[len(runeStack)-1]
			}
			if p == last {
				runeStack = runeStack[:len(runeStack)-1]
				continue
			}
		}
		runeStack = append(runeStack, r)
	}

	return len(runeStack) == 0
}
