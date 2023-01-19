package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type Valve struct {
	name string
}

func readValves(f io.Reader) map[string]Valve {
	pattern := regexp.MustCompile(
		`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z, ]+)`,
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		valveInfo := pattern.FindAllStringSubmatch(line, 3)[0]
		fmt.Println(valveInfo)
	}
	return nil
}

/*
AA,0 ===================
||.............\\      \\
BB,13===CC,2===DD,20   II,0
...............\\       \\
...............EE,20    JJ,21
...............\\
...............FF,0
...............\\
...............GG,0
...............\\
...............HH,22
Probably can do some apriori thing here
*/
func main() {
	fh, err := os.Open("day16.txt")
	if err != nil {
		log.Fatal("Input file not found")
	}
	valves := readValves(fh)
	fmt.Println(valves)
}
