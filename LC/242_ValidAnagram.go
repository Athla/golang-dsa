package lc

import (
	"slices"
	"strings"
)

//sort the arrays and compare if they're the same

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

// Make a slice of ints that has a size of 26 (one for each english letter)
// iterate over word one, get the ascii values of it as a int and save as a idx, add one to this idx
// do the same for the other one, but decrement the value of the current idx
// if all the values in the slice are 0, it's anagram, else it's not
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
