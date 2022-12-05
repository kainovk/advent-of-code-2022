package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("cmd/day_3/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	totalSum := 0
	for sc.Scan() {
		line := sc.Text()
		half := len(line) / 2
		left := line[:half]
		right := line[half:]

		totalSum += sumDuplicatesPriorities(left, right)
	}
	fmt.Println(totalSum)
}

func sumDuplicatesPriorities(left, right string) int {
	seen := make(map[rune]bool)
	for _, r := range left {
		seen[r] = true
	}

	result := 0
	dupSet := make(map[rune]bool)
	for _, r := range right {
		if seen[r] == true {
			exist := dupSet[r]
			if !exist {
				result += priority(r)
				dupSet[r] = true
			}
		}
	}
	return result
}

func priority(r rune) int {
	p := 0
	if r <= 90 {
		p = int(r - rune(38))
	} else {
		p = int(r - rune(96))
	}
	return p
}
