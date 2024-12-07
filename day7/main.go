package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		fmt.Println("Time consumed", time.Since(t))
	}()

	f, _ := os.Open("input.txt")
	// f, _ := os.Open("example.txt")
	scanner := bufio.NewScanner(f)

	var ans atomic.Int64
	var wg sync.WaitGroup
	for scanner.Scan() {
		text := scanner.Text()
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			tokens := strings.Fields(text)
			tokens[0] = tokens[0][:len(tokens[0])-1]
			// fmt.Println(tokens)
			target, _ := strconv.Atoi(tokens[0])
			nums := make([]int, 0, len(tokens)-1)
			for i := 1; i < len(tokens); i++ {
				num, _ := strconv.Atoi(tokens[i])
				nums = append(nums, num)
			}

			if possible(nums, target) {
				// fmt.Println("Possible", nums, target)
				ans.Add(int64(target))
			}
		}(text)
	}

	wg.Wait()
	fmt.Println(ans.Load())
}

// part 2
func possible(nums []int, target int) bool {
	var backtrack func(i int, state int) bool

	backtrack = func(i int, state int) bool {
		if i == len(nums) {
			return target == state
		}
		if state > target {
			return false
		}
		possible := backtrack(i+1, state*nums[i]) || backtrack(i+1, state+nums[i])
		if possible {
			return true
		}
		base := 10
		for base <= nums[i] {
			base *= 10
		}
		return backtrack(i+1, base*state+nums[i])
	}
	return backtrack(1, nums[0])
}

// part 1
// func possible(nums []int, target int) bool {
// 	var backtrack func(i int, state int) bool
//
// 	backtrack = func(i int, state int) bool {
// 		if i == len(nums) {
// 			return target == state
// 		}
// 		if state > target {
// 			return false
// 		}
// 		return backtrack(i+1, state*nums[i]) || backtrack(i+1, state+nums[i])
// 	}
// 	return backtrack(1, nums[0])
// }
