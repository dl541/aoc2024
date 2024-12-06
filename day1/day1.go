package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

// Part 2
func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lst1 := make([]int, 0, 100)
	counter := make(map[int]int, 100)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		n1, _ := strconv.Atoi(text)
		lst1 = append(lst1, n1)
		scanner.Scan()
		text = scanner.Text()
		n2, _ := strconv.Atoi(text)
		counter[n2]++
	}

	slices.Sort(lst1)
	ans := 0
	for _, num := range lst1 {
		ans += num * counter[num]
	}
	fmt.Println(ans)
}

// Part 1
// func main() {
// 	file, _ := os.Open("input.txt")
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	lst1 := make([]int, 0, 100)
// 	lst2 := make([]int, 0, 100)
// 	scanner.Split(bufio.ScanWords)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		n1, _ := strconv.Atoi(text)
// 		lst1 = append(lst1, n1)
// 		scanner.Scan()
// 		text = scanner.Text()
// 		n2, _ := strconv.Atoi(text)
// 		lst2 = append(lst2, n2)
// 	}
//
// 	slices.Sort(lst1)
// 	slices.Sort(lst2)
// 	ans := 0
// 	for i := range lst1 {
// 		ans += abs(lst1[i] - lst2[i])
// 	}
// 	fmt.Println(ans)
// }

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
