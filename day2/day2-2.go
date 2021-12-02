package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Amount    int
}

func readInput() ([]Command, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var inputs []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) != 2 {
			continue
		}
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		inputs = append(inputs, Command{parts[0], amount})
	}
	return inputs, nil
}

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	horizontal, vertical, aim := 0, 0, 0
	for _, input := range inputs {
		switch input.Direction {
		case "forward":
			horizontal += input.Amount
			vertical += input.Amount * aim
		case "down":
			aim += input.Amount
		case "up":
			aim -= input.Amount
		}
	}
	fmt.Printf("result: horizontal = %d, vertical = %d, aim = %d, final = %d\n", horizontal, vertical, aim, horizontal*vertical)
}
