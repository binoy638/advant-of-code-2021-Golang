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
	input := getStringArrayFromFile("input.txt")
	ones := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	zeroes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var gammaRate, epsilonRate string
	var intGamma, intEpsilon int64

	for _, val := range input {

		for i, bit := range val {
			if string(bit) == "0" {
				zeroes[i] += 1
			}
			if string(bit) == "1" {
				ones[i] += 1
			}
		}
	}

	for i, onebit := range ones {
		if onebit > zeroes[i] {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	if i, err := strconv.ParseInt(gammaRate, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		intGamma = i
	}

	if i, err := strconv.ParseInt(epsilonRate, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		intEpsilon = i
	}

	fmt.Println(intGamma * intEpsilon)

}
