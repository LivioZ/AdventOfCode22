package day10

import (
	"bufio"
	"fmt"
	"os"
)

var targetCycles map[int]bool = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

func Part1() {
	input, _ := os.Open("day10/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	r := 1
	cycleCount := 1
	signalStrength := 0

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
			if targetCycles[cycleCount] {
				signalStrength += cycleCount * r
			}
			// check if end of operation
			if i+1 >= opCycles {
				// add value to register
				r += val
			}
			cycleCount++
		}
	}
	fmt.Printf("Sum of signal strengths during specific cycles: %v\n", signalStrength)
}
