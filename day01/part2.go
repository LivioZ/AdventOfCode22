package day01

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Part2() {
	input, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	sum := 0
	first, second, third := -math.MaxInt32, -math.MaxInt32, -math.MaxInt32

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > first {
				third = second
				second = first
				first = sum
			} else if sum > second {
				third = second
				second = sum
			} else if sum > third {
				third = sum
			}
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Not a number\n", err)
			}
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum = first + second + third
	fmt.Printf("Sum of the three largest calories carried: %v\n", sum)
}
