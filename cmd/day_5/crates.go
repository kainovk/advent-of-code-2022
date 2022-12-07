package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("cmd/day_5/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var stacks [9][]rune

	sc.Scan()
	for ; sc.Text() != ""; sc.Scan() {
		for i, r := range sc.Text() {
			if r >= 'A' && r <= 'Z' {
				stacks[i/4] = append(stacks[i/4], r)
			}
		}
	}
	for _, stack := range stacks {
		stack = invert(stack)
	}

	for sc.Scan() {
		var count, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &count, &from, &to)

		stacks[from-1], stacks[to-1] = move(stacks[from-1], stacks[to-1], count)
	}
	result := ""
	for i := 0; i < len(stacks); i++ {
		l := len(stacks[i])
		result += string(stacks[i][l-1:])
	}
	fmt.Println(result)
}

func invert(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func move(from, to []rune, times int) ([]rune, []rune) {
	fl := len(from)
	crates := from[fl-times:]
	from = from[:fl-times]
	for i := 0; i < times; i++ {
		to = append(to, crates[i])
	}
	return from, to
}
