package main

import (
	"fmt"
)

func main() {
	maze := []string{
		"#######",
		"#     #",
		"# ### #",
		"# #   #",
		"# ### #",
		"#     #",
		"#######",
	}
	wall := "#"
	start := Point{1, 1}
	end := Point{5, 5}
	path := Solve(maze, wall, start, end)
	fmt.Println("Path:", path)
}
