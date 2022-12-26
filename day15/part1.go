package day15

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type point struct {
	x, y int
}

type interval struct {
	x1, x2 int
}

func Part1() {
	in, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	sc := bufio.NewScanner(in)

	intervals := make([]interval, 0)
	beacons := make(map[point]bool)
	const Y = 2000000

	for sc.Scan() {
		var s, b point
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)
		beacons[b] = true

		distance := int(math.Abs(float64(s.x-b.x)) + math.Abs(float64(s.y-b.y)))

		if distance > int(math.Abs(float64(Y-s.y))) {
			var iv interval
			xdist := (distance - int(math.Abs(float64(Y-s.y))))

			iv.x1 = s.x - xdist
			iv.x2 = s.x + xdist
			intervals = append(intervals, iv)
		}
	}

	minx := math.MaxInt
	maxx := 0

	for _, iv := range intervals {
		if iv.x1 < minx {
			minx = iv.x1
		}
		if iv.x2 > maxx {
			maxx = iv.x2
		}
	}

	var res int
	// bad: O(n^2)
	for x := minx; x <= maxx; x++ {
		if beacons[point{x, Y}] {
			continue
		}
		for _, iv := range intervals {
			if iv.x1 <= x && x <= iv.x2 {
				res++
				break
			}
		}
	}

	fmt.Printf("Positions in row %v that cannot contain a beacon: %v\n", Y, res)
}
