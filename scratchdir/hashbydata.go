package main

import "fmt"

type (
	empty struct{}
	t     interface{}
	set   map[t]empty
)

func main() {
	m := make(map[t]string)
	fmt.Println(fmt.Sprintf("%+v", m))
	m[struct{ x string }{x: "hello there"}] = "hello there"
	fmt.Println(fmt.Sprintf("%+v", m))

	m[struct{ x string }{x: "hello there"}] = "goodbye"
	fmt.Println(fmt.Sprintf("%+v", m))

	ptr := &struct{ x string }{x: "hello there"}
	pv := *ptr
	m[pv] = "ptr"
	fmt.Println(fmt.Sprintf("%+v", m))
}
