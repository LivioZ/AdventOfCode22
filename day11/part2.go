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

func Part2() {
	input, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var monkeys = make([]monkey, 0)
	var bigMod int = 1

	for scanner.Scan() {
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
		if scanner.Text() == "  Operation: new = old * old" {
			newMonkey.operation = func(val int) int { return (val * val) }
		} else {
			fmt.Sscanf(scanner.Text(), "  Operation: new = old %c %d", &op, &val)
			newMonkey.operation = buildOperation2(op, val)
		}

		scanner.Scan()

		// test line
		var divBy int
		var iftrue int
		var iffalse int
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &divBy)
		bigMod *= divBy
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &iftrue)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &iffalse)
		newMonkey.testAndThrow = buildTestAndThrow(divBy, iftrue, iffalse)

		monkeys = append(monkeys, newMonkey)
		scanner.Scan()
	}

	var nInspects = make([]int, len(monkeys))

	for round := 0; round < 10000; round++ { // 10000 rounds
		for k, currentMonkey := range monkeys { // every monkey
			for _, item := range currentMonkey.items { // every item of monkey
				newWorryLvl := currentMonkey.operation(item) % bigMod
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
	fmt.Printf("The level of monkey business after 10000 rounds is: %v\n", moreInspects1*moreInspects2)
}

func buildOperation2(op rune, val int) func(int) int {
	if op == '+' {
		return func(i int) int {
			return (i + val)
		}
	} else {
		return func(i int) int {
			return (i * val)
		}
	}
}
