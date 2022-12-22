package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	packets = append(packets, []any{[]any{float64(2)}}, []any{[]any{float64(6)}})

	// insertion sort
	for i := 1; i < len(packets); i++ {
		v := packets[i]
		j := i - 1
		for j >= 0 && !less(packets[j], v) {
			packets[j+1] = packets[j]
			j = j - 1
		}
		packets[j+1] = v
	}

	res := 1

	for i, packet := range packets {
		if p, ok := packet.([]any); ok {
			if len(p) == 1 {
				if p1, ok := p[0].([]any); ok {
					if len(p1) == 1 {
						if p2, ok := p1[0].(float64); ok {
							if p2 == 2.0 || p2 == 6.0 {
								res *= i + 1
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Product of indexes of divider packets after sorting: %v\n", res)
}

func less(left, right any) bool {
	if isOrdered(left, right) == 1 {
		return true
	} else {
		return false
	}
}
