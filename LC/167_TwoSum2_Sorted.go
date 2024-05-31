package lc

func twoSum(numbers []int, target int) []int {
	top := len(numbers) - 1
	bot := 0

	for bot < top {
		if numbers[bot]+numbers[top] == target {
			return []int{bot, top}
		} else if numbers[bot] > target {
			bot++
		} else if numbers[top] < target {
			top--
		}
	}

	return []int{}
}
