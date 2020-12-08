package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type rule struct {
	Contains map[string]int64
}

func main() {
	f, err := os.Open("./7/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[string]rule)

	for scanner.Scan() {
		rule := rule{Contains: make(map[string]int64)}

		line := scanner.Text()

		words := strings.Split(line, " ")

		for i, word := range words {
			if unicode.IsDigit(rune(word[0])) {

				count, _ := strconv.ParseInt(string(word[0]), 10, 32)

				color := words[i+1] + words[i+2]
				rule.Contains[color] = count
			}
		}

		bagName := words[0] + words[1]

		rules[bagName] = rule
	}

	bagsFound := make(map[string]bool)

	findParents("shinygold", rules, bagsFound)
	fmt.Printf("Total bags that can contain shiny gold: %+v\n", len(bagsFound))

	childrenFound := findChildren("shinygold", rules, 1)
	fmt.Printf("Total children contained in your shiny gold bag: %d", childrenFound-1)
}

func findParents(colorLookingFor string, rules map[string]rule, bagsFound map[string]bool) {
	for s, r := range rules {
		if r.Contains[colorLookingFor] != 0 {
			bagsFound[s] = true
			findParents(s, rules, bagsFound)
		}
	}
}

func findChildren(colorLookingFor string, rules map[string]rule, count int64) int64 {
	var children int64
	rule := rules[colorLookingFor]

	for bag, bagCount := range rule.Contains {
		children += findChildren(bag, rules, count) * bagCount
	}

	return count + children
}
