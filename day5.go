package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type crates struct {
	stacks [][]byte
}

func (c *crates) topCrates() string {
	b := make([]byte, len(c.stacks))
	for i, stack := range c.stacks {
		if len(stack) > 0 {
			b[i] = stack[len(stack)-1]
		}
	}
	return string(b)
}

func (c *crates) move(m move) {
	// copy slices
	from := c.stacks[m.from-1]
	to := c.stacks[m.to-1]

	for i := 0; i < m.amount; i++ {
		pop, popped := from[len(from)-1], from[:len(from)-1]
		from = popped
		to = append(to, pop)
	}

	// set slices
	c.stacks[m.to-1] = to
	c.stacks[m.from-1] = from
}

func (c *crates) moveMultiple(m move) {
	// copy slices
	from := c.stacks[m.from-1]
	to := c.stacks[m.to-1]

	pop, from := from[len(from)-m.amount:], from[:len(from)-m.amount]

	// push
	c.stacks[m.to-1] = append(to, pop...)
	// set popped
	c.stacks[m.from-1] = from
}

func (c *crates) parseLine(l []byte) {
	// by four
	for i := 0; i < len(l); i += 4 {
		if 'A' <= l[i+1] && l[i+1] <= 'Z' {
			// push front: https://github.com/golang/go/wiki/SliceTricks#push-frontunshift
			c.stacks[i/4] = append([]byte{l[i+1]}, c.stacks[i/4]...)
		}
	}
}

func newCargo(s *bufio.Scanner) *crates {
	// build cargo
	s.Scan()
	line := s.Bytes()
	//[x].
	//1234
	cargo := &crates{stacks: make([][]byte, len(line)/4+1)}
	cargo.parseLine(line)

	for s.Scan() {
		line = s.Bytes()
		if len(line) == 0 {
			break
		}
		cargo.parseLine(line)
	}
	return cargo
}

type move struct {
	amount int
	from   int
	to     int
}

func parseMove(l string) move {
	regex := regexp.MustCompile(`move (?P<amount>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	matches := regex.FindStringSubmatch(l)
	m := make([]int, 3)
	for i, match := range matches[1:] {
		matchInt, err := strconv.Atoi(match)
		m[i] = matchInt
		if err != nil {
			log.Fatal("Got unexpected input in move")
		}
	}
	return move{
		amount: m[0],
		from:   m[1],
		to:     m[2],
	}
}

func main() {
	f, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal("Could not open input file")
	}

	// part 1
	scanner := bufio.NewScanner(f)
	cargo := newCargo(scanner)
	for scanner.Scan() {
		l := scanner.Text()
		m := parseMove(l)
		cargo.move(m)
	}
	fmt.Println("Part 1:", cargo.topCrates())

	f.Seek(0, 0)

	// part 2
	scanner = bufio.NewScanner(f)
	cargo = newCargo(scanner)
	for scanner.Scan() {
		l := scanner.Text()
		m := parseMove(l)
		cargo.moveMultiple(m)
	}
	fmt.Println("Part 2:", cargo.topCrates())
}
