package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type cpu struct {
	register  int
	cycle     int
	signalLog []int
}

func (c *cpu) instruction(op string, arg []string) {
	addx := func(arg []string) {
		x, err := strconv.Atoi(arg[0])
		if err != nil {
			log.Fatal("Got invalid arg to addx", err)
		}
		c.signalLog = append(c.signalLog, (c.cycle+1)*c.register, (c.cycle+2)*c.register)
		c.register += x
		c.cycle += 2
	}
	noop := func(arg []string) {
		c.signalLog = append(c.signalLog, (c.cycle+1)*c.register)
		c.cycle += 1
	}

	switch op {
	case "addx":
		addx(arg)
	case "noop":
		noop(arg)
	default:
		log.Fatal("Got unexpected input")
	}
}

func NewCpu() *cpu {
	return &cpu{
		register:  1,
		cycle:     0,
		signalLog: []int(nil),
	}
}

func main() {
	f, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal("could open input file", err)
	}

	// Part 1: Keep a logahead of signals
	cpu := NewCpu()
	s := bufio.NewScanner(f)
	for s.Scan() {
		r := regexp.MustCompile(`(?P<instruction>\w+) ?(?P<arg>[-0-9 ]+)?`)
		m := r.FindStringSubmatch(s.Text())
		switch words := len(m[1:]); words {
		case 0:
			log.Fatal("Got unexpected input")
		case 1:
			cpu.instruction(m[1], []string(nil))
		default:
			cpu.instruction(m[1], m[2:])
		}
	}

	runningSum := 0
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		signal := cpu.signalLog[c-1]
		fmt.Println(fmt.Sprintf("cycle %d:", c), signal)
		runningSum += signal
	}
	fmt.Println("Part 1:", runningSum)

	// Part 2: Scan the signal log
	spri, te := 0, 2
	for cycle, signal := range cpu.signalLog {
		spri, te = (signal/(cycle+1))-1, (signal/(cycle+1))+1
		c := cycle % 40
		if spri <= c && c <= te {
			cpu.signalLog[cycle] = 1
		} else {
			cpu.signalLog[cycle] = 0
		}
	}
	for i := 0; i < len(cpu.signalLog); i += 40 {
		for _, visi := range cpu.signalLog[i : i+40] {
			if visi == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
