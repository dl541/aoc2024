package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var file string

func init() {
	flag.StringVar(&file, "file", "example.txt", "File from which to read the input data")
}

func main() {
	flag.Parse()
	file, _ := os.Open(file)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		nums := make([]int, 0, len(line))
		for _, token := range line {
			digit, _ := strconv.Atoi(token)
			nums = append(nums, digit)
		}
		// if isSafe(nums) {
		// 	count++
		// }
		if isSafeBruteForce(nums) {
			count++
		}
	}

	fmt.Println(count)
}

// part 2
func isSafe(nums []int) bool {
	diffs := make([]int, 0, len(nums))

	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}

	return isSafeDiffsSkippable(diffs, 1) || isSafeDiffsSkippable(diffs, -1)
}

func isSafeDiffsSkippable(diffs []int, sign int) bool {
	for i, d := range diffs {
		if d*sign <= 0 || abs(d) > 3 || abs(d) < 1 {
			if i > 0 {
				oldDiff := diffs[i]
				newDiff := diffs[i] + diffs[i-1]
				diffs[i] = newDiff
				if isSafeDiffs(diffs[i:], sign) {
					return true
				}
				diffs[i] = oldDiff
			}
			if i < len(diffs)-1 {
				oldDiff := diffs[i+1]
				newDiff := diffs[i] + diffs[i+1]
				diffs[i+1] = newDiff
				if isSafeDiffs(diffs[i+1:], sign) {
					return true
				}
				diffs[i+1] = oldDiff
			}
			return false
		}
	}
	return true
}

func isSafeDiffs(diffs []int, sign int) bool {
	for _, d := range diffs {
		if d*sign <= 0 || abs(d) > 3 || abs(d) < 1 {
			return false
		}
	}
	return true
}

func isSafeBruteForce(nums []int) bool {
	diffs := make([]int, 0, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}

	if isSafeDiffs(diffs, 1) || isSafeDiffs(diffs, -1) {
		return true
	}
	for i := 0; i < len(diffs)-1; i++ {
		newDiff := diffs[i] + diffs[i+1]
		concat := slices.Concat(diffs[:i+1], diffs[i+2:])
		concat[i] = newDiff
		if isSafeDiffs(concat, 1) || isSafeDiffs(concat, -1) {
			return true
		}
	}
	return false
}

// part 1
// func isSafe(nums []int) bool {
// 	for i := 0; i < len(nums)-1; i++ {
// 		if dist := abs(nums[i] - nums[i+1]); dist > 3 || dist < 1 {
// 			return false
// 		}
// 		if (nums[i]-nums[i+1])*(nums[0]-nums[1]) <= 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
