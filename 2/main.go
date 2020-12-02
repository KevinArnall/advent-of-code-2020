package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("./2/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var p1Count int

	var p2Count int

	for scanner.Scan() {
		line := scanner.Text()
		if isPasswordValid(line) {
			p1Count++
		}
		if isPasswordValidP2(line) {
			p2Count++
		}
	}

	fmt.Printf("Found %d valid passwords for part 1\n", p1Count)
	fmt.Printf("Found %d valid passwords for part 2\n", p2Count)
}

func isPasswordValid(line string) bool {

	values := strings.Split(line, " ")

	letter := string(values[1][0])

	counts := strings.Split(values[0], "-")
	min, _ := strconv.Atoi(counts[0])
	max, _ := strconv.Atoi(counts[1])

	password := values[2]

	count := 0

	for i := range password {
		if string(password[i]) == letter {
			count++
		}
	}

	if count >= min && count <= max {
		return true
	}

	return false
}

func isPasswordValidP2(line string) bool {

	values := strings.Split(line, " ")

	letter := string(values[1][0])

	counts := strings.Split(values[0], "-")
	position1, _ := strconv.Atoi(counts[0])
	position2, _ := strconv.Atoi(counts[1])

	// Have to account for array index being one less
	position1--
	position2--

	password := values[2]

	if (string(password[position1]) == letter && string(password[position2]) != letter) ||
		(string(password[position2]) == letter && string(password[position1]) != letter) {
		return true
	}

	return false
}
