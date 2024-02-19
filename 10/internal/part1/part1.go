package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coordinate struct {
	x int
	y int
}

var xMax int = 0
var yMax int = 0
var floor []string

func isInBounds(c coordinate) bool {
	return c.x >= 0 && c.x < xMax && c.y >= 0 && c.y < yMax
}

// return the coordinate of the next cell and whether we're at "S" again or not
func getNext(curr coordinate, prev coordinate) (coordinate, bool) {
	// Note: will need to use global xMax, yMax and floor

	// what is current character
	var next coordinate
	fmt.Println(string(floor[curr.y][curr.x]))
	switch floor[curr.y][curr.x] {
	case '|':
		if prev.y == curr.y-1 {
			next = coordinate{curr.x, curr.y + 1}
		} else {
			next = coordinate{curr.x, curr.y - 1}
		}
	case '-':
		if prev.x == curr.x-1 {
			next = coordinate{curr.x + 1, curr.y}
		} else {
			next = coordinate{curr.x - 1, curr.y}
		}
	case 'L':
		if prev.x == curr.x+1 {
			next = coordinate{curr.x, curr.y - 1}
		} else {
			next = coordinate{curr.x + 1, curr.y}
		}
	case 'J':
		if prev.x == curr.x-1 {
			next = coordinate{curr.x, curr.y - 1}
		} else {
			next = coordinate{curr.x - 1, curr.y}
		}
	case 'F':
		if prev.x == curr.x {
			next = coordinate{curr.x + 1, curr.y}
		} else {
			next = coordinate{curr.x, curr.y + 1}
		}
	case '7':
		if prev.x == curr.x-1 {
			next = coordinate{curr.x, curr.y + 1}
		} else {
			next = coordinate{curr.x - 1, curr.y}
		}

	}
	if !isInBounds(next) {
		fmt.Println("error: we could not find the next pipe ", "curr:", curr, ";next:", next)
		return coordinate{0, 0}, true
	}
	return next, (floor[next.y][next.x] == 'S')
}

func Part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	lineWithStart := 0

	floor = make([]string, 0)

	// bespoke first iteration to get the xMax
	scanner.Scan()
	line := scanner.Text()
	xMax = len(line)
	if strings.Contains(line, "S") {
		lineWithStart = 0
	}
	floor = append(floor, line)

	// subsequent iterations to get the floor
	i := 1
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "S") {
			lineWithStart = i
		}
		floor = append(floor, line)
		i++
	}

	yMax = i

	// get the position of the start (which'll be our previous)
	y := lineWithStart
	x := strings.Index(floor[lineWithStart], "S")

	// positions where our current pipe may be
	adjacentsToStart := make(map[string]bool)
	if (isInBounds(coordinate{x, y - 1})) {
		adjacentsToStart["top"] = true
	}
	if (isInBounds(coordinate{x - 1, y})) {
		adjacentsToStart["left"] = true
	}
	if (isInBounds(coordinate{x + 1, y})) {
		adjacentsToStart["right"] = true
	}
	if (isInBounds(coordinate{x, y + 1})) {
		adjacentsToStart["bottom"] = true
	}

	prev := coordinate{x, y}
	var curr coordinate
	for k, _ := range adjacentsToStart {
		switch k {
		case "top":
			top := floor[y-1][x]
			if top == '|' || top == 'F' || top == '7' {
				curr = coordinate{x, y - 1}
			}
		case "left":
			left := floor[y][x-1]
			if left == '-' || left == 'L' || left == 'F' {
				curr = coordinate{x, y - 1}
			}
		case "right":
			right := floor[y][x+1]
			if right == '-' || right == 'J' || right == '7' {
				curr = coordinate{x, y + 1}
			}
		case "bottom":
			left := floor[y-1][x]
			if left == '|' || left == 'L' || left == 'J' {
				curr = coordinate{x, y - 1}
			}
		default:
			continue
		}
	}

	fmt.Println(curr, " is the curr")
	fmt.Println(prev, " is the start aka previous")
	dist := 1
	for {
		next, isBackAtStart := getNext(curr, prev)
		if isBackAtStart {
			break
		} else {
			prev = curr
			curr = next
			dist++
		}
		fmt.Println(prev, "->", curr)
	}

	if dist%2 == 1 {
		dist = dist/2 + 1
	} else {
		dist = dist / 2
	}
	fmt.Println("dist: ", dist)

}
