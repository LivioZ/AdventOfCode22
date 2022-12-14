package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

func Part1() {
	input, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	forest := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		forest = append(forest, row)
	}

	isVisible := make(map[point]bool)

	//Horizontal view, from left and right simultaneously
	maxLeft := make([]int, len(forest))
	maxRight := make([]int, len(forest))
	for i := range forest {
		for j := range forest[0] {
			if j == 0 {
				maxLeft[i] = forest[i][j]
				maxRight[i] = forest[i][len(forest[0])-1]
				isVisible[point{i, j}] = true
				isVisible[point{i, len(forest[0]) - 1}] = true
				continue
			}
			if forest[i][j] > maxLeft[i] {
				isVisible[point{i, j}] = true
				maxLeft[i] = forest[i][j]
			}
			if forest[i][len(forest[0])-1-j] > maxRight[i] {
				isVisible[point{i, len(forest[0]) - 1 - j}] = true
				maxRight[i] = forest[i][len(forest[0])-1-j]
			}
		}
	}

	//Vertical view, from top and bottom simultaneously
	maxTop := make([]int, len(forest))
	maxDown := make([]int, len(forest))
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if i == 0 {
				maxTop[j] = forest[i][j]
				maxDown[j] = forest[len(forest)-1][j]
				isVisible[point{i, j}] = true
				isVisible[point{len(forest) - 1, j}] = true
				continue
			}
			if forest[i][j] > maxTop[j] {
				isVisible[point{i, j}] = true
				maxTop[j] = forest[i][j]
			}
			if forest[len(forest)-1-i][j] > maxDown[j] {
				isVisible[point{len(forest) - 1 - i, j}] = true
				maxDown[j] = forest[len(forest)-1-i][j]
			}
		}
	}
	fmt.Printf("Visible trees from outside the grid: %v\n", len(isVisible))
}
