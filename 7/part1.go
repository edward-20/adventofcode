package main

import (
	"bufio"
	"os"
	"strings"
	"slices"
	"strconv"
	"fmt"
)

var cardStrength = map[byte]int {
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func comparingHands(a [2]string, b [2]string) int {
	for i := range a {
		if cardStrength[a[0][i]] > cardStrength[b[0][i]] {
			return 1
		} else if cardStrength[a[0][i]] == cardStrength[b[0][i]] {
			continue
		} else {
			return -1
		}
	}

	return 0
}

func totalHands(hands [][2]string, rank *int) int{
	total := 0
	for _, hand := range hands {
		bid, err := strconv.Atoi(hand[1])
		if err != nil {
			panic(err)
		}
		total += bid * (*rank)
		fmt.Println(hand[0]," ", bid, "*", *rank, "=", total)
		*rank++
	}
	return total
}

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)	

	hands := make([][2]string, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var both [2]string 
		l := strings.Split(line, " ")
		both[0] = l[0]
		both[1] = l[1]
		hands = append(hands, both)
	}

	five := make([][2]string, 0, 0)
	four := make([][2]string, 0, 0)
	fullHouse := make([][2]string, 0, 0)
	three := make([][2]string, 0, 0)
	twoPair := make([][2]string, 0, 0)
	onePair := make([][2]string, 0, 0)
	highCard := make([][2]string, 0, 0)

	for _, hand := range hands {
		cards := hand[0]
		// bid := hand[1]

		parsedHand := make(map[rune]int)

		for _, card := range cards {
			parsedHand[card]++
		}

		switch len(parsedHand) {
			case 1:
				five = append(five, hand)
			case 2:
				// full house or four
				isFullHouse := false
				for _, value := range parsedHand {
					if value == 2 {
						isFullHouse = true
						break
					}
				}
				if isFullHouse {
					fullHouse = append(fullHouse, hand)
				} else {
					four = append(four, hand)
				}
			case 3:
				// twoPair or three
				isThree := false
				for _, v := range parsedHand {
					if v == 3 {
						isThree = true
						break
					}
				}
				if isThree {
					three = append(three, hand)
				} else {
					twoPair = append(twoPair, hand)
				}
			case 4:
				// onePair
				onePair = append(onePair, hand)
			case 5:
				// highCard
				highCard = append(highCard, hand)

		}
	}
	slices.SortFunc(five, comparingHands)
	slices.SortFunc(four, comparingHands)
	slices.SortFunc(three, comparingHands)
	slices.SortFunc(fullHouse, comparingHands)
	slices.SortFunc(twoPair, comparingHands)
	slices.SortFunc(onePair, comparingHands)
	slices.SortFunc(highCard, comparingHands)

	total := 0
	rank := 1
	total += totalHands(onePair, &rank)
	total += totalHands(twoPair, &rank)
	total += totalHands(three, &rank)
	total += totalHands(fullHouse, &rank)
	total += totalHands(four, &rank)
	total += totalHands(five, &rank)
	fmt.Println(total)

}
