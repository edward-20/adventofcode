package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, error) {
	buf := strings.Fields(line)
	var result []int
	for i := 0; i < len(buf); i++ {
		num, err := strconv.Atoi(buf[i])
		if err != nil {
			return result, err
		}
		result = append(result, num)
	}
	return result, nil
}

func isAllElementsTheSame(l []int) bool {
	item := l[0]
	for i := 0; i < len(l); i++ {
		if l[i] != item {
			return false
		}
	}
	return true
}

func getNextNumber(numbers []int) int {
	// fmt.Println("one invokation")
	layers := make([][]int, 0)

	// initialise the first layer
	firstLayer := make([]int, len(numbers)-1, len(numbers)-1)
	for i := 1; i < len(numbers); i++ {
		firstLayer[i-1] = numbers[i] - numbers[i-1]
	}

	layers = append(layers, firstLayer)

	// check if the previous layer was all same, if not make the next layer
	for i := 0; !isAllElementsTheSame(layers[i]); i++ {
		layer := make([]int, cap(layers[i])-1, cap(layers[i])-1)
		for j := 1; j < cap(layers[i]); j++ {
			layer[j-1] = layers[i][j] - layers[i][j-1]
		}
		layers = append(layers, layer)
	}

	result := layers[len(layers)-1][0]
	// fmt.Println(result)
	// go back up the layers, starting from the second last (the last non-homgeneous layer)
	for i := len(layers) - 2; i >= 0; i-- {
		result = layers[i][len(layers[i])-1] + result
		// fmt.Println(result)
	}
	return numbers[len(numbers)-1] + result
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error when reading file")
		return
	}

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		numbers, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Println("line parse error: unable to convert string to slice of integers")
			return
		} else {
			result += getNextNumber(numbers)
		}
	}

	fmt.Println("Answer is", result)

}
