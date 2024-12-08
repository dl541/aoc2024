package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

type Pos struct {
	r, c int
}

func main() {
	// f, _ := os.Open("example.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	board := make([]string, 0, 100)
	for scanner.Scan() {
		text := scanner.Text()
		board = append(board, text)
	}

	positions := make(map[byte][]Pos, 100)

	for r, row := range board {
		for c := range row {
			if row[c] != '.' {
				positions[row[c]] = append(positions[row[c]], Pos{r, c})
			}
		}
	}

	antinodes := getAntinodes(positions, board)
	keys := make([]Pos, 0, len(antinodes))
	for k := range antinodes {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(k1, k2 Pos) int {
		if k1.r == k2.r {
			return cmp.Compare(k1.c, k2.c)
		}
		return cmp.Compare(k1.r, k2.r)
	})

	fmt.Println(keys)
	fmt.Println(len(antinodes))
}

// part 2
func getAntinodes(positions map[byte][]Pos, board []string) map[Pos]struct{} {
	antinodes := make(map[Pos]struct{}, 100)
	for _, row := range positions {
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				p1 := row[i]
				p2 := row[j]
				antinodes[p1] = struct{}{}
				antinodes[p2] = struct{}{}
				diff := Pos{p1.r - p2.r, p1.c - p2.c}
				// p2 - diff
				n := 1
				for {
					p3 := Pos{p2.r - n*diff.r, p2.c - n*diff.c}
					if p3.r >= 0 && p3.c >= 0 && p3.r < len(board) && p3.c < len(board[0]) {
						antinodes[p3] = struct{}{}
					} else {
						break
					}
					n++
				}
				n = 1
				// p1 + diff
				for {
					p3 := Pos{p1.r + n*diff.r, p1.c + n*diff.c}
					if p3.r >= 0 && p3.c >= 0 && p3.r < len(board) && p3.c < len(board[0]) {
						antinodes[p3] = struct{}{}
					} else {
						break
					}
					n++
				}

			}
		}
	}
	return antinodes
}

// part 1
// func getAntinodes(positions map[byte][]Pos, board []string) map[Pos]struct{} {
// 	antinodes := make(map[Pos]struct{}, 100)
// 	for _, row := range positions {
// 		for i := 0; i < len(row); i++ {
// 			for j := i + 1; j < len(row); j++ {
// 				p1 := row[i]
// 				p2 := row[j]
// 				diff := Pos{p1.r - p2.r, p1.c - p2.c}
// 				p3 := Pos{p1.r + diff.r, p1.c + diff.c}
// 				p4 := Pos{p2.r - diff.r, p2.c - diff.c}
//
// 				if p3.r >= 0 && p3.c >= 0 && p3.r < len(board) && p3.c < len(board[0]) {
// 					antinodes[p3] = struct{}{}
// 				}
// 				if p4.r >= 0 && p4.c >= 0 && p4.r < len(board) && p4.c < len(board[0]) {
// 					antinodes[p4] = struct{}{}
// 				}
// 			}
// 		}
// 	}
// 	return antinodes
// }
