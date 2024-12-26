package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Pair struct {
	index, size int
}

func main() {
	f, _ := os.Open("input.txt")
	// f, _ := os.Open("example.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	nums := make([]int, 0, 10000)
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		nums = append(nums, num)
	}

	expanded := expand(nums)
	// fmt.Println(expanded)
	compacted := compact(expanded)
	// fmt.Println(compacted)
	// // fmt.Println(expanded)
	// fmt.Println(getCheckSum(expanded))
	fmt.Println(getCheckSum(compacted))
}

func getCheckSum(compacted []Pair) int {
	ans := 0
	count := 0
	for _, p := range compacted {
		for range p.size {
			if p.index != -1 {
				ans += p.index * count
			}
			count++
		}
	}
	return ans
}

func compact(pairs []Pair) []Pair {
	sizeToPairs := make(map[int][]Pair, 10)

	remainingPairs := make(map[Pair]struct{}, len(pairs))
	for i := len(pairs) - 1; i >= 0; i-- {
		p := pairs[i]
		if p.index != -1 {
			sizeToPairs[p.size] = append(sizeToPairs[p.size], p)
			remainingPairs[p] = struct{}{}
		}
	}

	res := make([]Pair, 0, len(pairs))
	for _, pair := range pairs {
		// fmt.Println(sizeToPairs, res)
		if pair.index != -1 {
			if _, ok := remainingPairs[pair]; !ok {
				res = append(res, Pair{-1, pair.size})
				continue
			}
			res = append(res, pair)
			delete(remainingPairs, pair)
		} else {
			for pair.size > 0 {
				chosen := Pair{math.MinInt, -1}
				for size, queue := range sizeToPairs {
					if size <= pair.size && len(queue) != 0 && queue[0].index > chosen.index {
						chosen = queue[0]
					}
				}

				if chosen.size != -1 {
					res = append(res, chosen)
					delete(remainingPairs, chosen)
					for len(sizeToPairs[chosen.size]) > 0 {
						if _, ok := remainingPairs[sizeToPairs[chosen.size][0]]; !ok {
							sizeToPairs[chosen.size] = sizeToPairs[chosen.size][1:]
						} else {
							break
						}
					}
					pair.size -= chosen.size
				} else {
					res = append(res, pair)
					break
				}
			}
		}
	}
	return res
}

func expand(nums []int) []Pair {
	expanded := make([]Pair, 0, 100000)
	for i, num := range nums {
		token := -1
		if i%2 == 0 {
			token = i / 2
		}
		expanded = append(expanded, Pair{token, num})
	}
	return expanded
}

// func getCheckSum(nums []int) int {
// 	total := 0
// 	for i, num := range nums {
// 		if num == -1 {
// 			break
// 		}
// 		total += i * num
// 	}
// 	return total
// }
//
// func compact(nums []int) {
// 	left := 0
// 	right := len(nums) - 1
// 	for left < right {
// 		// fmt.Println(left, right)
// 		if nums[left] != -1 {
// 			left++
// 			continue
// 		}
// 		if nums[right] == -1 {
// 			right--
// 			continue
// 		}
// 		nums[left], nums[right] = nums[right], nums[left]
// 		left++
// 		right--
// 	}
// }
//
// func expand(nums []int) []int {
// 	expanded := make([]int, 0, 100000)
// 	for i, num := range nums {
// 		token := -1
// 		if i%2 == 0 {
// 			token = i / 2
// 		}
// 		for range num {
// 			expanded = append(expanded, token)
// 		}
// 	}
//
// 	return expanded
// }
