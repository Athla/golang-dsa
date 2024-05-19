package main

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Point struct {
	x int
	y int
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// Base cases
	// 1. Off the Map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y <= 0 || curr.y > len(maze) {
		return false
	}
	// On a wall
	if maze[curr.y][curr.x] == wall[0] {
		return false
	}
	// Is the End
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}
	// Have been seen
	if seen[curr.y][curr.x] {
		return false
	}
	// 3 steps
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)
	// recurse
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]
		next := Point{x: curr.x + x, y: curr.y + y}
		if walk(maze, wall, next, end, seen, path) {
			return true
		}
	}
	// post
	(*path) = (*path)[:len(*path)-1]
	return false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point
	walk(maze, wall, start, end, seen, &path)

	return path
}
