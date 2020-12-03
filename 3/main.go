package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("./3/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	treesProduct := traverseSlope(lines, 1, 1) *
		traverseSlope(lines, 3, 1) *
		traverseSlope(lines, 5, 1) *
		traverseSlope(lines, 7, 1) *
		traverseSlope(lines, 1, 2)

	log.Printf("Product of trees: %d", treesProduct)
}

func traverseSlope(lines []string, right int, down int) int {

	var column int
	var trees int

	for i := 0; i < len(lines); i += down {
		column = column % len(lines[i])
		if lines[i][column] == '#' {
			trees++
		}
		column += right
	}

	log.Printf("Found %d trees", trees)
	return trees
}
