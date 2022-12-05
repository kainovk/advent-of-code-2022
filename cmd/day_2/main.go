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

	total := 0
	for sc.Scan() {
		line := []byte(sc.Text())
		opp := int(line[0]+1) % 3
		me := int(line[2]+2) % 3

		total += calculatePoints(opp, me)
	}
	fmt.Println(total)
}

func calculatePoints(opp, me int) int {
	result := 0
	if me == 0 {
		result += (opp + 2) % 3
	} else if me == 1 {
		result += 3
		result += opp
	} else {
		result += 6
		result += (opp + 1) % 3
	}
	return result + 1
}
