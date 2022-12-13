package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type (
	t     interface{}
	empty struct{}
)

type packet struct {
	data []t
}

func intOrList(v t) int {
	switch v.(type) {
	case int:
		return 0
	case []t:
		return 1
	default:
		return -1
	}
}

type continuation struct {
	i    int
	ok   bool
	a, b []t
}

func comparePackets(p1, p2 *packet, comp func(a, b int) bool) (int, bool) {
	var cp func(i int, ok bool, a, b []t) (int, bool)

	var ctn []continuation
	cp = func(i int, ok bool, a, b []t) (int, bool) {
		if !ok {
			return i, false
		}
		if al, bl := i >= len(a), i >= len(b); al || bl {
			if bl && !al {
				return i, false
			}
			return i, ok
		}
		at, bt := intOrList(a), intOrList(b)
		if at == 1 && bt == 1 {
		} else if at == 1 && bt == 0 {
		} else if at == 0 && bt == 1 {
		} else if at == 0 && bt == 0 {
		} else {
		}

		return i, true
	}
	ctn = append(ctn, continuation{0, true, p1.data, p2.data})
	for len(ctn) > 0 {
	}
	return cp(0, true, p1.data, p2.data)
}

func parsePacket(b []byte) *packet {
	var parseList func(list []t, b []byte) []t

	parseList = func(list []t, b []byte) []t {
		for i := 0; i < len(b); i++ {
			switch {
			case b[0] == '[' && b[1] == ']':
				list = append(list, []t{})
				return list
			case b[i] == '[':
				j := 1
				subchildren := 1
				for subchildren > 0 {
					if b[i+j] == '[' {
						subchildren += 1
					}
					if b[i+j] == ']' {
						subchildren = subchildren - 1
					}
					j++
				}
				list = append(list, parseList([]t{}, b[i+1:i+j-1]))
				i = i + j
			default:
				j := 0
				for ; i+j < len(b); j++ {
					stop := false
					switch b[i+j] {
					case ']':
						stop = true
					case ',':
						stop = true
					}
					if stop {
						break
					}
				}
				s := string(b[i : i+j])
				if len(s) == 0 {
					s = string(b[len(b)-1])
				}
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal("This is no a number!", i, j, s)
				}
				list = append(list, num)
				i = i + j
			}
		}
		return list
	}

	return &packet{data: parseList([]t{}, b[1:len(b)-1])}
}

func main() {
	f, err := os.Open("day13.txt")
	if err != nil {
		log.Fatal("Could not read input file")
	}

	s := bufio.NewScanner(f)
	packets := []*packet{}
	for s.Scan() {
		if len(s.Bytes()) > 0 {
			packets = append(packets, parsePacket(s.Bytes()))
		}
	}
	for i := 0; i < len(packets); i += 2 {
		fmt.Println(packets[i])
		fmt.Println(comparePackets(packets[i], packets[i+1], func(a, b int) bool { return a < b }))
	}
}
