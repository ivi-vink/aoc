package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	items       []*item
	operation   func(item *item) int
	test        func(item *item) *monkey
	by          int
	inspections int
}

func NewMonkey(items []int, by int, op func(i *item) int, test func(i *item) *monkey) *monkey {
	itemPtrs := make([]*item, len(items))
	for i, worry := range items {
		itemPtrs[i] = NewItem(worry)
	}
	return &monkey{items: itemPtrs, operation: op, test: test, by: by}
}

func (m *monkey) throw(receiver *monkey) {
	throw, leftover := m.items[0], m.items[1:]
	m.items = leftover
	receiver.items = append(receiver.items, throw)
}

type item struct {
	worry int
}

func (i *item) operation(m *monkey, r, mod int) {
	worry := m.operation(i)
	m.inspections += 1
	i.worry = (worry / r) % mod
}

func (i *item) test(m *monkey) *monkey {
	return m.test(i)
}

func NewItem(worry int) *item {
	return &item{
		worry: worry,
	}
}

// Can someone tell me how to do this better? Multiline regex?
func makeMonkeys(monkeyCount int, f *os.File) []*monkey {
	startingItemsRegex := regexp.MustCompile(`\s+Starting items: ([0-9 ,]+)`)
	operationRegex := regexp.MustCompile(`\s+Operation: new = ([0-9A-Za-z]+) ([+\*]{1}) ([0-9A-Za-z]+)`)
	testRegex := regexp.MustCompile(`\s+Test: divisible by (\d+)`)
	testConditionRegex := regexp.MustCompile(`\s+If (true|false): throw to monkey (\d+)`)

	monkeys := make([]*monkey, monkeyCount)
	makeTest := func(by int, mtrue, mfalse int) func(i *item) *monkey {
		return func(i *item) *monkey {
			if i.worry%by == 0 {
				return monkeys[mtrue]
			} else {
				return monkeys[mfalse]
			}
		}
	}
	makeOp := func(a1, comb, a2 string) func(i *item) int {
		b := []string{a1, a2}
		var c func(a, b int) int
		if comb == "+" {
			c = func(a, b int) int {
				return a + b
			}
		} else if comb == "*" {
			c = func(a, b int) int {
				return a * b
			}
		}
		b0, _ := strconv.Atoi(b[0])
		b1, _ := strconv.Atoi(b[1])
		return func(i *item) int {
			switch {
			case b[0] == "old" && b[1] == "old":
				return c(i.worry, i.worry)
			case b[0] == "old":
				return c(i.worry, b1)
			case b[1] == "old":
				return c(b0, i.worry)
			default:
				return c(b0, b1)
			}
		}
	}

	f.Seek(0, 0)
	s := bufio.NewScanner(f)
	startingItems := []int{}
	testMap := make(map[string]int)
	opMap := make(map[string]string)
	monkeyPtr := 0
	for {
		eof := s.Scan()

		if !eof || len(s.Bytes()) == 0 {
			monkeys[monkeyPtr] = NewMonkey(
				startingItems,
				testMap["by"],
				makeOp(opMap["arg1"], opMap["combination"], opMap["arg2"]),
				makeTest(testMap["by"], testMap["true"], testMap["false"]),
			)
			monkeyPtr += 1
			testMap = make(map[string]int)
			opMap = make(map[string]string)
			if monkeyPtr >= len(monkeys) {
				break
			}
		}

		m := startingItemsRegex.FindStringSubmatch(s.Text())
		if len(m) == 2 {
			itemStr := strings.Split(m[1], ", ")
			startingItems = make([]int, len(itemStr))
			for i, strItem := range itemStr {
				v, err := strconv.Atoi(strItem)
				if err != nil {
					log.Fatal("Could not get item worry from string", err)
				}
				startingItems[i] = v
			}
		}
		m = operationRegex.FindStringSubmatch(s.Text())
		if len(m) == 4 {
			opMap["arg1"] = m[1]
			opMap["combination"] = m[2]
			opMap["arg2"] = m[3]
		}
		m = testRegex.FindStringSubmatch(s.Text())
		if len(m) == 2 {
			by, err := strconv.Atoi(m[1])
			if err != nil {
				log.Fatal("Could not convert divisible by", err)
			}
			testMap["by"] = by
		}
		m = testConditionRegex.FindStringSubmatch(s.Text())
		if len(m) == 3 {
			to, err := strconv.Atoi(m[2])
			if err != nil {
				log.Fatal("Could not convert monkey to throw to", err)
			}
			testMap[m[1]] = to
		}
	}
	return monkeys
}

func monkeyBusiness(monkeys []*monkey, rounds, worryFactor, mod int) int {
	for round := 0; round < rounds; round++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				item := m.items[0]
				item.operation(m, worryFactor, mod)
				to := item.test(m)
				m.throw(to)
			}
		}
	}

	monkey, business := 0, 0
	for _, m := range monkeys {
		if m.inspections > monkey {
			monkey, business = m.inspections, monkey
		} else if m.inspections > business {
			monkey, business = monkey, m.inspections
		}
	}
	return monkey * business
}

func main() {
	f, err := os.Open("day11.txt")
	if err != nil {
		log.Fatal("could open input file", err)
	}

	monkeyRegex := regexp.MustCompile(`Monkey (\d+)`)
	monkeyCount := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := monkeyRegex.FindStringSubmatch(s.Text())
		if len(m) > 0 {
			monkeyCount += 1
		}
	}

	monkeys := makeMonkeys(monkeyCount, f)
	mod := 1
	for _, m := range monkeys {
		mod *= m.by
	}
	mb := monkeyBusiness(monkeys, 20, 3, mod)
	fmt.Println("Part 1:", mb)

	monkeys = makeMonkeys(monkeyCount, f)
	mb = monkeyBusiness(monkeys, 10_000, 1, mod)
	fmt.Println("Part 2:", mb)
}
