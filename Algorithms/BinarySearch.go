package algorithms

func BinarySearch(nums []int, target int) int {
	top := len(nums)
	bot := 0
	var mid int = (top + bot) / 2
	for bot < top {
		midPoint := nums[mid]
		if midPoint < target {
			bot = mid
			mid = (top + bot) / 2
		} else if midPoint > target {
			top = mid
			mid = (top + bot) / 2
		} else {
			return mid
		}
	}

	return -1
}
