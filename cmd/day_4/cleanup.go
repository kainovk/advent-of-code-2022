package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("cmd/day_4/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	cFullyOverlaps := 0
	cOverlaps := 0
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, ",")
		if fullyContains(parts[0], parts[1]) {
			cFullyOverlaps++
		}
		if overlaps(parts[0], parts[1]) {
			cOverlaps++
		}
	}
	fmt.Println(cFullyOverlaps)
	fmt.Println(cOverlaps)
}

func overlaps(s1, s2 string) bool {
	elf1 := strings.Split(s1, "-")
	elf2 := strings.Split(s2, "-")
	start1, _ := strconv.ParseInt(elf1[0], 10, 64)
	start2, _ := strconv.ParseInt(elf2[0], 10, 64)
	end1, _ := strconv.ParseInt(elf1[1], 10, 64)
	end2, _ := strconv.ParseInt(elf2[1], 10, 64)
	if start1 <= end2 && start2 <= end1 ||
		start1 <= start2 && end2 <= end1 {
		return true
	}
	return false
}

func fullyContains(s1, s2 string) bool {
	elf1 := strings.Split(s1, "-")
	elf2 := strings.Split(s2, "-")
	start1, _ := strconv.ParseInt(elf1[0], 10, 64)
	start2, _ := strconv.ParseInt(elf2[0], 10, 64)
	end1, _ := strconv.ParseInt(elf1[1], 10, 64)
	end2, _ := strconv.ParseInt(elf2[1], 10, 64)
	if start1 <= start2 && end2 <= end1 ||
		start2 <= start1 && end1 <= end2 {
		return true
	}
	return false
}
