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
	last *position
}

func (p position) near(pp position) bool {
	return (p.x-pp.x)*(p.x-pp.x) <= 1 && (p.y-pp.y)*(p.y-pp.y) <= 1
}

func (p position) diagonal(pp position) bool {
	return (p.x-pp.x)*(p.x-pp.x) == 4 && (p.y-pp.y)*(p.y-pp.y) == 4
}

var (
	visited = make(map[position]bool)
	hPos    = position{0, 0, nil}
)

const knots = 10 - 1

func main() {
	file, _ := os.Open("cmd/day_9/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	makeRope()

	visited[hPos] = true
	for sc.Scan() {
		line := sc.Text()
		m := strings.Split(line, " ")
		vec := mapToPosition(m[0])
		steps, _ := strconv.Atoi(m[1])
		move := position{vec.x * steps, vec.y * steps, nil}
		makeMove(move)
		printRope()
	}
	fmt.Println("Visited:", len(visited))
}

func printRope() {
	iter := &hPos
	for i := 0; i < knots; i++ {
		fmt.Println("i:", i, "x:", iter.x, "y:", iter.y)
		iter = iter.last
	}
	fmt.Println("i:", knots, "x:", iter.x, "y:", iter.y)
	fmt.Println("---")
}

func makeRope() {
	iter := &hPos
	for i := 0; i < knots; i++ {
		iter.last = &position{0, 0, nil}
		iter = iter.last
	}
}

func makeMove(move position) {
	if move.x == 0 {
		for move.y != 0 {
			makeMoveY(move.y)
			move.y += -1 * move.y / int(math.Abs(float64(move.y)))
			cling()
		}
	} else {
		for move.x != 0 {
			makeMoveX(move.x)
			move.x += -1 * move.x / int(math.Abs(float64(move.x)))
			cling()
		}
	}
}

func cling() {
	iter := &hPos
	for i := 0; i < knots; i++ {
		last := iter.last
		if !iter.near(*last) && !iter.diagonal(*last) {
			if math.Abs(float64(iter.x-last.x)) <= 1 {
				last.x = iter.x
			} else {
				last.y = iter.y
			}
			if iter.x == last.x {
				if iter.y > last.y {
					for iter.y-1 != last.y {
						last.y++
						if last.last == nil {
							visited[*last] = true
						}
					}
				} else {
					for iter.y+1 != last.y {
						last.y--
						if last.last == nil {
							visited[*last] = true
						}
					}
				}
			} else {
				if iter.x > last.x {
					for iter.x-1 != last.x {
						last.x++
						if last.last == nil {
							visited[*last] = true
						}
					}
				} else {
					for iter.x+1 != last.x {
						last.x--
						if last.last == nil {
							visited[*last] = true
						}
					}
				}
			}
		} else if iter.diagonal(*last) {
			if iter.x > last.x {
				last.x++
			} else {
				last.x--
			}
			if iter.y > last.y {
				last.y++
			} else {
				last.y--
			}
			if last.last == nil {
				visited[*last] = true
			}
		}
		iter = last
	}
}

func makeMoveX(x int) {
	if x > 0 {
		hPos.x++
	} else {
		hPos.x--
	}
}

func makeMoveY(y int) {
	if y > 0 {
		hPos.y++
	} else {
		hPos.y--
	}
}

func mapToPosition(s string) position {
	switch s {
	case "U":
		return position{0, 1, nil}
	case "D":
		return position{0, -1, nil}
	case "L":
		return position{-1, 0, nil}
	case "R":
		return position{1, 0, nil}
	}
	return position{0, 0, nil}
}
