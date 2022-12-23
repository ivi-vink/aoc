package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	rock  byte = '#'
	air   byte = '.'
	abyss byte = '~'
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type coord struct {
	x, y int
}

type cave struct {
	scan                []byte
	width, height, xmin int
}

func (c *cave) stringify() string {
	cave := ""
	for i := 0; i < len(c.scan); i += c.height {
		cave += string(c.scan[i:i+c.height]) + "\n"
	}
	return cave
}

func (c *cave) setCoord(co coord, value byte) {
	c.scan[(co.x-c.xmin)*c.height+co.y] = value
}

func (c *cave) getCoord(x, y int) byte {
	nx, ny := (x-c.xmin)*c.height, y
	if (x-c.xmin) >= c.width || y >= c.height {
		return abyss
	}
	if nx+ny < 0 {
		return abyss
	}
	return c.scan[nx+ny]
}

func (c *cave) drawRocks(c1, c2 coord) {
	draw := func(co coord, x1, x2 int, up func(i int) coord) {
		dx := x2 - x1
		di := 1
		if dx != 0 {
			if dx < 0 {
				dx = -dx
				di = -di
			}
			for i := 0; abs(i) < dx+1; i += di {
				c.setCoord(up(i), rock)
			}
		}
	}
	draw(c1, c1.x, c2.x, func(i int) coord { return coord{c1.x + i, c1.y} })
	draw(c1, c1.y, c2.y, func(i int) coord { return coord{c1.x, c1.y + i} })
}

func (c *cave) rockFormations(rockFormations [][]coord) {
	for _, r := range rockFormations {
		for j, co := range r {
			if j > 0 {
				c.drawRocks(r[j-1], co)
			}
		}
	}
}

func (c *cave) newFloor() []byte {
	floor := make([]byte, c.height)
	for i := range floor {
		floor[i] = air
	}
	floor[len(floor)-1] = rock
	return floor
}

type sand byte

func (s sand) fall(cave *cave, floor bool) bool {
	co := &coord{500, 0}
	left, beneath, right := s.options(cave, co.x, co.y)
	if left != air && beneath != air && right != air {
		return false
	}
	for {
		left, beneath, right := s.options(cave, co.x, co.y)
		if beneath == air {
			co.y++
			continue
		}
		if left == air {
			co.x--
			co.y++
			continue
		}
		if right == air {
			co.x++
			co.y++
			continue
		}
		if left == abyss || beneath == abyss || right == abyss {
			if floor {
				cave.scan = append(cave.newFloor(), cave.scan...)
				cave.xmin = cave.xmin - 1
				cave.width = cave.width + 1

				cave.scan = append(cave.scan, cave.newFloor()...)
				cave.width = cave.width + 1
				continue
			}
			return false
		}

		if beneath == rock || beneath == byte(s) {
			cave.setCoord(*co, byte(s))
			return true
		}
	}
	fmt.Println("Something went wrong")
	return false
}

func (s sand) options(cave *cave, x, y int) (byte, byte, byte) {
	return cave.getCoord(x-1, y+1),
		cave.getCoord(x, y+1),
		cave.getCoord(x+1, y+1)
}

func main() {
	f, err := os.Open("day14.txt")
	if err != nil {
		log.Fatal("Could not read input file")
	}

	s := bufio.NewScanner(f)
	rockFormations := [][]coord{}
	xmin, xmax, ymax := 1<<31, 0, 0
	for s.Scan() {
		line := s.Text()
		coordStrings := strings.Split(line, "->")
		coords := make([]coord, len(coordStrings))
		for i, coordString := range coordStrings {
			nums := strings.Split(strings.Trim(coordString, " "), ",")
			x, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			if x > xmax {
				xmax = x
			}
			if x < xmin {
				xmin = x
			}
			if y > ymax {
				ymax = y
			}
			coords[i] = coord{x, y}
		}
		rockFormations = append(rockFormations, coords)
	}

	width, height := (xmax-xmin)+1, (ymax-0)+1
	scan := make([]byte, width*height)
	for i := range scan {
		scan[i] = air
	}
	c := &cave{scan, width, height, xmin}
	c.rockFormations(rockFormations)

	sandCounter := 0
	var snd sand = 'o'
	for snd.fall(c, false) {
		sandCounter++
	}
	fmt.Println("Part 1: ", sandCounter)

	correctedHeight := height + 2
	scan = make([]byte, width*correctedHeight)
	for i := range scan {
		scan[i] = air
	}
	c = &cave{scan, width, correctedHeight, xmin}
	floor := []coord{
		{xmin, correctedHeight - 1},
		{xmin + width - 1, correctedHeight - 1},
	}
	rockFormations = append(rockFormations, floor)
	c.rockFormations(rockFormations)

	sandCounter = 1
	for snd.fall(c, true) {
		sandCounter++
	}
	fmt.Println("Part 2: ", sandCounter)
}
