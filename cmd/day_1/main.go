package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("cmd/day_1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		max             int
		currentCalories int
		snacks          []int
	)
	for scanner.Scan() {
		line := scanner.Text()
		calories, err := strconv.Atoi(line)
		currentCalories += calories

		if err != nil || err == io.EOF {
			if max < currentCalories {
				max = currentCalories
			}
			snacks = append(snacks, currentCalories)
			currentCalories = 0
		}
	}
	sort.Ints(snacks)
	topThree := snacks[len(snacks)-3:]
	fmt.Println(sum(topThree))
}

func sum(ints []int) int {
	result := 0
	for _, v := range ints {
		result += v
	}
	return result
}
