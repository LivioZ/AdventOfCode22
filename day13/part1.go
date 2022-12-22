package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Part1() {
	input, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var indexSum int

	for i := 1; scanner.Scan(); i++ {
		// package 1
		line := []byte(scanner.Text())
		var p1 []any // each packet (line) is always a list
		err = json.Unmarshal(line, &p1)
		if err != nil {
			log.Fatal(err)
		}

		// package 2
		scanner.Scan()
		line = []byte(scanner.Text())
		var p2 []any // each packet (line) is always a list
		err = json.Unmarshal(line, &p2)
		if err != nil {
			log.Fatal(err)
		}

		if isOrdered(p1, p2) == 1 {
			indexSum += i
		}

		scanner.Scan() // skip empty line
	}
	fmt.Printf("Sum of indexes of pairs in right order: %v\n", indexSum)
}

func isOrdered(left, right any) int {
	// determine if left and right are slices or numbers
	// numbers extracted with json.Unmarshal() are type float64
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			// both numbers
			a := left.(float64)
			b := right.(float64)
			if a < b {
				return 1
			}
			if a == b {
				return 0
			}
			return -1
		case []any:
			// make left a slice
			left = []any{left}
		}
	case []any:
		switch right.(type) {
		case float64:
			// make left a slice
			right = []any{right}
		}
	}

	// now we have two slices
	var l, r = left.([]any), right.([]any)
	var i int
	for i < len(l) && i < len(r) {
		cond := isOrdered(l[i], r[i])
		if cond == 1 {
			return 1
		}
		if cond == -1 {
			return -1
		}

		i++
	}

	if i == len(l) {
		if len(l) == len(r) {
			// no comparison made order decision
			return 0
		} else {
			// left run out of items first
			return 1
		}
	}

	// right run out of items first
	return -1
}
