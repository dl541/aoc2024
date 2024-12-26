package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARD_SIZE = 71

// const BOARD_SIZE = 7

func parse() [][]int {
	// f, _ := os.Open("example.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	blocks := make([][]int, 0, 1000)
	for scanner.Scan() {
		txt := scanner.Text()
		tokens := strings.FieldsFunc(txt, func(t rune) bool {
			return t == ','
		})
		y, _ := strconv.Atoi(tokens[0])
		x, _ := strconv.Atoi(tokens[1])
		blocks = append(blocks, []int{x, y})
	}
	return blocks
}

func getBoard(blocks [][]int) [][]int {
	board := make([][]int, BOARD_SIZE)
	for i := range BOARD_SIZE {
		board[i] = make([]int, BOARD_SIZE)
	}

	for _, b := range blocks {
		board[b[0]][b[1]] = 1
	}
	return board
}

type Pos struct {
	r, c int
}

var DIFFS = []Pos{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func bfs(board [][]int) int {
	q := []Pos{{0, 0}}
	ans := 0
	for len(q) > 0 {
		qLen := len(q)
		for _, pos := range q {
			if pos.r == BOARD_SIZE-1 && pos.c == BOARD_SIZE-1 {
				return ans
			}
			for _, diff := range DIFFS {
				nextPos := Pos{diff.r + pos.r, diff.c + pos.c}
				if nextPos.r < 0 || nextPos.c < 0 || nextPos.r >= BOARD_SIZE || nextPos.c >= BOARD_SIZE {
					continue
				}
				if board[nextPos.r][nextPos.c] == 1 {
					continue
				}
				board[nextPos.r][nextPos.c] = 1
				q = append(q, Pos{nextPos.r, nextPos.c})
			}
		}
		q = q[qLen:]
		ans++
	}
	return -1
}

func main() {
	blocks := parse()
	left := 0
	right := len(blocks)
	for left < right {
		mid := (left + right) / 2
		board := getBoard(blocks[:mid+1])
		ans := bfs(board)
		if ans == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(blocks[left])
}
