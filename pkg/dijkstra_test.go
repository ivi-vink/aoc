package algos

import (
	"strings"
	"testing"
)

const day12input string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

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

func (s *square) Index() int { return s.i }

type squareMap [][]Node

func (s squareMap) Neighbors(n Node) []Node
func (s squareMap) UpdateOrEnd(n Node) bool

func readInput(input string) (start, end *square, squares [][]Node, flat []Node) {
	y := 0
	for _, s := range strings.Split(input, "\n") {
		squareRow := make([]Node, len(s))
		for x, c := range s {
			if c == 'S' {
				s := &square{x, y, x + (y * len(s)), byte2height('a')}
				squareRow[x] = s
				start = s
			} else if c == 'E' {
				s := &square{x, y, x + (y * len(s)), byte2height('z')}
				squareRow[x] = s
				end = s
			} else {
				s := &square{x, y, x + (y * len(s)), byte2height(byte(c))}
				squareRow[x] = s
			}
		}
		squares = append(squares, squareRow)
		flat = append(flat, squareRow...)
		y++
	}
	return
}

func Test_dijkstra(t *testing.T) {
	start, end, squares, flat := readInput(day12input)
	type args struct {
		searchSpace []Node
		start       Node
		noder       Noder
	}
	tests := []struct {
		name string
		args args
	}{
		{"dummy", args{nil, nil, nil}},
		{"test", args{flat, start}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dijkstra(tt.args.searchSpace, tt.args.start, tt.args.noder)
		})
	}
}
