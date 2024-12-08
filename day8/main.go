package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sync"
	"time"
)

type Pos struct {
	r, c int
}

func main() {
	t := time.Now()
	defer func() {
		fmt.Println("Time consumed", time.Since(t))
	}()
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

	// fmt.Println(keys)
	fmt.Println(len(antinodes))
}

type Antinodes struct {
	nodes map[Pos]struct{}
	mtx   sync.Mutex
}

func NewAntinodes() Antinodes {
	return Antinodes{
		nodes: make(map[Pos]struct{}, 100),
	}
}

func scanRow(row []Pos, antinodes *Antinodes, wg *sync.WaitGroup, board []string) {
	defer wg.Done()
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			p1 := row[i]
			p2 := row[j]

			func() {
				antinodes.mtx.Lock()
				defer antinodes.mtx.Unlock()
				antinodes.nodes[p1] = struct{}{}
				antinodes.nodes[p2] = struct{}{}
			}()

			diff := Pos{p1.r - p2.r, p1.c - p2.c}
			// p2 - diff
			n := 1
			for {
				p3 := Pos{p2.r - n*diff.r, p2.c - n*diff.c}
				if p3.r >= 0 && p3.c >= 0 && p3.r < len(board) && p3.c < len(board[0]) {
					func() {
						antinodes.mtx.Lock()
						defer antinodes.mtx.Unlock()
						antinodes.nodes[p3] = struct{}{}
					}()
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
					func() {
						antinodes.mtx.Lock()
						defer antinodes.mtx.Unlock()
						antinodes.nodes[p3] = struct{}{}
					}()
				} else {
					break
				}
				n++
			}
		}
	}
}

// part 2
func getAntinodes(positions map[byte][]Pos, board []string) map[Pos]struct{} {
	antinodes := NewAntinodes()

	var wg sync.WaitGroup
	for _, row := range positions {
		wg.Add(1)
		go scanRow(row, &antinodes, &wg, board)
	}
	wg.Wait()
	return antinodes.nodes
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
