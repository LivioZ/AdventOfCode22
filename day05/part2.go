package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func Part2() {
	input, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	line := []rune(scanner.Text())
	stacks := make([]stack, (len(line)+1)/4)
	for scanner.Text()[:2] != " 1" {
		line = []rune(scanner.Text())
		for index, char := range line {
			if unicode.IsLetter(char) {
				stacks[index/4].addToBottom(char)
			}
		}
		scanner.Scan()
	}

	for scanner.Scan() {
		line := scanner.Text()
		var num, from, to int
		var tempStack stack
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &num, &from, &to)
		if err != nil {
			continue
		}
		for i := 1; i <= num; i++ {
			tempStack.elements = append([]rune{stacks[from-1].pop()}, tempStack.elements...)
		}
		stacks[to-1].elements = append(stacks[to-1].elements, tempStack.elements...)
	}

	fmt.Printf("Crates on top of the stacks (crane 9001): ")
	for _, s := range stacks {
		fmt.Printf("%v", string(s.pop()))
	}
	fmt.Println()
}
