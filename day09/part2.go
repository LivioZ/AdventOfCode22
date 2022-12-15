package day09

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	input, _ := os.Open("day09/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var direction string
	var steps int
	var tailVisited = make(map[coord]bool)
	var rope = make([]coord, 10) // rope of 10 knots
	tailVisited[rope[9]] = true  // tail starts at coord{0, 0}

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &steps)
		for ; steps > 0; steps-- {
			switch direction {
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			case "U":
				rope[0].y--
			case "D":
				rope[0].y++
			}

			for i := 1; i < len(rope); i++ {
				// rope[i]: node we are considering
				// rope[i-1]: its head
				// same logic as part 1
				rope[i].moveTo(&rope[i-1])
			}
			tailVisited[rope[9]] = true
		}
	}
	fmt.Printf("Number of coorditions the tail visited: %v\n", len(tailVisited))
}
