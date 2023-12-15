package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func waysToBeat (time int, distance int) int {
	numWays := 0
	for i := 1; i < time; i++ {
		if i * (time - i) > distance {
			numWays++
		}
	}
	return numWays
}
func main () {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter time given and distance to beat (as space separated values): ")
		text, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		result := strings.Split(strings.TrimSpace(text), " ")
		t := result[0]
		d := result[1]

		time, terr := strconv.Atoi(t)
		if terr != nil {
			fmt.Println("error converting time")
			break
		}

		dist, derr := strconv.Atoi(d)
		if derr != nil {
			fmt.Println("error converting distance")
			break
		}
		fmt.Println("Ways of beating:", waysToBeat(time, dist))
	}
}
