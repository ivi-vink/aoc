package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	startSquare = 'S'
	endSquare   = 'E'
)

func byte2height(b byte) int {
	return int(b - 'a')
}

type square struct {
	x, y, i int
	height  int
}

type heightmap struct {
	squares [][]*square

	work      []*square
	pathlog   map[int]*square
	distances map[int]int
}

func (h *heightmap) distance(s *square, o *square) int {
	return 1
}

func (h *heightmap) neighbors(s *square) []*square {
	potential := []*square{}
	if s.x > 0 {
		potential = append(potential, h.squares[s.y][s.x-1])
	}
	if s.x < len(h.squares[0])-1 {
		potential = append(potential, h.squares[s.y][s.x+1])
	}
	if s.y > 0 {
		potential = append(potential, h.squares[s.y-1][s.x])
	}
	if s.y < len(h.squares)-1 {
		potential = append(potential, h.squares[s.y+1][s.x])
	}

	ns := []*square{}
	for _, n := range potential {
		if (n.height - s.height) <= 1 {
			ns = append(ns, n)
		}
	}
	return ns
}

func (hm *heightmap) shortestPath(start *square, pred func(sqr *square) bool) ([]*square, bool) {
	infinity := len(hm.squares) * len(hm.squares[0])
	hm.work = make([]*square, len(hm.squares)*len(hm.squares[0]))
	for i := range hm.squares {
		for j := range hm.squares[i] {
			s := hm.squares[i][j]
			hm.distances[s.i] = infinity
			hm.pathlog[s.i] = nil
			hm.work[j+(i*len(hm.squares[0]))] = s
		}
	}
	hm.distances[start.i] = 0

	var sqr *square
	for len(hm.work) > 0 {
		var i int
		sqr = hm.work[0]
		for j, sqrmin := range hm.work {
			if hm.distances[sqrmin.i] < hm.distances[sqr.i] {
				sqr = sqrmin
				i = j
			}
		}
		if sqr == nil {
			log.Fatal("No min distance square found")
		}
		hm.work[i], hm.work[len(hm.work)-1] = hm.work[len(hm.work)-1], hm.work[i]
		hm.work = hm.work[:len(hm.work)-1]
		if pred(sqr) {
			break
		}
		for _, n := range hm.neighbors(sqr) {
			d := hm.distances[sqr.i] + hm.distance(sqr, n)
			if d < hm.distances[n.i] {
				hm.distances[n.i] = d
				hm.pathlog[n.i] = sqr
			}
		}
	}
	path := []*square{}
	if hm.pathlog[sqr.i] == nil {
		return path, false
	}
	ptr := sqr.i
	for hm.pathlog[ptr] != nil {
		p := hm.pathlog[ptr]
		path = append(path, p)
		ptr = p.i
	}
	return path, true
}

func main() {
	f, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal("Could not read input data", err)
	}
	s := bufio.NewScanner(f)
	y := 0
	d := make(map[int]int)
	hm := &heightmap{
		pathlog:   make(map[int]*square),
		distances: d,
	}
	otherStarts := []*square{}
	var start, end *square
	for s.Scan() {
		squareRow := make([]*square, len(s.Bytes()))
		for x, c := range s.Bytes() {
			if c == 'S' {
				s := &square{x, y, x + (y * len(s.Bytes())), byte2height('a')}
				squareRow[x] = s
				start = s
			} else if c == 'E' {
				s := &square{x, y, x + (y * len(s.Bytes())), byte2height('z')}
				squareRow[x] = s
				end = s
			} else {
				s := &square{x, y, x + (y * len(s.Bytes())), byte2height(c)}
				if c == 'a' {
					otherStarts = append(otherStarts, s)
				}
				squareRow[x] = s
			}
		}
		hm.squares = append(hm.squares, squareRow)
		y++
	}

	p1path, ok := hm.shortestPath(start, func(sqr *square) bool {
		return sqr == end
	})
	if ok {
		fmt.Println("Part 1:", hm.distances[end.i])
	}

	shortest := len(p1path)
	for _, strt := range otherStarts {
		path, ok := hm.shortestPath(strt, func(sqr *square) bool {
			return sqr == end || hm.distances[sqr.i] >= len(p1path)
		})
		if ok {
			if len(path) < shortest {
				shortest = len(path)
			}
		}
	}
	fmt.Println("Part 2:", shortest)
}
