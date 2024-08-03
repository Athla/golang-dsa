def BinarySearch(nums: list[int], target: int) -> int:
    top = len(nums)
    bot = 0
    mid = (top + bot) // 2
    while bot < top:
        midPoint = nums[mid]
        if midPoint < target:
            bot = mid
            mid = (top + bot) // 2
        elif midPoint > target:
            top = mid
            mid = (top + bot) // 2
        else:
            return mid

    return -1


def search(self, nums: List[int], target: int) -> int:
    bot, top = 0, len(nums) - 1
    while bot <= top:
        m = (bot + top) // 2
        if nums[m] > target:
            top = m - 1
        elif nums[m] < target:
            bot = m + 1
        else:
            return m
    return -1
