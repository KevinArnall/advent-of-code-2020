package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	f, err := os.Open("./5/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	max := 0
	seatIds := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		row := binarySearch(line[:7], 128)
		col := binarySearch(line[7:], 8)
		seatID := row*8 + col

		seatIds = append(seatIds, seatID)

		if seatID > max {
			max = seatID
		}
	}

	fmt.Printf("Highest seat ID: %d\n", max)

	sort.Ints(seatIds)

	for i, id := range seatIds {
		if i == 0 {
			continue
		}

		if seatIds[i-1] != id-1 {
			fmt.Printf("Your seat id is: %d\n", id-1)
			break
		}
	}
}

func binarySearch(str string, rows int) int {

	min := 0
	max := rows
	location := rows / 2

	for _, letter := range str {
		if letter == 'F' || letter == 'L' {
			max = location
			location = (max + min) / 2
		} else {
			min = location
			location = (max + min) / 2
		}
	}
	return location
}
