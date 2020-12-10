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
	fmt.Printf("Answer for part 1: %d\n", ans)

	ans = part2(numbers, ans)
	fmt.Printf("Answer for part 2: %d\n", ans)

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

func part2(numbers []int, ans int) int {
	start := 0
	end := 1

	for end != len(numbers) {
		total := 0
		min := ans
		max := 0
		for _, number := range numbers[start:end] {
			if number > max {
				max = number
			}
			if number < min {
				min = number
			}
			total += number
		}
		if total == ans {
			return min + max
		} else if total > ans {
			start++
		} else if total < ans {
			end++
		}
	}

	return -1
}
