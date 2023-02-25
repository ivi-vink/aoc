package main

import "fmt"

type flyer interface {
	fly() string
	dive() string
}

type printer func(s string) string

func (p printer) fly() string {
	return p("I'm flying!")
}

func (p printer) dive() string {
	return p("I'm diving!")
}

func printAndReturn(s string) string {
	fmt.Println(s)
	return s
}

type bird struct {
	flyer
}

func main() {
	bf := bird{
		flyer: printer(printAndReturn),
	}
	bf.fly()

	bs.fly()
}
