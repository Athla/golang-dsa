package lc

func TwoSum(nums []int, target int) []int {
	// map of val, idx
	diffmap := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if val, found := diffmap[diff]; found {
			return []int{val, i}
		}
		diffmap[v] = i
	}

	return []int{}
}
