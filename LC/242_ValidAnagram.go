package lc

import (
	"slices"
	"strings"
)

func IsAnagram(s string, t string) bool {
	s_slice := strings.Split(s, "")
	t_slice := strings.Split(t, "")
	slices.Sort(s_slice)
	slices.Sort(t_slice)

	if slices.Equal(s_slice, t_slice) {
		return true
	}

	return false
}

func IsAnagram2(s string, t string) bool {
	chars := make([]int, 26)

	for _, c := range s {
		i := int(c - 'a')
		chars[i]++
	}

	for _, c := range t {
		i := int(c - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
