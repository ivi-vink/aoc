package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type filesystem struct {
	root *dir
}

type dir struct {
	parent  *dir
	subdirs map[string]*dir
	files   map[string]*file
}

func newDir(parent *dir) *dir {
	return &dir{
		parent:  parent,
		subdirs: make(map[string]*dir),
		files:   make(map[string]*file),
	}
}

type file struct {
	name string
	size int
}

func changeDir(fs *filesystem, wd *dir, arg string) *dir {
	switch arg {
	case "..":
		return wd.parent
	case "/":
		return fs.root
	default:
		d, ok := wd.subdirs[arg]
		if !ok {
			log.Fatal("Tried going into non existing directory")
		}
		return d
	}
}

func lsOut(wd *dir, words []string) {
	if words[0] == "dir" {
		wd.subdirs[words[1]] = newDir(wd)
		return
	}
	if num, err := strconv.Atoi(words[0]); err == nil {
		wd.files[words[1]] = &file{
			name: words[1],
			size: num,
		}
	}
}

// Assume only one command arg
func commands(fs *filesystem, wd *dir, s *bufio.Scanner) {
	lsout := false
	for s.Scan() {
		words := strings.Split(s.Text(), " ")
		if words[0] == "$" {
			lsout = false

			switch words[1] {
			case "ls":
				lsout = true
				continue
			case "cd":
				wd = changeDir(fs, wd, words[2])
			}
		}

		if lsout {
			lsOut(wd, words)
		}
	}
}

func byMaxSize(fs *filesystem, maxSize int) (int, []int) {
	return byPred(fs, maxSize, func(a int) bool {
		return a <= maxSize
	})
}

func byMinSize(fs *filesystem, minSize int) (int, []int) {
	return byPred(fs, minSize, func(a int) bool {
		return a >= minSize
	})
}

func byPred(fs *filesystem, by int, pred func(a int) bool) (int, []int) {
	dirSizes := []int{}
	wd := fs.root
	var bms func(wd *dir) int
	bms = func(wd *dir) int {
		size := 0
		for _, f := range wd.files {
			size += f.size
		}
		for _, d := range wd.subdirs {
			size += bms(d)
		}
		if pred(size) {
			dirSizes = append(dirSizes, size)
		}
		return size
	}

	return bms(wd), dirSizes
}

func main() {
	f, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal("Could not open input file", err)
	}

	// Part 1
	s := bufio.NewScanner(f)
	fs := &filesystem{
		root: newDir(nil),
	}
	var wd *dir = nil
	// build the fs tree
	commands(fs, wd, s)
	// recurse the three
	total, dirSizes := byMaxSize(fs, 100_000)
	sum := 0
	for _, d := range dirSizes {
		sum += d
	}
	fmt.Println("Part 1:", total, sum)

	// Part 2
	total, _ = byMaxSize(fs, 0)
	free := (70_000_000 - total)
	need := (30_000_000 - free)
	total, dirSizes = byMinSize(fs, need)
	smallest := dirSizes[0]
	for _, d := range dirSizes {
		if d < smallest {
			smallest = d
		}
	}
	fmt.Println("Part 2", need, smallest)
}
