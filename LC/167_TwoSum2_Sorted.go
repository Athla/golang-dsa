package lc

// Ordered version of twosum, it's simpler to do because the two pointers existence
// Always check if the sum is bellow or above the target, and move pointers accordingly
func twoSum(numbers []int, target int) []int {
	top := len(numbers) - 1
	bot := 0
	sum := 0
	for bot < top {
		sum = numbers[bot] + numbers[top]
		if sum == target {
			return []int{bot, top}
		} else if sum > target {
			bot++
		} else if sum < target {
			top--
		}
	}

	return []int{}
}
