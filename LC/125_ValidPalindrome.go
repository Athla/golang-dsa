package lc

import "unicode"

func isPalindrome(s string) bool {
	end := len(s) - 1
	start := 0
	r_string := []rune(s)
	for start <= end {
		bot := r_string[start]
		top := r_string[end]
		if unicode.IsLetter(bot) && unicode.IsLetter(top) && !unicode.IsDigit(bot) && !unicode.IsDigit(top) {
			if bot == top {
				start++
				end--
				continue
			} else {
				return false
			}
		}
	}

	return true
}
