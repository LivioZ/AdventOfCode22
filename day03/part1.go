package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	sum := 0

	input, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rucksack := scanner.Text()
		comp1 := scanner.Text()[:len(rucksack)/2]
		comp2 := scanner.Text()[len(rucksack)/2:]

		set := make(map[rune]bool)

		for _, char := range comp1 {
			set[char] = true
		}

		for _, char := range comp2 {
			if set[char] {
				if char >= 65 && char <= 90 {
					sum = sum + (int(char) - 38)
				} else if char >= 97 && char <= 122 {
					sum = sum + (int(char) - 96)
				}
				break
			}
		}
	}

	fmt.Printf("Sum of priorities: %v\n", sum)
}
