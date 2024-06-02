package lc

// Create a hash map and store the seen values with a empty value, that way, if it's found in the hashmap, it's a duplicate
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	if len(nums) <= 1 {
		return false
	}
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
