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
	score := getBestScenicScore(grove)
	fmt.Println(visible)
	fmt.Println(score)
}

func getBestScenicScore(g [][]int) int {
	row := len(g)
	col := len(g[0])
	bestScore := 0
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			score := getScore(g, i, j)
			sScore := countScenicScore(score)
			if bestScore < sScore {
				bestScore = sScore
			}
		}
	}
	return bestScore
}

func getScore(g [][]int, i, j int) []int {
	score := make([]int, 4)
	row := len(g)
	col := len(g[0])
	var cur int
	for top := i - 1; top >= 0; top-- {
		score[0]++
		cur = g[top][j]
		if cur >= g[i][j] {
			break
		}
	}
	for btm := i + 1; btm < row; btm++ {
		score[1]++
		cur = g[btm][j]
		if cur >= g[i][j] {
			break
		}
	}
	for left := j - 1; left >= 0; left-- {
		score[2]++
		cur = g[i][left]
		if cur >= g[i][j] {
			break
		}
	}
	for right := j + 1; right < col; right++ {
		score[3]++
		cur = g[i][right]
		if cur >= g[i][j] {
			break
		}
	}
	return score
}

func countScenicScore(ss []int) int {
	result := 1
	for _, s := range ss {
		result *= s
	}
	return result
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
	for top := i - 1; top >= 0; top-- {
		if g[top][j] >= g[i][j] {
			t = true
		}
	}
	for btm := i + 1; btm < row; btm++ {
		if g[btm][j] >= g[i][j] {
			b = true
		}
	}
	for left := j - 1; left >= 0; left-- {
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
