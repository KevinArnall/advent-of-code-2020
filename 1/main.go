package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("./1/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numbers []int64

	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		numbers = append(numbers, number)
	}

	find2Naive(numbers)
	find2Optimized(numbers)
	find3Naive(numbers)

}

// O(n^2)
func find2Naive(numbers []int64) {
	for _, num1 := range numbers {
		for _, num2 := range numbers {
			if (num1 + num2) == 2020 {
				product := num1 * num2
				fmt.Printf("Num1: %d, Num2: %d, Product: %d\n", num1, num2, product)
				return
			}
		}
	}
}

// O(n)
func find2Optimized(numbers []int64) {
	occurred := make(map[int64]bool)

	for _, num1 := range numbers {
		num2 := 2020 - num1

		if occurred[num2] {
			product := num1 * num2
			fmt.Printf("Num1: %d, Num2: %d, Product: %d\n", num1, num2, product)
		} else {
			occurred[num1] = true
		}
	}
}

// O(n^3), can optimize using the same method as find2Optimized to get O(n^2)
func find3Naive(numbers []int64) {
	for _, num1 := range numbers {
		for _, num2 := range numbers {
			for _, num3 := range numbers {
				if (num1 + num2 + num3) == 2020 {
					product := num1 * num2 * num3
					fmt.Printf("Num1: %d, Num2: %d, Num3: %d, Product: %d\n", num1, num2, num3, product)
					return
				}
			}
		}
	}
}
