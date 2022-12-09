package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part1() {
	input, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	sum, max := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Not a number\n", err)
			}
			sum = sum + num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Max number of calories: %v\n", max)
}
