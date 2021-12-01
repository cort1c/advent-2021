package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var inputs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	prev, result := -1, 0
	for _, input := range inputs {
		if prev != -1 && input > prev {
			result++
		}
		prev = input
	}
	fmt.Printf("result = %d\n", result)
}
