package day10

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	input, _ := os.Open("day10/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	r := 1
	cycleCount := 0

	for scanner.Scan() {
		var op string
		var val int
		var opCycles int
		fmt.Sscanf(scanner.Text(), "%s %d", &op, &val)

		if op == "addx" {
			opCycles = 2
		} else if op == "noop" {
			opCycles = 1
		}

		for i := 0; i < opCycles; i++ {
			// pixel rows from 0 to 39
			cycleCount %= 40

			// check sprite position overlap
			if cycleCount == r-1 || cycleCount == r || cycleCount == r+1 {
				fmt.Printf("##")
			} else {
				fmt.Printf("  ")
			}

			// check if end of operation
			if i+1 >= opCycles {
				// add value to register
				r += val
			}

			// if last pixel of the row
			if cycleCount == 39 {
				fmt.Println()
			}
			cycleCount++
		}
	}
}
