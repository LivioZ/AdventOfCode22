package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isFile   bool
	children map[string]*node
	parent   *node
}

var result int

func Part1() {
	input, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var currentDirectory *node
	var root *node

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// $ cd dir
		if len(line) == 3 {
			if line[2] == "/" {
				if root == nil {
					root = &node{"/", 0, false, make(map[string]*node), nil}
				}
				currentDirectory = root
			} else if line[2] == ".." {
				currentDirectory = currentDirectory.parent
			} else {
				currentDirectory = currentDirectory.children[line[2]]
			}
		} else if line[1] == "ls" {
			continue
		} else if line[0] == "dir" {
			currentDirectory.children[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
		} else {
			size, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)
			}
			currentDirectory.children[line[1]] = &node{line[1], size, true, nil, currentDirectory}
		}
	}

	calculateTreeSize(root)
	sumSize(root)
	fmt.Printf("Sum of size of directories with a total size of at most 100000: %v\n", result)
}

// visit to calculate all directories' size
func calculateTreeSize(node *node) int {
	if node.isFile {
		return node.size
	} else if len(node.children) == 0 {
		node.size = 0
		return node.size
	} else {
		result := 0
		for _, child := range node.children {
			result += calculateTreeSize(child)
		}
		node.size = result
		return result
	}
}

// visit to sum only directories with size > 100 000
func sumSize(node *node) {
	if !node.isFile && node.size <= 100000 {
		result += node.size
	}
	for _, child := range node.children {
		sumSize(child)
	}
}
