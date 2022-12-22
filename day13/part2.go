package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func Part2() {
	input, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var packets []any

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

		packets = append(packets, p1, p2)

		scanner.Scan() // skip empty line
	}

	// add divider packets
	div1 := []any{[]any{float64(2)}}
	div2 := []any{[]any{float64(6)}}
	packets = append(packets, div1, div2)

	// sort packets
	sort.Slice(packets, func(i, j int) bool {
		return isOrdered(packets[i], packets[j]) == 1
	})

	res := 1

	// find dividers
	for i, p := range packets {
		if isOrdered(p, div1) == 0 || isOrdered(p, div2) == 0 {
			res *= i + 1
		}
	}

	fmt.Printf("Product of indexes of divider packets after sorting: %v\n", res)
}
