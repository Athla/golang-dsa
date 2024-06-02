package lc

// TopKFrequent finds the k most frequent elements in nums.
// It uses a map to count the frequency of each element and a slice of slices to group elements by frequency.
// Finally, it constructs the result by iterating from the highest frequency to the lowest until k elements are collected.

func TopKFrequent(nums []int, k int) (res []int) {
	countFrequency := map[int]int{}
	countSlice := make([][]int, len(nums)+1)

	for _, n := range nums {
		if count, ok := countFrequency[n]; ok {
			countFrequency[n] = count + 1
		} else {
			countFrequency[n] = 1
		}

	}

	for num, count := range countFrequency {
		countSlice[count] = append(countSlice[count], num)
	}

	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return res
		}

	}

	return res
}
