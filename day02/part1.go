package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
rock		A  X  1
paper		B  Y  2
scissors	C  Z  3
*/

var val = map[string]int{
	"A": 1,
	"X": 1,
	"B": 2,
	"Y": 2,
	"C": 3,
	"Z": 3,
}

func winner(p1, p2 string) int {
	// I'm p2
	if val[p1] == val[p2] {
		return val[p2] + 3
	} else if (val[p1] == 1 && val[p2] == 2) || (val[p1] == 2 && val[p2] == 3) || (val[p1] == 3 && val[p2] == 1) {
		return val[p2] + 6
	} else {
		return val[p2] + 0
	}
}

func Part1() {
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
		p2 := scanner.Text()

		sum = sum + winner(p1, p2)
	}

	fmt.Printf("Total score: %v\n", sum)
}
