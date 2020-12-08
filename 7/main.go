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
	Contains []bag
}

type bag struct {
	Color string
	Count int64
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
		rule := rule{}

		line := scanner.Text()

		words := strings.Split(line, " ")

		for i, word := range words {
			if unicode.IsDigit(rune(word[0])) {

				count, _ := strconv.ParseInt(string(word[0]), 10, 32)

				b := bag{
					Color: words[i+1] + words[i+2],
					Count: count,
				}
				rule.Contains = append(rule.Contains, b)
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
		for _, bag := range r.Contains {
			if bag.Color == colorLookingFor {
				bagsFound[s] = true
				findParents(s, rules, bagsFound)
			}
		}
	}
}

func findChildren(colorLookingFor string, rules map[string]rule, count int64) int64 {
	for s, r := range rules {
		if colorLookingFor == s {
			var children int64

			for _, bag := range r.Contains {
				children += findChildren(bag.Color, rules, count) * bag.Count
			}

			count += children
		}
	}
	return count
}
