package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("cmd/day_1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var max int
	var currentCalories int
	for scanner.Scan() {
		line := scanner.Text()
		calories, err := strconv.Atoi(line)
		currentCalories += calories

		if err != nil || err == io.EOF {
			if max < currentCalories {
				max = currentCalories
			}
			currentCalories = 0
		}
	}
	fmt.Println(max)
}
