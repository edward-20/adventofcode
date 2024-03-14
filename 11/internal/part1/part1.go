package part1

import (
	"bufio"
	"fmt"
	"os"
)

/*
We will use Floyd Warshall to find all pairs of shortest paths
*/
func Part1(file *os.File) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileLines := make([][]rune, 0)

	for fileScanner.Scan() {
		fileLines = append(fileLines, []rune(fileScanner.Text()))
	}

	for i, line := range fileLines {
		for j, ch := range line {
			if ch == '#' {
				fileLines[i][j] = 'a'
			}

		}
	}

	for _, line := range fileLines {
		fmt.Println(string(line))
	}

}
