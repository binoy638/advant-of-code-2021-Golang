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
	var inputs = getIntArrayFromFile("input.txt")
	var sliding = []int{}

	count := 0

	for i := 0; i < len(inputs)-2; i++ {
		sum := inputs[i] + inputs[i+1] + inputs[i+2]
		sliding = append(sliding, sum)
	}

	for i := 0; i < len(sliding)-1; i++ {
		if sliding[i] < sliding[i+1] {
			count++
		}
	}

	fmt.Println(count)

}
