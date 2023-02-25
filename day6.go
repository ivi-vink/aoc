package main

import (
	"fmt"
	"log"
	"os"
)

func inside(c byte, b []byte) int {
	for i, a := range b {
		if c == a {
			return i
		}
	}
	return -1
}

func packetMarker(buf []byte, l int) int {
	i, j := 0, 1
	for (j - i) < l {
		n := inside(buf[j], buf[i:j])
		if n != -1 {
			i += n + 1
		} else {
			j++
		}
	}
	return j
}

func main() {
	f, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal("Could not open input file", err)
	}
	buf := make([]byte, 32*1024)
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal("Got error reading into buffer", err)
	}
	// Part 1
	r := packetMarker(buf, 4)
	fmt.Println("Part 1", r, string(buf[r-4:r]))

	// Part 2
	r = packetMarker(buf, 14)
	fmt.Println("Part 2", r, string(buf[r-14:r]))
}
