package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("./6/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	group := make(map[rune]bool)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += len(group)
			group = make(map[rune]bool)
			continue
		}

		for _, letter := range line {
			group[letter] = true
		}
	}

	total += len(group)

	fmt.Printf("Total questions answered: %d\n", total)
}
