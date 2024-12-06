package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var count atomic.Uint64

func main() {
	t := time.Now()
	defer func() {
		fmt.Println("Time consumed", time.Since(t))
	}()

	f, _ := os.Open("input.txt")
	// f, _ := os.Open("example.txt")
	scanner := bufio.NewScanner(f)

	board := make([][]byte, 0, 100)
	for scanner.Scan() {
		text := scanner.Text()
		board = append(board, []byte(text))
	}

	blockCandidates, _ := traverse(board, Pos{-1, -1})
	fmt.Println("Block blockCandidates", len(blockCandidates))

	// for candidate := range blockCandidates {
	// 	testSingleThreaded(board, candidate)
	// }

	var wg sync.WaitGroup
	for candidate := range blockCandidates {
		wg.Add(1)
		go copyAndTest(board, candidate, &wg)
	}

	wg.Wait()
	fmt.Println("Loop count", count.Load())
}

func testSingleThreaded(board [][]byte, candidate Pos) {
	if _, loop := traverse(board, candidate); loop {
		count.Add(1)
		// fmt.Println("Loop if block", candidate)
	}
}

func copyAndTest(board [][]byte, candidate Pos, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, loop := traverse(board, candidate); loop {
		count.Add(1)
		// fmt.Println("Loop if block", candidate)
	}
}

type Pos struct {
	r, c int
}

func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func countVisited(board [][]byte) int {
	count := 0
	for _, row := range board {
		for _, b := range row {
			if b == 'X' {
				count++
			}
		}
	}
	return count
}

func findStartPos(grid [][]byte) (Pos, bool) {
	for r, row := range grid {
		for c := range row {
			if row[c] == '^' {
				return Pos{
					r, c,
				}, true
			}
		}
	}

	return Pos{}, false
}

var DIFFS = []Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func traverse(grid [][]byte, banned Pos) (map[Pos]struct{}, bool) {
	memo := make([][]int, len(grid))
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
	}

	blockSet := make(map[Pos]struct{}, 1000)
	pos, ok := findStartPos(grid)
	if !ok {
		return nil, false
	}

	d := 0
	for pos.r >= 0 && pos.r < len(grid) && pos.c >= 0 && pos.c < len(grid[0]) {
		if ((memo[pos.r][pos.c] >> d) & 1) > 0 {
			return nil, true
		}
		memo[pos.r][pos.c] |= (1 << d)
		diff := DIFFS[d]

		nextPos := Pos{pos.r + diff.r, pos.c + diff.c}
		if nextPos.r >= 0 && nextPos.r < len(grid) && nextPos.c >= 0 && nextPos.c < len(grid[0]) {
			if grid[nextPos.r][nextPos.c] == '#' || banned == nextPos {
				d = (d + 1) % 4
				continue
			} else {
				blockSet[nextPos] = struct{}{}
			}
		}

		pos = nextPos
	}

	return blockSet, false
}

// part 1
// func main() {
// 	f, _ := os.Open("input.txt")
// 	// f, _ := os.Open("example.txt")
// 	scanner := bufio.NewScanner(f)
//
// 	board := make([][]byte, 0, 100)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		board = append(board, []byte(text))
// 	}
//
// 	fmt.Println("Before")
// 	printBoard(board)
// 	traverse(board)
// 	fmt.Println("After")
// 	printBoard(board)
//
// 	fmt.Println("Visited: ", countVisited(board))
// }
//
// type Pos struct {
// 	r, c int
// }
//
// func printBoard(board [][]byte) {
// 	for _, row := range board {
// 		fmt.Println(string(row))
// 	}
// }
//
// func countVisited(board [][]byte) int {
// 	count := 0
// 	for _, row := range board {
// 		for _, b := range row {
// 			if b == 'X' {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
//
// func findStartPos(grid [][]byte) (Pos, bool) {
// 	for r, row := range grid {
// 		for c := range row {
// 			if row[c] == '^' {
// 				return Pos{
// 					r, c,
// 				}, true
// 			}
// 		}
// 	}
//
// 	return Pos{}, false
// }
//
// func traverse(grid [][]byte) {
// 	DIFFS := []Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
// 	pos, ok := findStartPos(grid)
// 	if !ok {
// 		return
// 	}
//
// 	d := 0
// 	for pos.r >= 0 && pos.r < len(grid) && pos.c >= 0 && pos.c < len(grid[0]) {
// 		grid[pos.r][pos.c] = 'X'
// 		diff := DIFFS[d]
//
// 		nextPos := Pos{pos.r + diff.r, pos.c + diff.c}
// 		if nextPos.r >= 0 && nextPos.r < len(grid) && nextPos.c >= 0 && nextPos.c < len(grid[0]) {
// 			if grid[nextPos.r][nextPos.c] == '#' {
// 				d = (d + 1) % 4
// 				continue
// 			}
// 		}
//
// 		pos = nextPos
// 	}
// }
