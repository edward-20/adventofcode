package main

import (
	"day10/internal/part1"
	"day10/internal/part2"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		1 : part 1 practice
		2 : part 1 actual
		4 : part 2 practice
		8 : part 2 actual
	*/
	args := os.Args[1]

	actual, err := os.Open("input.txt")
	practice, err := os.Open("practice.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer actual.Close()
	defer practice.Close()

	flags, err := strconv.Atoi(args)

	if err != nil {
		fmt.Println(err)
		return
	}

	var bitarray [4]bool

	for i := 3; i >= 0; i-- {
		bit := 1 << i
		if flags-bit >= 0 {
			flags -= bit
			bitarray[i] = true
		}
	}

	if bitarray[0] {
		part1.Part1(practice)
	}
	if bitarray[1] {
		part1.Part1(actual)
	}
	if bitarray[2] {
		part2.Part2(practice)
	}
	if bitarray[3] {
		part2.Part2(actual)
	}

	return
}
