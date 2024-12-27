package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse() []string {
	// f, _ := os.Open("example.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	board := make([]string, 0, 100)
	for scanner.Scan() {
		t := scanner.Text()
		board = append(board, t)
	}

	return board
}

type Pos struct {
	r, c int
}

var DIFFS = []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func getPath(board []string) []Pos {
	path := make([]Pos, 0, 1000)

	start := Pos{}
	end := Pos{}
	for r, row := range board {
		for c, b := range row {
			if b == 'S' {
				start = Pos{r, c}
			} else if b == 'E' {
				end = Pos{r, c}
			}
		}
	}

	path = append(path, start)

	for {
		pos := path[len(path)-1]

		if pos == end {
			break
		}

		for _, diff := range DIFFS {
			nextPos := Pos{pos.r + diff.r, pos.c + diff.c}
			if nextPos.r < 0 || nextPos.c < 0 || nextPos.r >= len(board) || nextPos.c >= len(board[0]) {
				continue
			}
			if board[nextPos.r][nextPos.c] == '#' {
				continue
			}
			if len(path) > 1 && path[len(path)-2] == nextPos {
				continue
			}
			path = append(path, nextPos)
			break
		}
	}

	return path
}

const (
	AT_LEAST_SAVE  = 100
	CHEAT_DURATION = 20
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	board := parse()
	for _, row := range board {
		fmt.Println(row)
	}

	path := getPath(board)
	// fmt.Println(path)

	ans := 0

	fmt.Println("Length of path is ", len(path))
	for i, pos := range path {
		for offset := AT_LEAST_SAVE + 1; offset+i < len(path); offset++ {
			nextPos := path[offset+i]
			dist := abs(nextPos.r-pos.r) + abs(nextPos.c-pos.c)
			// dist must be small enough to cheat and the time saved = offset - dist must be greater than the threshold
			if dist <= CHEAT_DURATION && offset-dist >= AT_LEAST_SAVE {
				// fmt.Println(nextPos, pos)
				ans++
			}
		}
	}
	fmt.Println(ans)
}
