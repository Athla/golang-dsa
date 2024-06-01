# Sort the array
# Check the combinations
# When arriving into a > 0 val as starting value, end and return what was find

from typing import List


def ThreeSum(self, nums: List[int]) -> List[List[int]]:
	nums.sort()
	res = []
	for i, a in enumerate(nums):
		if i > 0 and a == nums[i - 1]:
			continue
		l, r = i + 1, len(nums) - 1
		while l < r:
			threeSum = a + nums[l] + nums[r]
			if threeSum > 0: 
				r -= 1
			elif threeSum < 0:
				l += 1
			else:
				res.append([a, nums[l], nums[r]])
				l += 1
				while nums[l] == nums[l - 1] and l < r:
					l += 1 
	return res


