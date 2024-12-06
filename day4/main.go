package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var file string

func init() {
	flag.StringVar(&file, "file", "example.txt", "File from which to read the input data")
}

// part 2
func main() {
	flag.Parse()
	file, _ := os.Open(file)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	count := 0
	for r := range grid {
		for c := range grid[0] {
			count += search(grid, r, c)
		}
	}
	fmt.Println(count)
}

func search(grid []string, r, c int) int {
	if r-1 < 0 || c-1 < 0 || r+1 >= len(grid) || c+1 >= len(grid[0]) {
		return 0
	}
	if grid[r][c] != 'A' {
		return 0
	}

	if (grid[r-1][c-1] != 'M' || grid[r+1][c+1] != 'S') && (grid[r-1][c-1] != 'S' || grid[r+1][c+1] != 'M') {
		return 0
	}

	if (grid[r+1][c-1] != 'M' || grid[r-1][c+1] != 'S') && (grid[r+1][c-1] != 'S' || grid[r-1][c+1] != 'M') {
		return 0
	}
	return 1
}

// part 1
// func search(grid []string, r, c, rDiff, cDiff int) int {
// 	for i := range "XMAS" {
// 		newR := r + rDiff*i
// 		newC := c + cDiff*i
// 		if newR < 0 || newC < 0 || newR >= len(grid) || newC >= len(grid[0]) || grid[newR][newC] != "XMAS"[i] {
// 			return 0
// 		}
// 	}
// 	return 1
// }

// func main() {
// 	flag.Parse()
// 	file, _ := os.Open(file)
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
// 	grid := make([]string, 0, 100)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		grid = append(grid, line)
// 	}
//
// 	count := 0
// 	DIFFS := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
// 	for r := range grid {
// 		for c := range grid[0] {
// 			for _, diff := range DIFFS {
// 				count += search(grid, r, c, diff[0], diff[1])
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }
