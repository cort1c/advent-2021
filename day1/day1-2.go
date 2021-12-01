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

func sumWindow(window []int) int {
	return window[0] + window[1] + window[2]
}

func updateWindow(window []int, inputs []int, index int) {
	window[0] = inputs[index-2]
	window[1] = inputs[index-1]
	window[2] = inputs[index]
}

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	prev := -1
	result := 0
	window := []int{0, 0, 0}
	for i := 2; i < len(inputs); i++ {
		updateWindow(window, inputs, i)
		sum := sumWindow(window)
		if prev != -1 && sum > prev {
			result++
		}
		prev = sum
	}
	fmt.Printf("result = %d\n", result)
}
