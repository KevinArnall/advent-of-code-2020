package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./9/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		number, _ := strconv.Atoi(line)

		numbers = append(numbers, number)
	}

	ans := part1(numbers)
	fmt.Printf("Answer for number 1: %d", ans)
}

func part1(numbers []int) int {
	start := 0
	end := 25

	for _, currentNum := range numbers[end:] {
		var found bool

		for _, num1 := range numbers[start:end] {
			for _, num2 := range numbers[start:end] {
				if num1+num2 == currentNum {
					found = true
				}
			}
		}

		if !found {
			return currentNum
		}

		start++
		end++
	}
	return -1
}
