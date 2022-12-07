package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	name    string
	parent  *directory
	subdirs []*directory
	fsize   int
	dsize   int
}

func main() {
	file, _ := os.Open("cmd/day_7/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	root := directory{name: "/", parent: nil}
	var curDir = &root

	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")

		// Build the device directory structure.
		if line[0] == "$" && line[1] == "cd" {
			if line[2] == ".." {
				curDir = curDir.parent
			} else if line[2] != "/" {
				index := findDirectoryIndex(curDir.subdirs, line[2])
				curDir = curDir.subdirs[index]
			}
		} else if line[0] == "dir" {
			curDir.subdirs = append(curDir.subdirs,
				&directory{name: line[1], parent: curDir})
		} else {
			fsize, _ := strconv.Atoi(line[0])
			curDir.fsize += fsize
		}
	}
	setDirectoriesSize(&root)

	fmt.Println("Total size: ", root.fsize+root.dsize)
	result1 := partOne(root)
	fmt.Println(result1)

	const space = 70000000
	const need = 30000000
	taken := root.fsize + root.dsize
	size := getDirectoryToDeleteSize(taken, space-need, root)
	fmt.Println(size)
}

func getDirectoryToDeleteSize(space, bound int, root directory) int {
	var candidates []int
	search(space, bound, root, &candidates)
	sort.Ints(candidates)
	return candidates[0]
}

func search(space, bound int, root directory, c *[]int) {
	size := root.fsize + root.dsize
	if space-size <= bound {
		*c = append(*c, size)
	}
	for _, dir := range root.subdirs {
		search(space, bound, *dir, c)
	}
}

func partOne(root directory) int {
	result := 0
	if root.fsize+root.dsize <= 100000 {
		result += root.fsize + root.dsize
	}
	for _, dir := range root.subdirs {
		result += partOne(*dir)
	}
	return result
}

func setDirectoriesSize(root *directory) {
	iter := &root
	for _, dir := range (*iter).subdirs {
		(*iter).dsize += setDirectorySize(dir)
	}
}

func setDirectorySize(dir *directory) int {
	iter := &dir
	for _, subDir := range (*iter).subdirs {
		dir.dsize += setDirectorySize(subDir)
	}
	return dir.fsize + dir.dsize
}

func findDirectoryIndex(subdirs []*directory, dName string) int {
	for i, dir := range subdirs {
		if dir.name == dName {
			return i
		}
	}
	return -1
}
