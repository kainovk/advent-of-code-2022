package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func (p position) near(pp position) bool {
	return (p.x-pp.x)*(p.x-pp.x) <= 1 && (p.y-pp.y)*(p.y-pp.y) <= 1
}

var (
	visited = make(map[position]bool)
	hPos    = position{0, 0}
	tPos    = position{0, 0}
)

func main() {
	file, _ := os.Open("cmd/day_9/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	visited[hPos] = true
	for sc.Scan() {
		line := sc.Text()
		m := strings.Split(line, " ")
		vec := mapToPosition(m[0])
		steps, _ := strconv.Atoi(m[1])
		move := position{vec.x * steps, vec.y * steps}
		makeMove(move)
	}
	fmt.Println("Head position:", hPos)
	fmt.Println("Tail position:", tPos)
	fmt.Println("Visited:", len(visited))
}

func makeMove(move position) {
	hPos.x += move.x
	hPos.y += move.y
	if !hPos.near(tPos) {
		if math.Abs(float64(hPos.x-tPos.x)) <= 1 {
			tPos.x = hPos.x
		} else {
			tPos.y = hPos.y
		}
		if hPos.x == tPos.x {
			if hPos.y > tPos.y {
				for hPos.y-1 != tPos.y {
					tPos.y++
					visited[tPos] = true
				}
			} else {
				for hPos.y+1 != tPos.y {
					tPos.y--
					visited[tPos] = true
				}
			}
		} else {
			if hPos.x > tPos.x {
				for hPos.x-1 != tPos.x {
					tPos.x++
					visited[tPos] = true
				}
			} else {
				for hPos.x+1 != tPos.x {
					tPos.x--
					visited[tPos] = true
				}
			}
		}
	}
}

func mapToPosition(s string) position {
	switch s {
	case "U":
		return position{0, 1}
	case "D":
		return position{0, -1}
	case "L":
		return position{-1, 0}
	case "R":
		return position{1, 0}
	}
	return position{0, 0}
}
