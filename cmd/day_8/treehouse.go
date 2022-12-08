package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("cmd/day_8/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var grove [][]int

	for sc.Scan() {
		line := sc.Text()
		trees := mapTrees(line)
		grove = append(grove, trees)
	}
	visible := countVisible(grove)
	fmt.Println(visible)
}

func countVisible(g [][]int) int {
	row := len(g)
	col := len(g[0])
	c := row*col - (row-2)*(col-2)

	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if visible(g, i, j) {
				c++
			}
		}
	}
	return c
}

func visible(g [][]int, i, j int) bool {
	row := len(g)
	col := len(g[0])
	var t, b, l, r bool
	for top := 0; top < i; top++ {
		if g[top][j] >= g[i][j] {
			t = true
		}
	}
	for btm := i + 1; btm < row; btm++ {
		if g[btm][j] >= g[i][j] {
			b = true
		}
	}
	for left := 0; left < j; left++ {
		if g[i][left] >= g[i][j] {
			l = true
		}
	}
	for right := j + 1; right < col; right++ {
		if g[i][right] >= g[i][j] {
			r = true
		}
	}
	return !(t && b && l && r)
}

func mapTrees(l string) []int {
	result := make([]int, len(l))
	for i, r := range l {
		result[i] = int(r - '0')
	}
	return result
}
