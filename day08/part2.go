package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part2() {
	input, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	forest := make([][]int, 0)
	scenicScore := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		var oneRow []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			oneRow = append(oneRow, 1)
			row = append(row, num)
		}
		forest = append(forest, row)
		scenicScore = append(scenicScore, oneRow)
	}

	for i := range forest {
		for j := range forest[0] {
			// right
			d := 1 // number of visible trees
			for j+d < (len(forest[i])-1) && forest[i][j+d] < forest[i][j] {
				d += 1
			}
			scenicScore[i][j] *= d

			// left
			d = 1
			for j-d > 0 && forest[i][j-d] < forest[i][j] {
				d += 1
			}
			scenicScore[i][j] *= d

			// down
			d = 1
			for i+d < (len(forest)-1) && forest[i+d][j] < forest[i][j] {
				d += 1
			}
			scenicScore[i][j] *= d

			// up
			d = 1
			for i-d > 0 && forest[i-d][j] < forest[i][j] {
				d += 1
			}
			scenicScore[i][j] *= d
		}
	}

	bestScenicScore := 0
	for i := range scenicScore {
		for j := range scenicScore {
			if scenicScore[i][j] > bestScenicScore {
				bestScenicScore = scenicScore[i][j]
			}
		}
	}

	fmt.Printf("Highest scenic score: %v\n", bestScenicScore)
}
