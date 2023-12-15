package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

type Mapping struct {
	SourceString string
	DestString string
	SDR [][3]int
}

func compare (a [3]int, b [3]int) int {
	if a[1] < b[1] {
		return -1
	} else if a[1] == b[1] {
		return 0
	} else {
		return 1
	}
}

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// get the initial seeds
	scanner.Scan()
	initialSeedsStrings := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
	initialSeedsStrings = initialSeedsStrings[1:]

	initialSeeds := make([]int, 0)
	for _, v := range initialSeedsStrings {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		initialSeeds = append(initialSeeds, num)
	}

	mappings := make([]Mapping, 0)

	// for each mapping
	for scanner.Scan() {
		mapp := Mapping{}

		line := scanner.Text()
		// if its a line describing the mapping
		if strings.Contains(line, "map") {
			mappingStrings := strings.Split(line, " ")[0]
			sdStrings := strings.Split(mappingStrings, "-")
			mapp.SourceString = sdStrings[0]
			mapp.DestString = sdStrings[2]

			sdr := make([][3]int, 0)
			// for each line detailing a numeric mapping
			for scanner.Scan() {
				l := scanner.Text()
				// if we've reached the end of the mappings
				if l == "" {
					break
				}
				mapping := strings.Split(l, " ")

				d, err := strconv.Atoi(mapping[0])
				if err != nil {
					panic(err)
				}

				s, err := strconv.Atoi(mapping[1])
				if err != nil {
					panic(err)
				}

				r, err := strconv.Atoi(mapping[2])
				if err != nil {
					panic(err)
				}

				thisSdr := [3]int{s, d, r}
				sdr = append(sdr, thisSdr)
			}

			mapp.SDR = sdr

			mappings = append(mappings, mapp)

		}


	}

	locations := make([]int, 0, len(initialSeeds))
	// for each seed
	for _, seed := range initialSeeds {
		fmt.Println()
		target := seed
		for _, mapping := range mappings {
			// iterate through SDR to find the element of the array
			// that pertains to the current target
			found := false
			for _, singleMap := range mapping.SDR {
				src := singleMap[0]
				dst := singleMap[1]
				ran := singleMap[2]

				if target >= src && target < src + ran {
					fmt.Println(mapping.SourceString, mapping.DestString, singleMap)
					fmt.Println(target)
					target = dst + (target - src)
					fmt.Println(target)
					found = true
					break
				}
			}
			if !found {
				fmt.Println(mapping.SourceString, mapping.DestString, "identity map")
				fmt.Println(target)
				fmt.Println(target)
				// leave the target as is
			}
		}
		locations = append(locations, target)
	}

	lowest_location := locations[0]
	for _, v := range locations[1:] {
		if v < lowest_location {
			lowest_location = v
		}
	}
	fmt.Println()
	fmt.Println(lowest_location)
}
