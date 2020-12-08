package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type rule struct {
	Contains []string
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
				rule.Contains = append(rule.Contains, words[i+1]+words[i+2])
			}
		}

		bagName := words[0] + words[1]

		rules[bagName] = rule

	}
	//fmt.Printf("%+v", rules)

	bagsFound := make(map[string]bool)

	find("shinygold", rules, bagsFound)

	//fmt.Printf("%+v\n", bagsFound)
	fmt.Printf("Total bags that can contain shiny gold: %+v", len(bagsFound))
}

func find(bag string, rules map[string]rule, bagsFound map[string]bool) {

	for s, r := range rules {
		for _, bagName := range r.Contains {
			if bagName == bag {
				//fmt.Printf("Child: %s Parent: %s\n", bag, s)
				bagsFound[s] = true
				find(s, rules, bagsFound)
			}
		}
	}
	//fmt.Printf("No more parents found for %s\n", bag)
}
