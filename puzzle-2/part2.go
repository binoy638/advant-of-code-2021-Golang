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

	var horizontal, depth, aim int

	for _, val := range input {
		direction := val[0 : len(val)-2]
		value, _ := strconv.Atoi(val[len(val)-1:])

		if direction == "forward" {
			horizontal += value
			depth += aim * value
		} else {
			if direction == "up" {
				aim -= value
			} else {
				aim += value
			}
		}

	}

	fmt.Println(horizontal * depth)

}
