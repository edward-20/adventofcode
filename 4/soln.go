package main

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"sort"
	"os"
	"strconv"
)

func binSearch(query int, array []int, l int, r int) bool {
	if l > r {
		return false
	}
	// base case
	if l == r {
		if query != array[l] {
			return false
		}
		return true
	}

	halfway := int((r + l) / 2)

	if query == array[halfway] {
		return true
	} else if query > array[halfway]{
		return binSearch(query, array, halfway + 1, r)
	} else {
		return binSearch(query, array, l, halfway - 1)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// get the lines
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	biggestCard := len(lines)
	cardsProduced := make([][]int, 0, biggestCard)

	for i := 0; i < biggestCard; i++ {
		cardsProduced = append(cardsProduced, make([]int, 0))
	}

	// compute for each card the corresponding cards that would be won from
	// them
	for i, line := range lines {
		numbers := strings.Split(line, ":")[1]
		currentGame := int(i + 1)


		ourAndWinningNumbers := strings.Split(numbers, "|")
		winningNumbersStrings := strings.Fields(strings.TrimSpace(ourAndWinningNumbers[0]))
		ourNumbersStrings := strings.Fields(strings.TrimSpace(ourAndWinningNumbers[1]))

		// winning
		winningNumbers := make([]int, len(winningNumbersStrings))
		for i, v := range winningNumbersStrings {
			n, err := strconv.Atoi(v) 
			if err != nil {
				panic(err)
			}
			winningNumbers[i] = n
		}

		// ours 
		ourNumbers := make([]int, len(ourNumbersStrings))
		for i, v := range ourNumbersStrings {
			n, err := strconv.Atoi(v) 
			if err != nil {
				panic(err)
			}
			ourNumbers[i] = n
		}

		sort.Ints(winningNumbers)
		sort.Ints(ourNumbers)

		score := 0
		// for each of our numbers do a binary search for the number on
		// winning numbers
		for _, v := range ourNumbers {
			if binSearch(v, winningNumbers, 0, len(winningNumbers) - 1) {
				score++
			}
		}

		curr := cardsProduced[currentGame - 1]
		for i:= currentGame + 1; i < currentGame + 1 + score && i < biggestCard; i++ {
			curr = append(curr, i)
		}
		cardsProduced[currentGame - 1] = curr
		fmt.Println("Card", currentGame, "has", score, "matching numbers", "so", cardsProduced[currentGame - 1])

	}

	// make two data structures: cardsUsed and cardsWon
	cardsWon := make([]int, 0)
	cardsUsed := make([]int, 0)

	// add the produced cards to cardsWon
	for _, producedCardsList := range cardsProduced {
		for _, v := range producedCardsList {
			cardsWon = append(cardsWon, v)
		}
	}

	// add the originals to cardsUsed
	for i := range lines {
		cardsUsed = append(cardsUsed, i + 1)
	}
	fmt.Println(cardsUsed)

	for len(cardsWon) > 0 {
		// get the top of cardsWon
		popped := cardsWon[0]
		cardsWon = cardsWon[1:]
		// append to cardsWon the cards you got from popped
		cardsWon = append(cardsWon, cardsProduced[popped - 1]...)
		// append to cardsUsed the popped
		cardsUsed = append(cardsUsed, popped)
	}


	fmt.Println(len(cardsUsed))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

