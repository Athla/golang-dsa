package main

import (
	"fmt"

	lc "github.com/Athla/golang-dsa/LC"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(lc.TwoSum(nums, target))

}
