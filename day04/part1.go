package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// Checks if a range contains the other (interchangeably).
//
// Ranges are in the form of a slice.
//
// Example: ["82-82", "8-83"]
func containedRanges(ranges []string) bool {
	range1 := strings.Split(ranges[0], "-") // ["82", "82"]
	range2 := strings.Split(ranges[1], "-") // ["8", "83"]
	n1 := toInt(range1[0])
	n2 := toInt(range1[1])
	n3 := toInt(range2[0])
	n4 := toInt(range2[1])

	/*
			n1---n2		 n1-n2
			 n3-n4		n3---n4
	*/
	return (n3 >= n1 && n4 <= n2) || (n1 >= n3 && n2 <= n4)
}

func Part1() {
	sum := 0

	input, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text() // "82-82,8-83"
		ranges := strings.Split(line, ",") // ["82-82", "8-83"]
		
		if containedRanges(ranges) {
			sum++
		}
	}

	fmt.Printf("%v assignment pairs fully contain the other\n", sum)
}