package day07

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const totSpace = 70000000
const neededFreeSpace = 30000000

var lowerSizeDirToDelete int = math.MaxInt

func Part2() {
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

	unusedSpace := totSpace - root.size
	spaceToDelete := neededFreeSpace - unusedSpace
	fmt.Println(spaceToDelete)
	findLowerSizeDirToDelete(root, spaceToDelete)
	fmt.Printf("Smallest directory to delete: %v\n", lowerSizeDirToDelete)
}

func findLowerSizeDirToDelete(node *node, spaceToDelete int) {
	if node.size >= spaceToDelete && !node.isFile && node.size < lowerSizeDirToDelete {
		lowerSizeDirToDelete = node.size
	}
	for _, child := range node.children {
		findLowerSizeDirToDelete(child, spaceToDelete)
	}
}
