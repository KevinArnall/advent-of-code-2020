package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./10/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	adapters := make([]int, 0)

	for scanner.Scan() {
		adapter, _ := strconv.Atoi(scanner.Text())

		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)

	part1(adapters)
}

func part1(adapters []int) {
	diff1 := 0
	diff3 := 0
	currentVoltage := 0

	for _, adapter := range adapters {
		if adapter-currentVoltage == 1 {
			diff1++
		} else if adapter-currentVoltage == 3 {
			diff3++
		}
		currentVoltage = adapter
	}

	// Don't forget the last jump to the device
	diff3++

	fmt.Printf("Diff1: %d, Diff3: %d, Product: %d\n", diff1, diff3, diff1*diff3)
}
