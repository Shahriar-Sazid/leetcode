package solutions

import (
	"math"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	length := int(math.Floor(math.Log10(float64(x)))) + 1
	for i := 0; i < length/2; i++ {
		d := (int(math.Pow10(length - i - 1)))
		lDigit := y / d
		y = y % d

		rDigit := x % 10
		x = x / 10

		if lDigit != rDigit {
			return false
		}
	}
	return true
}
