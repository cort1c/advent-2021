package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const WINDOW_SIZE int = 3

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
	result := 0
	for _, w := range window {
		result += w
	}
	return result
}

func updateWindow(window []int, inputs []int, index int) {
	for i := 0; i < WINDOW_SIZE; i++ {
		window[i] = inputs[index-(WINDOW_SIZE-1-i)]
	}
}

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	prev, result := -1, 0
	window := make([]int, WINDOW_SIZE)
	for i := WINDOW_SIZE - 1; i < len(inputs); i++ {
		updateWindow(window, inputs, i)
		sum := sumWindow(window)
		if prev != -1 && sum > prev {
			result++
		}
		prev = sum
	}
	fmt.Printf("result = %d\n", result)
}
