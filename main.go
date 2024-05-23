package main

import (
	"fmt"
	// lc "github.com/Athla/golang-dsa/LC"
	algo "github.com/Athla/golang-dsa/Algorithms"
)

func main() {
	arr := []int{5, 6, 7, 2, 1, 6, 7, 8, 0}
	fmt.Println(arr)
	algo.QuickSort(arr)
	fmt.Println(arr)
}
