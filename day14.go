package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type cave struct {
	scan []byte
}

func main() {
	f, err := os.Open("day14.txt")
	if err != nil {
		log.Fatal("Could not read input file")
	}

	s := bufio.NewScanner(f)
	coords := []coord{}
	xmin, xmax := 1<<31, 0
	ymin, ymax := 1<<31, 0
	for s.Scan() {
		line := s.Text()
		for _, coordString := range strings.Split(line, "->") {
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
			if y < ymin {
				ymin = y
			}
			coords = append(coords, coord{x, y})
		}
	}
	fmt.Println(coords, xmin, xmax, ymin, ymax)

	cave := &cave{scan: make([]byte, (xmax-xmin)*(ymax-ymin))}
}
