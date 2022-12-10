package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("cmd/day_10/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var (
		x          = 1
		addx       = 0
		cycleCount = 0
		pSum       = 0
		queue      = make([]int, 2)
	)

	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, " ")
		if len(split) == 2 {
			addx, _ = strconv.Atoi(split[1])
			queue = append(queue, addx)
			for i := 0; i < 2; i++ {
				pSum += getPower(x, cycleCount)
				queue, x = cycle(queue, x)
				cycleCount++
			}
		} else {
			pSum += getPower(x, cycleCount)
			queue, x = cycle(queue, x)
			cycleCount++
		}
		queue = append(queue, 0)
	}
	fmt.Println(pSum)
}

func getPower(x, count int) int {
	if count%40 == 20 {
		fmt.Println("powerful signal at x:", x, "count:", count)
		return x * count
	}
	return 0
}

func cycle(signals []int, x int) ([]int, int) {
	x += signals[0]
	signals = signals[1:]
	return signals, x
}
