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

	group := make(map[rune]int)
	totalP1 := 0
	peopleInGroup := 0
	answeredByEveryoneInGroup := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totalP1 += len(group)
			for _, q := range group {
				if q == peopleInGroup {
					answeredByEveryoneInGroup++
				}
			}

			group = make(map[rune]int)
			peopleInGroup = 0
			continue
		}

		peopleInGroup++

		for _, letter := range line {
			group[letter]++
		}
	}

	totalP1 += len(group)
	for _, q := range group {
		if q == peopleInGroup {
			answeredByEveryoneInGroup++
		}
	}

	fmt.Printf("Total questions answered for part 1: %d\n", totalP1)
	fmt.Printf("Total questions answered for part 2: %d\n", answeredByEveryoneInGroup)
}
