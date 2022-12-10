package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func intersect(s1, s2 map[rune]bool) map[rune]bool {
	s_intersect := make(map[rune]bool)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for key := range s1 {
		if s2[key] {
			s_intersect[key] = true
		}
	}

	return s_intersect
}

func Part2() {
	sum := 0

	input, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rucksack1 := scanner.Text()
		scanner.Scan()
		rucksack2 := scanner.Text()
		scanner.Scan()
		rucksack3 := scanner.Text()

		set1 := make(map[rune]bool)
		set2 := make(map[rune]bool)
		set3 := make(map[rune]bool)

		for _, char := range rucksack1 {
			set1[char] = true
		}
		for _, char := range rucksack2 {
			set2[char] = true
		}
		for _, char := range rucksack3 {
			set3[char] = true
		}

		intersection := intersect(intersect(set1, set2), set3)

		for k := range intersection {
			if k >= 65 && k <= 90 {
				sum = sum + (int(k) - 38)
			} else if k >= 97 && k <= 122 {
				sum = sum + (int(k) - 96)
			}
			break
		}
	}

	fmt.Printf("Sum of priorities of groups: %v\n", sum)
}
