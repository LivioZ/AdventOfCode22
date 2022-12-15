package day09

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type pos struct {
	i, j int
}

func (t *pos) moveTo(h *pos) {
	if (math.Abs(float64(h.i - t.i))) > (math.Abs(float64(h.j - t.j))) {
		// higher vertical distance
		t.j = h.j      // same horizontal position
		if h.i > t.i { // head under tail
			t.i = h.i - 1
		} else if h.i < t.i { // head over tail
			t.i = h.i + 1
		}
	} else if (math.Abs(float64(h.j - t.j))) > (math.Abs(float64(h.i - t.i))) {
		// higher horizontal distance
		t.i = h.i      // same vertical position
		if h.j > t.j { // head on the right of tail
			t.j = h.j - 1
		} else if h.j < t.j { // head on the left of tail
			t.j = h.j + 1
		}
	}
}

func Part1() {
	input, _ := os.Open("day09/input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var direction string
	var steps int
	var grid = make(map[pos]bool)
	grid[pos{0, 0}] = true
	var h, t pos

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &steps)
		for steps > 0 {
			switch direction {
			case "L":
				h.j -= 1
			case "R":
				h.j += 1
			case "U":
				h.i -= 1
			case "D":
				h.i += 1
			}
			steps--
			t.moveTo(&h)
			grid[t] = true
		}
	}
	fmt.Printf("Number of positions the tail visited: %v\n", len(grid))
}
