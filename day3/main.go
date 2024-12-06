package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var file string

func init() {
	flag.StringVar(&file, "file", "example.txt", "File from which to read the input data")
}

func parseByRegex(bytes []byte, enabled bool) (int, bool) {
	// regex := regexp.MustCompile(`mul\([\d]+,[\d]+\)`)
	numRegex := regexp.MustCompile(`[\d]+`)
	// dontRegex := regexp.MustCompile(`don\'t\(\)`)
	// doRegex := regexp.MustCompile(`do\(\)`)
	overallRegex := regexp.MustCompile(`(mul\([\d]+,[\d]+\))|(don\'t\(\))|(do\(\))`)
	s := string(bytes)
	sum := 0
	for _, sub := range overallRegex.FindAllString(s, -1) {
		fmt.Println(sub)
		if sub[0] == 'm' {
			if !enabled {
				continue
			}
			nums := numRegex.FindAllString(sub, -1)
			fmt.Println(nums)
			n2, _ := strconv.Atoi(nums[1])
			n1, _ := strconv.Atoi(nums[0])
			sum += n1 * n2
		} else if sub[:3] == "don" {
			enabled = false
		} else if sub[:2] == "do" {
			enabled = true
		}
	}
	fmt.Println(sum)
	return sum, enabled
}

func main() {
	flag.Parse()
	file, _ := os.Open(file)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()

		num, res := parseByRegex([]byte(line), enabled)
		enabled = res
		sum += num
		// sum += parse([]byte(line))
	}
	fmt.Println(sum)
}

// func parse(bytes []byte) int {
// 	sum := 0
// 	for probe := 0; probe < len(bytes); probe++ {
// 		pt := probe
// 		if pt+4 <= len(bytes) && string(bytes[pt:pt+4]) == "mul(" {
// 			fmt.Println("pt now", string(bytes[pt:]))
// 			pt += 4
// 			left := pt
// 			for ; pt < len(bytes) && bytes[pt] != ','; pt++ {
// 			}
// 			if pt >= len(bytes) || bytes[pt] != ',' {
// 				continue
// 			}
// 			n1, err := strconv.Atoi(string(bytes[left:pt]))
// 			if err != nil {
// 				continue
// 			}
// 			pt++
//
// 			right := pt
// 			for ; pt < len(bytes) && bytes[pt] != ')'; pt++ {
// 			}
// 			if pt >= len(bytes) || bytes[pt] != ')' {
// 				continue
// 			}
// 			n2, err := strconv.Atoi(string(bytes[right:pt]))
// 			if err != nil {
// 				continue
// 			}
// 			fmt.Println(n1, n2)
// 			sum += n1 * n2
// 		}
// 	}
//
// 	fmt.Println(sum)
// 	return sum
// }
//
// func parseByRegex(bytes []byte) int {
// 	regex := regexp.MustCompile(`mul\([\d]+,[\d]+\)`)
// 	numRegex := regexp.MustCompile(`[\d]+`)
// 	s := string(bytes)
// 	sum := 0
// 	for _, sub := range regex.FindAllString(s, -1) {
// 		fmt.Println(sub)
// 		nums := numRegex.FindAllString(sub, -1)
// 		fmt.Println(nums)
// 		n1, _ := strconv.Atoi(nums[0])
// 		n2, _ := strconv.Atoi(nums[1])
// 		sum += n1 * n2
// 	}
// 	fmt.Println(sum)
// 	return sum
// }

// func main() {
// 	flag.Parse()
// 	file, _ := os.Open(file)
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
// 	sum := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		sum += parseByRegex([]byte(line))
// 		// sum += parse([]byte(line))
// 	}
// 	fmt.Println(sum)
// }
