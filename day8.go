// I'll come back later to this day
//
// The second part is bad with time O(n*(W+H)) -> could be better for sure like O(4n)
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type quadcopter struct {
	scans []*scan
}

// Scans all rows and columns, probably could be a bit better
func (q *quadcopter) scanMap(hmap [][]byte, xlen, ylen int) {
	for i, row := range hmap {
		q.scans[i] = NewScan(row)
	}
	for y := 0; y < ylen; y++ {
		col := make([]byte, xlen)
		for x := 0; x < xlen; x++ {
			col[x] = hmap[x][y]
		}
		q.scans[xlen+y] = NewScan(col)
	}
}

func (q *quadcopter) countVisibleTrees(hmap [][]byte, xlen, ylen int) int {
	// circumference
	visible := 2 * (xlen - 1 + ylen - 1)
	// inner rectangle
	for x := 1; x < xlen-1; x++ {
		for y := 1; y < ylen-1; y++ {
			tree := hmap[x][y]
			xv := q.scans[x].progressScan(tree)
			yv := q.scans[xlen+y].progressScan(tree)
			if xv {
				visible += 1
				continue
			}
			if yv {
				visible += 1
				continue
			}
		}
	}
	return visible
}

type scan struct {
	ahead   []byte
	visited byte
}

func NewScan(trees []byte) *scan {
	lastTree := len(trees) - 1
	s := scan{
		visited: trees[lastTree],
		ahead:   []byte{trees[lastTree]},
	}
	// stop before tree[0], it's not in the scan technically
	for j := lastTree - 1; j > 0; j-- {
		if trees[j] >= s.visited {
			s.visited = trees[j]
			s.ahead = append(s.ahead, trees[j])
		}
	}
	s.visited = trees[0]
	return &s
}

func (s *scan) progressScan(tree byte) bool {
	// tree was greatest in scan ahead
	if tree == s.ahead[len(s.ahead)-1] {
		// keep last tree
		if len(s.ahead) > 1 {
			s.ahead = s.ahead[:len(s.ahead)-1]
		}
	}

	if tree > s.visited {
		s.visited = tree
		return true
	}

	if tree > s.ahead[len(s.ahead)-1] {
		return true
	}
	return false
}

// Could be less repetitive with closures? But go kind of sucks with iterators
func scenicScore(tree byte, hmap [][]byte, x, xlen, y, ylen int) int {
	down := 0
	for i := x + 1; i < xlen; i++ {
		other := hmap[i][y]
		down = i - x
		if other >= tree {
			break
		}
	}
	up := 0
	for i := x - 1; i > -1; i-- {
		other := hmap[i][y]
		up = x - i
		if other >= tree {
			break
		}
	}
	left := 0
	for i := y + 1; i < ylen; i++ {
		other := hmap[x][i]
		left = i - y
		if other >= tree {
			break
		}
	}
	right := 0
	for i := y - 1; i > -1; i-- {
		other := hmap[x][i]
		right = y - i
		if other >= tree {
			break
		}
	}
	return up * down * left * right
}

func main() {
	// part 1
	f, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}

	xlen, ylen := 0, 0
	hmap := [][]byte{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		b := s.Bytes()
		if ylen == 0 {
			ylen = len(b)
		}
		bcopy := append([]byte(nil), b...)
		hmap = append(hmap, bcopy)
		xlen += 1
	}
	q := &quadcopter{
		scans: make([]*scan, xlen+ylen),
	}
	q.scanMap(hmap, xlen, ylen)
	vis := q.countVisibleTrees(hmap, xlen, ylen)
	fmt.Println("Part 1:", vis)

	// part 2
	runningScore := 0
	for x := 1; x < xlen-1; x++ {
		for y := 1; y < ylen-1; y++ {
			tree := hmap[x][y]
			if s := scenicScore(tree, hmap, x, xlen, y, ylen); s > runningScore {
				runningScore = s
			}
		}
	}
	fmt.Println("Part 2:", runningScore)
}
