package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	sum := 0

	input, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		// "82-82,8-83"
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)

		if (startSecond >= startFirst && endSecond <= endFirst) ||
			(startFirst >= startSecond && endFirst <= endSecond) {
			sum++
		}
	}
	fmt.Printf("%v assignment pairs fully contain the other\n", sum)
}
