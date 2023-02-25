package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type round struct {
	opponent shape
	player   shape
}

type roundOutcome struct {
	lost int
	draw int
	won  int
}

var outcome roundOutcome = roundOutcome{
	lost: 0,
	draw: 3,
	won:  6,
}

type shape rune

const (
	opponentRock, playerRock shape = iota + 'A', iota + 'X'
	opponentPaper, playerPaper
	opponentScissors, playerScissors
)

func shapePoints(p shape) int {
	return int((p - playerRock + 1) * 1)
}

func loseDrawWin(o shape) (shape, shape, shape) {
	switch o {
	case opponentRock:
		return playerScissors, playerRock, playerPaper
	case opponentPaper:
		return playerRock, playerPaper, playerScissors
	case opponentScissors:
		return playerPaper, playerScissors, playerRock
	default:
		return 0, 0, 0
	}
}

func outcomePoints(r round) int {
	player := int(r.player - playerPaper)
	opponent := int(r.opponent - opponentPaper)
	if player == opponent {
		return 3
	}
	if player+opponent == 0 && player < opponent {
		return 6
	} else if player+opponent == 0 && player > opponent {
		return 0
	}
	if player < opponent {
		return 0
	} else {
		return 6
	}
}

func outcomePointsCheating(p shape) int {
	switch p {
	case playerRock:
		return 0
	case playerPaper:
		return 3
	case playerScissors:
		return 6
	default:
		return 0
	}
}

func shapePointsCheating(r round) int {
	lose, draw, win := loseDrawWin(r.opponent)
	switch outcomePointsCheating(r.player) {
	case 0:
		return shapePoints(lose)
	case 3:
		return shapePoints(draw)
	case 6:
		return shapePoints(win)
	default:
		return 0
	}
}

func countRounds(f io.Reader, buf []byte) (int, error) {
	count := 0
	for {
		i, err := f.Read(buf)
		count += bytes.Count(buf[:i], []byte{'\n'})
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			log.Fatal("Error reading input file")
			return count, err
		}
	}
}

// A very verbose way to read lines into rounds without the bufio package
func getRounds(f io.Reader, buf []byte, rounds []round) []round {
	r := 0
	for {
		n, err := f.Read(buf)
		for i := 0; i < n; i += 4 {
			o := shape(buf[i])
			p := shape(buf[i+2])
			if playerRock <= p && p <= playerScissors {
				rounds[r].player = p
			}
			if opponentRock <= o && o <= opponentScissors {
				rounds[r].opponent = o
			}
			if rounds[r].player == 0 || rounds[r].opponent == 0 {
				rounds[r] = rounds[len(rounds)-1]
				rounds = rounds[:len(rounds)-1]
			}
			r++
		}
		if err == io.EOF {
			return rounds
		}
	}
}

func main() {
	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal("could not open input file")
	}
	var buf []byte = make([]byte, 32*1024)
	count, err := countRounds(f, buf)
	f.Seek(0, 0)
	var rounds []round = make([]round, count)
	rounds = getRounds(f, buf, rounds)

	points := 0
	for _, r := range rounds {
		points += shapePoints(r.player)
		points += outcomePoints(r)
	}
	fmt.Println("first part: ", points)

	points = 0
	for _, r := range rounds {
		points += shapePointsCheating(r)
		points += outcomePointsCheating(r.player)
	}
	fmt.Println("second part: ", points)
}
