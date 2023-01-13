package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type empty struct{}

type Coord struct {
	x, y int
}

type Sensor struct {
	loc           Coord
	closestBeacon Coord
}

type visiter []byte

func (v visiter) getNeighbors(c Coord, w, h int, shiftx, shifty int) []byte {
	up := (c.x + shiftx + (c.y+shifty-1)*w) % len(v)
	down := (c.x + shiftx + (c.y+shifty+1)*w) % len(v)
	left := (c.x + shiftx - 1 + (c.y+shifty)*w) % len(v)
	right := (c.x + shiftx + 1 + (c.y+shifty)*w) % len(v)

	b := make([]byte, 4)
	for i, n := range []int{up, down, left, right} {
		if n < 0 {
			b[i] = 2
		} else {
			b[i] = v[n]
		}
	}
	return b
}

func (v visiter) visit(c Coord, w, h int, shiftx, shifty int) {
	v[(c.x+shiftx+(c.y+shifty)*w)%len(v)] |= 1
}

func (v visiter) beacon(c Coord, w, h int, shiftx, shifty int) {
	v[(c.x+shiftx+(c.y+shifty)*w)%len(v)] |= 2
}

func manhattanDistance(c1, c2 Coord) int {
	return int(math.Abs(float64(c2.x-c1.x)) + math.Abs(float64(c2.y-c1.y)))
}

type rowVisiter []map[Coord]bool

func (v rowVisiter) visitRow(c Coord, beacon bool, shifty int) {
	if _, ok := v[c.y+shifty][c]; !ok {
		v[c.y+shifty][c] = beacon
	}
}

func main() {
	fi, err := os.Open("day15.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs := bufio.NewScanner(fi)

	inputPattern := regexp.MustCompile(
		`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`,
	)

	sensors := []Sensor(nil)
	allcoords := []Coord(nil)
	for bs.Scan() {
		line := bs.Text()
		coordstrings := inputPattern.FindAllStringSubmatch(line, 3)[0][1:]
		coords := make([]int, len(coordstrings))
		for i, c := range coordstrings {
			n, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			coords[i] = n
		}
		c := []Coord{{coords[0], coords[1]}, {coords[2], coords[3]}}
		allcoords = append(allcoords, c...)
		sensors = append(sensors, Sensor{c[0], c[1]})
	}

	xmax, xmin := 0, 0
	ymax, ymin := 0, 0
	for _, c := range allcoords {
		if c.x > xmax {
			xmax = c.x
		}
		if c.x < xmin {
			xmin = c.x
		}
		if c.y > ymax {
			ymax = c.y
		}
		if c.y < ymin {
			ymin = c.y
		}
	}
	padding := 0
	xmax, xmin = xmax+padding, xmin-padding
	ymax, ymin = ymax+padding, ymin-padding

	width, height := xmax-xmin, ymax-ymin
	shiftx, shifty := -xmin, -ymin
	visit := make(rowVisiter, height+1)
	for i := range visit {
		visit[i] = make(map[Coord]bool)
	}
	fmt.Println(len(visit), width, height, shiftx, shifty)

	for _, s := range sensors {
		sensorVisit := make(map[Coord]empty)
		toVisit := []Coord{s.loc}
		dist := manhattanDistance(s.loc, s.closestBeacon)
		visit.visitRow(s.closestBeacon, true, shifty)
		fmt.Println("new sensor", s.loc, s.closestBeacon, dist)
		for len(toVisit) > 0 {
			visiting := toVisit[0]
			toVisit = toVisit[1:]

			for _, c := range []Coord{
				{visiting.x, visiting.y + 1},
				{visiting.x, visiting.y - 1},
				{visiting.x - 1, visiting.y},
				{visiting.x + 1, visiting.y},
			} {
				if _, ok := sensorVisit[c]; !ok && manhattanDistance(s.loc, c) <= dist {
					if c.y+shifty >= len(visit) {
						visit = append(visit, make(map[Coord]bool))
					}
					if c.y+shifty < len(visit) {
						visit = append([]map[Coord]bool{make(map[Coord]bool)}, visit...)
						shifty++
					}
					visit.visitRow(c, false, shifty)
					sensorVisit[c] = empty{}
					toVisit = append(toVisit, c)
				}
			}
		}
	}
	count := 0
	fmt.Println(visit)
	for _, beacon := range visit[10+shifty] {
		if !beacon {
			count++
		}
	}
	fmt.Println(count)
}
