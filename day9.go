package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type head struct {
	x, y int
	knot *knot
}

func (h *head) move(direction byte, count int) *knot {
	switch direction {
	case 'U':
		h.y += count
	case 'R':
		h.x += count
	case 'D':
		h.y -= count
	case 'L':
		h.x -= count
	}
	return h.knot
}

type knot struct {
	x, y    int
	knot    *knot
	visited map[struct{ x, y int }]bool
}

func (t *knot) simulate(hx, hy int) (*knot, int, int) {
	xd := (hx - t.x)
	yd := (hy - t.y)

	s := func() int {
		return yd*yd + xd*xd
	}

	if s() > 2 {
		t.move(xd, yd)
		c := struct{ x, y int }{t.x, t.y}
		if t.visited != nil {
			_, ok := t.visited[c]
			if !ok {
				t.visited[c] = true
			}
		}
	}
	return t.knot, t.x, t.y
}

func (t *knot) move(xd, yd int) {
	if xd*xd > 0 {
		if xd > 0 {
			t.x += 1
		} else {
			t.x -= 1
		}
	}
	if yd*yd > 0 {
		if yd > 0 {
			t.y += 1
		} else {
			t.y -= 1
		}
	}
}

func NewRope(length int) (*head, *knot) {
	k := &knot{0, 0, nil, make(map[struct{ x, y int }]bool)}
	t := k
	for i := 0; i < length-2; i++ {
		k = &knot{0, 0, k, nil}
	}
	return &head{
		0, 0, k,
	}, t
}

func main() {
	f, err := os.Open("day9.txt")
	if err != nil {
		fmt.Println("Could not find input file")
	}

	h, t := NewRope(10)
	s := bufio.NewScanner(f)
	for s.Scan() {
		r := regexp.MustCompile(`(\w+) (\d+)`)
		m := r.FindSubmatch(s.Bytes())

		direction := m[1][0]
		count, err := strconv.Atoi(string(m[2]))
		if err != nil {
			log.Fatal("Invalid number string")
		}

		for i := 0; i < count; i++ {
			k := h.move(direction, 1)
			x, y := h.x, h.y
			for k != nil {
				k, x, y = k.simulate(x, y)
			}
		}
	}
	fmt.Println(len(t.visited) + 1)
}
