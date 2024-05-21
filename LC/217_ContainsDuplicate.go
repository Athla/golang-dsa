package lc

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
