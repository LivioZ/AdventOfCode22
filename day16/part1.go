package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type valve struct {
	flowRate  int
	tunnelsTo []string
}

func Part1() {
	in, err := os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	sc := bufio.NewScanner(in)

	valves := make(map[string]valve)

	for sc.Scan() {
		line := strings.Split(sc.Text(), "; tunnels lead to valves ") // ["Valve AA has flow rate=0", "DD, II, BB"]
		if len(line) == 1 {
			line = strings.Split(sc.Text(), "; tunnel leads to valve ") // ["Valve HH has flow rate=22", "GG"]
		}

		var n string
		var fr int
		var tt []string
		fmt.Sscanf(line[0], "Valve %s has flow rate=%d; tunnels lead to", &n, &fr)
		tt = strings.Split(line[1], ", ")

		valves[n] = valve{flowRate: fr, tunnelsTo: tt}
	}
}
