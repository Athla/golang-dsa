package lc

// Use a hashmap to keep the track of the val and idx, check if the difference between current value and target is in the diffmap
// if it is, return the current diff idx (val) and the found idx,
// else add the current value (v) to the hashmap with the value as idx
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
