package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	Amount       int
	Instruction  string
	NextPosition int
}

func main() {
	f, err := os.Open("./8/input.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var position int

	instructions := make(map[int]*node)

	for scanner.Scan() {
		line := scanner.Text()

		operation := line[:3]
		amount, _ := strconv.Atoi(line[4:])

		var nextPosition int

		if operation == "jmp" {
			nextPosition = position + amount
		} else {
			nextPosition = position + 1
		}

		instructions[position] = &node{
			NextPosition: nextPosition,
			Amount:       amount,
			Instruction:  operation,
		}

		position++
	}

	acc, _ := part1(instructions)
	fmt.Printf("Value in accumulator for part 1: %d\n", acc)

	acc = part2(instructions)
	fmt.Printf("Value in accumulator for part 2: %d\n", acc)
}

func part1(instructions map[int]*node) (int, error) {
	var position int
	var acc int
	visited := make(map[int]bool)

	for {

		if position == len(instructions) {
			return acc, nil
		}

		if visited[position] {
			break
		}

		visited[position] = true

		if instructions[position].Instruction == "acc" {
			acc += instructions[position].Amount
		}

		position = instructions[position].NextPosition
	}
	return acc, errors.New("couldn't reach the end")
}

func part2(instructions map[int]*node) int {
	var position int

	for {
		node := instructions[position]

		switchInstruction(node, position)

		acc, err := part1(instructions)
		if err == nil {
			return acc
		}

		// Need to make sure to switch it back because it's a pointer
		switchInstruction(node, position)

		position++
	}
}

func switchInstruction(node *node, position int) {
	if node.Instruction == "nop" {
		node.Instruction = "jmp"
		node.NextPosition = position + node.Amount
	} else if node.Instruction == "jmp" {
		node.Instruction = "nop"
		node.NextPosition = position + 1
	}
}
