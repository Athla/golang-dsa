package lc

import "sort"

// After picking the first number, implment ordered two sum, using two pointers to find the solution.
// It's a O(N^2) solution in worst case scenario
func Threesum(nums []int) (res [][]int) {
	sort.Ints(nums)
	var sum int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum = v + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{v, nums[l], nums[r]})
				r--
				for nums[l] == nums[l-1] && l < r {
					r--
				}
			}

		}
	}
	return res
}
