package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
rock		A  1
paper		B  2
scissors	C  3
*/

/*
lose	X
draw	Y
win		Z
*/

func myPlay(p1, end string) int {
	// lose
	if end == "X" {
		if val[p1] == 1 {
			return 3 + 0
		} else if val[p1] == 2 {
			return 1 + 0
		} else if val[p1] == 3 {
			return 2 + 0
		}
		// draw
	} else if end == "Y" {
		return val[p1] + 3
		// win
	} else if end == "Z" {
		if val[p1] == 1 {
			return 2 + 6
		} else if val[p1] == 2 {
			return 3 + 6
		} else if val[p1] == 3 {
			return 1 + 6
		}
	}
	return 0
}

func Part2() {
	input, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	sum := 0

	for scanner.Scan() {
		p1 := scanner.Text()
		scanner.Scan()
		roundEnd := scanner.Text()

		sum = sum + myPlay(p1, roundEnd)
	}

	fmt.Printf("Total score: %v\n", sum)
}
