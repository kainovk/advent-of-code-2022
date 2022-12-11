package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     *[]int
	operation rune
	operand   int
	divTest   int
	throw     [2]int
}

func (m monkey) play(item int) int {
	operand := m.operand
	if operand == -1 {
		operand = item
	}
	if m.operation == '+' {
		item += operand
	} else {
		item *= operand
	}
	return item / 3
}

func main() {
	file, _ := os.Open("cmd/day_11/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	monkeys := make([]monkey, 0)
	const rounds = 20

	for sc.Scan() {
		sc.Scan()
		line := strings.TrimSpace(sc.Text())
		split := strings.Split(line, " ")
		items := mapItems(split[2:])

		sc.Scan()
		line = strings.TrimSpace(sc.Text())
		split = strings.Split(line, " ")
		operation := rune(split[4][0])
		operand := mapOperand(split[5])

		sc.Scan()
		line = strings.TrimSpace(sc.Text())
		split = strings.Split(line, " ")
		divTest, _ := strconv.Atoi(split[3])

		throw := [2]int{0, 0}
		sc.Scan()
		line = strings.TrimSpace(sc.Text())
		split = strings.Split(line, " ")
		throw[0], _ = strconv.Atoi(split[5])
		sc.Scan()
		line = strings.TrimSpace(sc.Text())
		split = strings.Split(line, " ")
		throw[1], _ = strconv.Atoi(split[5])
		sc.Scan()

		monkeys = append(monkeys,
			monkey{
				&items,
				operation,
				operand,
				divTest,
				throw,
			},
		)
	}

	throwsCount := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for i, monkey := range monkeys {
			for _, item := range *monkey.items {
				item = monkey.play(item)
				targetMonkey := 0
				if item%monkey.divTest == 0 {
					targetMonkey = monkey.throw[0]
				} else {
					targetMonkey = monkey.throw[1]
				}
				*monkeys[targetMonkey].items = append(*monkeys[targetMonkey].items, item)
			}
			throwsCount[i] += len(*monkey.items)
			*monkey.items = make([]int, 0)
		}
	}

	sort.Ints(throwsCount)
	fmt.Println("Monkey business:",
		throwsCount[len(throwsCount)-1]*throwsCount[len(throwsCount)-2])
}

func mapOperand(o string) int {
	if o == "old" {
		return -1
	}
	op, _ := strconv.Atoi(o)
	return op
}

func mapItems(line []string) []int {
	l := len(line)
	items := make([]int, l)
	for i := 0; i < l-1; i++ {
		item := line[i][:len(line[i])-1]
		nItem, _ := strconv.Atoi(item)
		items[i] = nItem
	}
	item := line[l-1]
	nItem, _ := strconv.Atoi(item)
	items[l-1] = nItem
	return items
}
