package day11

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
	operation func(int)int
	testAndThrow func(int)int
}

var monkeys = make([]monkey, 0)

func Part1() {
	input, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	
	for i := 0; scanner.Scan(); i++ {
		scanner.Scan() // skip monkey line
		var newMonkey monkey

		// items line
		for _, item := range strings.Split(scanner.Text()[len("  Starting items: "):], ", ") {
			worryLevel, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			newMonkey.items = append(newMonkey.items, worryLevel)
		}

		scanner.Scan()

		// operation line
		var op rune
		var val int	
		if scanner.Text() == "  Operation: new = old * old"{
			newMonkey.operation = func(val int) int {return (val * val) / 3}
		} else {
			fmt.Sscanf(scanner.Text(), "  Operation: new = old %c %d", &op, &val)
			newMonkey.operation = buildOperation(op, val)
		}

		scanner.Scan()

		// test line
		var divBy int
		var iftrue int
		var iffalse int
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &divBy)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &iftrue)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &iffalse)
		newMonkey.testAndThrow = buildTestAndThrow(divBy, iftrue, iffalse)

		monkeys = append(monkeys, newMonkey)
		scanner.Scan()
	}

	var nInspects = make([]int, len(monkeys))

	for round := 0; round < 20; round++ { // 20 rounds
		for k, currentMonkey := range monkeys { // every monkey
			for _, item := range currentMonkey.items { // every item of monkey
				newWorryLvl := currentMonkey.operation(item)
				monkeys[currentMonkey.testAndThrow(newWorryLvl)].items = append(monkeys[currentMonkey.testAndThrow(newWorryLvl)].items, newWorryLvl)
			}
			nInspects[k] += len(monkeys[k].items)
			// all items thrown, clear item list
			monkeys[k].items = []int{}
		}
	}

	var moreInspects1, moreInspects2 = math.MinInt, math.MinInt
	for _, n := range nInspects {
		if n > moreInspects1 {
			moreInspects2 = moreInspects1
			moreInspects1 = n
		} else if n > moreInspects2 {
			moreInspects2 = n
		}
	}
	fmt.Printf("The level of monkey business after 20 rounds is: %v\n", moreInspects1 * moreInspects2)
}

func buildOperation(op rune, val int) func(int)int {
	if op == '+' {
		return func(i int) int {
			return (i + val) / 3
		}
	} else {
		return func(i int) int {
			return (i * val) / 3
		}
	}
}

func buildTestAndThrow(divBy, ift, iff int) func(int)int {
	return func(i int) int {
		if i % divBy == 0 {
			return ift
		} else {
			return iff
		}
	}
}