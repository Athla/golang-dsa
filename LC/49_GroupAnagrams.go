package lc

import (
	"slices"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	sortString := func(s string) string {
		splited := strings.Split(s, "")
		slices.Sort(splited)
		return strings.Join(splited, "")
	}

	for _, v := range strs {
		curr := sortString(v)
		groups[curr] = append(groups[curr], v)
	}
	ans := make([][]string, len(groups))
	i := 0
	for k := range groups {
		ans[i] = append(ans[i], groups[k]...)
		i++
	}

	return ans
}
