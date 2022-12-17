package day09

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y float64
}

func (t *coord) isTouching(h *coord) bool {
	return math.Abs(float64(h.x-t.x)) <= 1 && math.Abs(float64(h.y-t.y)) <= 1
}

func (t *coord) moveTo(h *coord) {
	if !t.isTouching(h) {
		var movx, movy float64
		if t.x == h.x { // if same x, x never changes
			movx = 0
		} else { // else we need to figure out the direction
			// h.x - t.x is the direction (positive or negative)
			// we normalize its value to 1 or -1
			// because we only have to move by at most 1 unit
			movx = (h.x - t.x) / math.Abs(float64(h.x-t.x))
		}

		// same for y
		if t.y == h.y {
			movy = 0
		} else {
			movy = (h.y - t.y) / math.Abs(float64(h.y-t.y))
		}

		// update the tail position
		t.x += movx
		t.y += movy
	}
}

func Part1() {
	input, err := os.Open("day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var direction string
	var steps int
	var tailVisited = make(map[coord]bool)
	tailVisited[coord{0, 0}] = true
	var h, t coord

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &steps)
		// one step at a time to get every position the tail goes
		for steps > 0 {
			switch direction {
			case "L":
				h.x--
			case "R":
				h.x++
			case "U":
				h.y++
			case "D":
				h.y--
			}
			steps--
			t.moveTo(&h)
			tailVisited[t] = true
		}
	}
	fmt.Printf("Number of positions the tail visited: %v\n", len(tailVisited))
}
