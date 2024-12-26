package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func parse() ([]string, []string) {
	fPatterns, _ := os.Open("input_patterns.txt")
	scannerPattern := bufio.NewScanner(fPatterns)
	scannerPattern.Scan()
	patterns := strings.FieldsFunc(scannerPattern.Text(), func(r rune) bool {
		return r == ','
	})
	for i, p := range patterns {
		patterns[i] = strings.TrimSpace(p)
	}
	fTowels, _ := os.Open("input_towels.txt")
	towels := make([]string, 0, 1000)
	scannerTowel := bufio.NewScanner(fTowels)
	for scannerTowel.Scan() {
		towels = append(towels, scannerTowel.Text())
	}
	return patterns, towels
}

func canDo(patterns map[string]bool, maxPattern int, towel string) int {
	fmt.Println(towel)
	var dp func(i int) int
	cache := make([]int, len(towel))
	for i := range cache {
		cache[i] = -1
	}
	dp = func(i int) int {
		if i == len(towel) {
			return 1
		}
		if cache[i] != -1 {
			return cache[i]
		}

		ans := 0
		for length := 1; length <= maxPattern && i+length <= len(towel); length++ {
			substr := towel[i : i+length]
			if patterns[substr] {
				ans += dp(i + length)
			}
		}

		cache[i] = ans
		return ans
	}
	ans := dp(0)
	// fmt.Println(cache)
	return ans
}

func main() {
	patterns, towels := parse()
	fmt.Println(len(patterns), len(towels))

	maxPattern := len(slices.MaxFunc(patterns, func(s1, s2 string) int {
		return cmp.Compare(len(s1), len(s2))
	}))

	patternsSet := make(map[string]bool, len(patterns))
	for _, p := range patterns {
		patternsSet[p] = true
	}
	ans := 0
	fmt.Println(patternsSet)
	for _, t := range towels {
		ans += canDo(patternsSet, maxPattern, t)
	}
	fmt.Println(ans)
}
