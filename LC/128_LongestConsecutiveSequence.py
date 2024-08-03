from typing import List


def longestConsecutive(nums: List[int]) -> int:
    setNums = set(nums)
    long = 0
    for n in setNums:
        if (n - 1) not in setNums:
            l = 0
            while n + l in setNums:
                l += 1
            long = max(l, long)

    return long
