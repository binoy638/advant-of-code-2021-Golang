package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getIntArrayFromFile(location string) []int {

	var inputs = []int{}
	content, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		input, _ := strconv.Atoi(line)

		inputs = append(inputs, input)

	}
	return inputs
}

func getStringArrayFromFile(location string) []string {

	var inputs = []string{}
	content, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()

		inputs = append(inputs, line)

	}
	return inputs
}

func main() {

	var input = getStringArrayFromFile("input.txt")

	var horizontal, depth int

	for i := 0; i < len(input); i++ {
		direction := input[i][0 : len(input[i])-2]
		value, _ := strconv.Atoi(input[i][len(input[i])-1:])

		if direction == "forward" {
			horizontal += value
		} else {
			if direction == "up" {
				depth -= value
			} else {
				depth += value
			}
		}

	}

	fmt.Println(horizontal * depth)

}
