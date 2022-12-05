package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("cmd/day_2/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	offset := 23
	total := 0
	for sc.Scan() {
		line := []byte(sc.Text())
		opp := int(line[0])
		me := int(line[2])
		diff := me - (opp + offset)

		points := 1
		if diff == 0 {
			points += 3
		} else if diff == 1 || diff == -2 {
			points += 6
		}

		if me == 'Y' {
			points += 1
		} else if me == 'Z' {
			points += 2
		}
		total += points
	}
	fmt.Println(total)
}
