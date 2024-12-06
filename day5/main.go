package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRules() map[int]map[int]struct{} {
	f, _ := os.Open("input_rules.txt")
	// f, _ := os.Open("example_rules.txt")
	scanner := bufio.NewScanner(f)

	before := make(map[int]map[int]struct{})
	for scanner.Scan() {
		text := scanner.Text()
		n1, _ := strconv.Atoi(text[:2])
		n2, _ := strconv.Atoi(text[3:])
		// fmt.Println(text)
		// fmt.Println(n1, n2)
		if before[n2] == nil {
			before[n2] = make(map[int]struct{})
		}
		before[n2][n1] = struct{}{}
	}
	return before
}

func parse() [][]int {
	f, _ := os.Open("input.txt")
	// f, _ := os.Open("example.txt")
	scanner := bufio.NewScanner(f)
	nums := make([][]int, 0, 100)
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.FieldsFunc(text, func(r rune) bool {
			return r == ','
		})
		row := make([]int, 0, 100)
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			row = append(row, num)
		}
		nums = append(nums, row)
	}
	return nums
}

// part 2
func fixOrder(row []int, before map[int]map[int]struct{}) []int {
	graph := make(map[int][]int, len(row))
	degree := make(map[int]int, len(row))
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			if _, ok := before[row[i]][row[j]]; ok {
				graph[row[j]] = append(graph[row[j]], row[i])
				degree[row[i]]++
			}
			if _, ok := before[row[j]][row[i]]; ok {
				graph[row[i]] = append(graph[row[i]], row[j])
				degree[row[j]]++
			}
		}
	}

	order := make([]int, 0, len(row))
	queue := make([]int, 0, len(row))
	for _, num := range row {
		if degree[num] == 0 {
			queue = append(queue, num)
		}
	}

	for len(queue) > 0 {
		num := queue[0]
		queue = queue[1:]
		order = append(order, num)

		for _, next := range graph[num] {
			degree[next]--
			if degree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return order
}

func main() {
	before := parseRules()
	nums := parse()
	// fmt.Println(before)
	// fmt.Println(nums)

	ans := 0
	for _, row := range nums {
		correct := true
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if _, ok := before[row[i]][row[j]]; ok {
					correct = false
				}
			}
		}
		if !correct {
			newRow := fixOrder(row, before)
			ans += newRow[len(newRow)/2]
		}
	}

	fmt.Println(ans)
}

// part 1
// func main() {
// 	before := parseRules()
// 	nums := parse()
// 	// fmt.Println(before)
// 	// fmt.Println(nums)
//
// 	ans := 0
// 	for _, row := range nums {
// 		correct := true
// 		for i := 0; i < len(row); i++ {
// 			for j := i + 1; j < len(row); j++ {
// 				if _, ok := before[row[i]][row[j]]; ok {
// 					correct = false
// 				}
// 			}
// 		}
// 		if correct {
// 			ans += row[len(row)/2]
// 		}
// 	}
//
// 	fmt.Println(ans)
// }
