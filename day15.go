// Part 2 could be further optimised by taking into account which sensors touch at manhattan distance + 1
// But the main trick was recognising traversing the circumference as a way to reduce the search space.
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

type Coord struct {
	x, y int
}

type Sensor struct {
	loc           Coord
	closestBeacon Coord
}

type visiter []byte

func manhattanDistance(c1, c2 Coord) int {
	return int(math.Abs(float64(c2.x-c1.x)) + math.Abs(float64(c2.y-c1.y)))
}

func points(s Sensor) map[Coord]struct{} {
	d := manhattanDistance(s.loc, s.closestBeacon)
	points := []Coord{{s.loc.x + d + 1, s.loc.y}}
	circ := make(map[Coord]struct{})
	for len(points) > 0 {
		p := points[0]
		points = points[1:]
		for _, a := range []Coord{
			{p.x - 1, p.y + 1},
			{p.x + 1, p.y - 1},
			{p.x - 1, p.y - 1},
			{p.x + 1, p.y + 1},
		} {
			if manhattanDistance(s.loc, a) == d+1 {
				_, ok := circ[a]
				if !ok {
					circ[a] = struct{}{}
					points = append(points, a)
				}
			}
		}
	}
	return circ
}

func findAtRow(distances []int, sensors []Sensor, row, begin, end int) (int, []Coord) {
	covered := 0
	notCovered := []Coord{}
	for i := begin; i < end; i++ {
		coord := Coord{i, row}
		cov := false
		for i, c := range sensors {
			if d := manhattanDistance(c.loc, coord); d <= distances[i] {
				cov = true
				if coord != c.loc && coord != c.closestBeacon {
					covered++
					break
				}
			}
		}
		if !cov {
			notCovered = append(notCovered, coord)
		}
	}
	return covered, notCovered
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
		sensors = append(sensors, Sensor{c[0], c[1]})
	}

	xmax, xmin := 0, 0
	ymax, ymin := 0, 0
	for _, c := range sensors {
		if c.loc.x > xmax {
			xmax = c.loc.x
		}
		if c.loc.x+c.closestBeacon.x > xmax {
			xmax = c.loc.x + c.closestBeacon.x
		}
		if c.loc.x < xmin {
			xmin = c.loc.x
		}
		if c.loc.x-c.closestBeacon.x < xmin {
			xmin = c.loc.x - c.closestBeacon.x
		}
		if c.loc.y > ymax {
			ymax = c.loc.y
		}
		if c.loc.y+c.closestBeacon.y > ymax {
			ymax = c.loc.y + c.closestBeacon.y
		}
		if c.loc.y < ymin {
			ymin = c.loc.y
		}
		if c.loc.y-c.closestBeacon.y < ymin {
			ymin = c.loc.y - c.closestBeacon.y
		}
	}

	distances := []int(nil)
	for _, c := range sensors {
		distances = append(distances, manhattanDistance(c.loc, c.closestBeacon))
	}
	fmt.Println(xmin, xmax, ymin, ymax)

	row := 2_000_000
	count, _ := findAtRow(distances, sensors, row, xmin, xmax)
	fmt.Println(count)

	lower, upper := 0, 4_000_000
	target := Coord{}
	t := false
	for _, s := range sensors {
		p := points(s)
		for c := range p {
			notCovered := true
			if c.x < lower || c.x > upper {
				continue
			}
			if c.y < lower || c.y > upper {
				continue
			}
			for _, other := range sensors {
				if other != s {
					d := manhattanDistance(other.loc, other.closestBeacon)
					if d >= manhattanDistance(other.loc, c) {
						notCovered = false
					}
				}
			}
			if notCovered {
				target = c
				t = true
				break
			}
		}
		if t {
			break
		}
	}
	fmt.Println(target)
}
