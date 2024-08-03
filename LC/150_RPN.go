package lc

import (
	"fmt"
	"strconv"
)

func EvalRPN(tokens []string) int {
	op := map[string]string{"+": "", "-": "", "*": "", "/": ""}
	var stack []int
	var res int

	for _, t := range tokens {
		_, ok := op[t]
		if !ok {
			i, _ := strconv.Atoi(t)
			stack = append(stack, i)
		} else {
			a := stack[stack[len(stack)-1]]
			b := stack[stack[len(stack)-2]]
			stack = stack[:stack[len(stack)-3]]
			switch t {
			case "+":
				res += a + b
			case "-":
				res += b - a
			case "*":
				res += a * b
			case "/":
				res += b / a
			default:
				fmt.Println("Invalid op")
			}

		}

	}

	return stack[len(stack)-1]
}
