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

	var count int

	for scanner.Scan() {
		line := scanner.Text()
		if isPasswordValid(line) {
			count++
		}
	}

	fmt.Printf("Found %d valid passwords", count)
}

type requirements struct {
	Letter string
	Min    int
	Max    int
}

func isPasswordValid(line string) bool {

	values := strings.Split(line, " ")

	letter := string(values[1][0])

	counts := strings.Split(values[0], "-")
	min, _ := strconv.Atoi(counts[0])
	max, _ := strconv.Atoi(counts[1])

	req := requirements{
		Letter: letter,
		Min:    min,
		Max:    max,
	}

	password := values[2]

	count := 0

	for i := range password {
		if string(password[i]) == req.Letter {
			count++
		}
	}

	if count >= req.Min && count <= req.Max {
		return true
	}

	return false
}
