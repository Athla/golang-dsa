package lc

func isValid(s string) bool {
	valid := map[byte]byte{'}': '{', ']': '[', ')': '('}

	stack := []byte{}

	for _, c := range []byte(s) {
		pair, ok := valid[c]
		if !ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return true
		}

		if stack[len(stack)-1] != pair {
			return false
		}

		stack = stack[:len(stack)-1]

	}

	if len(stack) == 0 {
		return true
	}

	return false

}
