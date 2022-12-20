package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// the idea is to start from E instead of S
// and look for the nearest a
func Part2() {
	input, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var start point

	visited := make(map[point]bool)
	distance := make(map[point]int)

	// 'S' = 83
	// 'E' = 69
	// 'a' to 'z' = 97 to 122

	for i := 0; scanner.Scan(); i++ {
		row := make([]rune, 0)
		for j, c := range scanner.Text() {
			row = append(row, c)
			if c == 'S' {
				row[j] = 'a'
			} else if c == 'E' {
				start = point{i, j}
				row[j] = 'z'
			}
			visited[point{i, j}] = false
			distance[point{i, j}] = 0
		}
		heightmap = append(heightmap, row)
	}

	// BFS
	queue := []point{start}
	visited[start] = true

	for len(queue) > 0 {
		u := queue[0]
		// remove first element from queue
		queue = queue[1:]
		// found end for first time (shortest path)
		if heightmap[u.i][u.j] == 'a' {
			fmt.Printf("Fewest steps from any 'a' position: %v\n", distance[u])
			break
		}
		// add neighbors to queue
		for _, item := range neighborsInverted(u) {
			if !visited[item] {
				visited[item] = true
				distance[item] = distance[u] + 1
				queue = append(queue, item)
			}
		}
	}
}

func neighborsInverted(p point) (neighbors []point) {
	n := len(heightmap)    // number of rows
	m := len(heightmap[0]) // number of columns

	// search neighbor up down left right
	for _, dir := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		ii := p.i + dir[0]
		jj := p.j + dir[1]

		// check rows and colums bounds
		if 0 <= ii && ii < n && 0 <= jj && jj < m {
			// check if neighbor is at most one step lower
			if heightmap[ii][jj] >= heightmap[p.i][p.j]-1 {
				neighbors = append(neighbors, point{ii, jj})
			}
		}
	}
	return
}
