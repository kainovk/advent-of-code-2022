package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("cmd/day_6/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanRunes)

	const seq = 14
	var queue []rune
	for i := 0; i < seq-1; i++ {
		sc.Scan()
		char := rune(sc.Text()[0])
		queue = append(queue, char)
	}

	result := seq - 1
	for sc.Scan() {
		result++
		char := rune(sc.Text()[0])
		queue = append(queue, char)
		if len(queue) == seq+1 {
			queue = queue[1:]
		}
		if len(toSet(queue)) == seq {
			break
		}
	}
	fmt.Println(result)
}

func toSet(q []rune) map[rune]bool {
	set := make(map[rune]bool)
	for _, r := range q {
		set[r] = true
	}
	return set
}
