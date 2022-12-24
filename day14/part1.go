package day14

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y int
}

var cave = make(map[point]bool)
var lastRock point
var unitsOfSandFallen int

func Part1() {
	input, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), " -> ") // positions = ["541,130", "541,122", ...]
		points := make([]point, 0)

		for _, pos := range positions { // points = [point{541,130}, point{541,122}, ...]
			var p point
			fmt.Sscanf(pos, "%d,%d", &p.x, &p.y)
			points = append(points, p)
		}

		// cave's rocks initialization
		for i := 1; i < len(points); i++ {
			start := points[i-1]
			end := points[i]

			// straight line in one direction at a time
			if start.x != end.x {
				// determine direction
				dirX := (end.x - start.x) / int(math.Abs(float64(end.x-start.x))) // 1 or -1
				// set rock positions step by step
				for x := start.x; x != end.x; x += dirX {
					cave[point{x, start.y}] = true
				}
			} else if start.y != end.y {
				dirY := (end.y - start.y) / int(math.Abs(float64(end.y-start.y))) // 1 or -1
				for y := start.y; y != end.y; y += dirY {
					cave[point{start.x, y}] = true
				}
			}
			// because end.x/end.y don't enter for loops
			cave[point{end.x, end.y}] = true

			if end.y > lastRock.y {
				lastRock.x, lastRock.y = end.x, end.y
			}

		}
	}

	sandStart := point{500, 0}
	for willRest(sandStart) {
		unitsOfSandFallen++
	}

	fmt.Printf("Number of sand units came to rest: %v\n", unitsOfSandFallen)
}

func willRest(p point) bool {
	if p.y+1 <= lastRock.y {
		if cave[point{p.x, p.y + 1}] { // down
			if cave[point{p.x - 1, p.y + 1}] { // down-left
				if cave[point{p.x + 1, p.y + 1}] { // down-right
					cave[point{p.x, p.y}] = true
					return true
				} else {
					return willRest(point{p.x + 1, p.y + 1})
				}
			} else {
				return willRest(point{p.x - 1, p.y + 1})
			}
		} else {
			return willRest(point{p.x, p.y + 1})
		}
	} else {
		return false
	}
}
