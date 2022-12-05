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
	var group []string
	commonItemPrioritySum := 0
	for sc.Scan() {
		line := sc.Text()
		half := len(line) / 2
		left := line[:half]
		right := line[half:]

		totalSum += sumDuplicatesPriorities(left, right)

		group = append(group, line)
		if len(group) == 3 {
			commonItemPrioritySum += calculateCommonItemPriority(group)
			group = nil
		}
	}
	fmt.Println(totalSum)
	fmt.Println(commonItemPrioritySum)
}

func calculateCommonItemPriority(group []string) int {
	cMap := make(map[rune]byte)
	for _, s := range group {
		set := toRuneSet(s)
		for r := range set {
			cMap[r]++
		}
	}

	result := 0
	for r, c := range cMap {
		if c == 3 {
			result += priority(r)
		}
	}
	return result
}

func toRuneSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, r := range s {
		set[r] = true
	}
	return set
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
		p = int(r) - 38
	} else {
		p = int(r) - 96
	}
	return p
}
