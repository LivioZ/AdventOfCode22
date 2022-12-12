package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type stack struct {
	elements []rune
}

func (s *stack) pop() (r rune) {
	r = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *stack) push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func Part1() {
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
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &num, &from, &to)
		if err != nil {
			continue
		}
		for i := 1; i <= num; i++ {
			stacks[to-1].push(stacks[from-1].pop())
		}
	}

	fmt.Printf("Crates on top of the stacks (crane 9000): ")
	for _, s := range stacks {
		fmt.Printf("%v", string(s.pop()))
	}
	fmt.Println()
}
