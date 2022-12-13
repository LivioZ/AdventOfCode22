package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	input, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	charStream := []rune(scanner.Text())

	const markerLength = 4

	if len(charStream) < markerLength {
		log.Fatalf("Stream shorter than %v characters", markerLength)
	}

	for i := 0; i < len(charStream)-markerLength; i++ {
		charSet := make(map[rune]bool)
		for j := 0; j < markerLength; j++ {
			charSet[charStream[i+j]] = true
		}
		if len(charSet) == markerLength {
			fmt.Printf("First marker after character %v\n", i+markerLength)
			break
		}
	}
}
