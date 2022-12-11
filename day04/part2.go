package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part2() {
	numberOfOverlaps := 0

	input, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int

		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)

		if (startSecond <= endFirst && endSecond >= startFirst) ||
			(startFirst <= endSecond && endFirst >= startSecond) {
			numberOfOverlaps++
		}
	}
	fmt.Printf("%v assignment pairs overlap\n", numberOfOverlaps)
}
