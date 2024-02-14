package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseSlice(s []int) {
	// Get the length of the slice
	length := len(s)

	// Iterate only half of the slice length
	for i := 0; i < length/2; i++ {
		// Swap elements from both ends
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// refactor code organisation to have main call part1 and part2 at the same time mebe
func getNextNumberPart2(numbers []int) int {
	reverseSlice(numbers)
	return getNextNumber(numbers)
}

func part2() {
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
			result += getNextNumberPart2(numbers)
		}
	}

	fmt.Println("Answer for part 2: ", result)
}
