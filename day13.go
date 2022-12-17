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

type compcode int

const (
	cont       compcode = -1
	rightOrder compcode = 0
	wrongOrder compcode = 1
)

// not very idiomatich i think
func newComparePackets(comp func(a, b int) compcode) func (p1, p2 *packet) (int, compcode) {
    return func (p1, p2 *packet) (int, compcode) {
        var cp func(i int, ok compcode, a, b []t) (int, compcode)

        cp = func(i int, ok compcode, a, b []t) (int, compcode) {
            if al, bl := i >= len(a), i >= len(b); al {
                if bl && al {
                    return i, cont
                }
                if bl && !al {
                    return i, wrongOrder
                }
                return i, rightOrder
            }

            if al, bl := i >= len(a), i >= len(b); bl {
                if bl && al {
                    return i, cont
                }
                if bl && !al {
                    return i, wrongOrder
                }
                return i, rightOrder
            }

            at, bt := intOrList(a[i]), intOrList(b[i])
            code := cont
            if at == 1 && bt == 1 {
                _, code = cp(0, ok, a[i].([]t), b[i].([]t))
            } else if at == 1 && bt == 0 {
                _, code = cp(0, ok, a[i].([]t), []t{b[i]})
            } else if at == 0 && bt == 1 {
                _, code = cp(0, ok, []t{a[i]}, b[i].([]t))
            } else if at == 0 && bt == 0 {
                code = comp(a[i].(int), b[i].(int))
            } else {
                log.Fatal("Unexpected input types")
            }

            if code != cont {
                return i, code
            }

            return cp((i + 1), code, a, b)
        }
        return cp(0, -1, p1.data, p2.data)
    }

}

func parsePacket(b []byte) *packet {
	var parseList func(list []t, b []byte) []t

	parseList = func(list []t, b []byte) []t {
		for i := 0; i < len(b); i++ {
			if b[i] == '[' && b[i+1] == ']' {
				list = append(list, []t{})
			} else if b[i] == '[' {
				j := 1
				subchildren := 1
				for subchildren > 0 {
					if i+j >= len(b) {
						break
					}
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
			} else {
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
					continue
				}
				list = append(list, num)
				i = i + j
			}
		}
		return list
	}

	return &packet{data: parseList([]t{}, b[1:len(b)-1])}
}

func swap(packets []*packet, key, i int) []*packet {
    packets[i], packets[key] = packets[key], packets[i]
    return packets
}

func quicksort(packets []*packet, comp func (p1, p2 *packet) (int, compcode)) []*packet {
    var qs func () 
    qs := func () {
        key := len(packets) - 1
        for i:=key-1; i>0; i-- {
            _, code := comp(packets[key], packets[i])
            if code == rightOrder {
                packets = swap(packets, key, i)
            }
        }
        if key <= 0 || key + 1 > len(packets) {
            return packets
        } else {
            left, right := packets[:key], packets[key+1:]
            quicksort()
            quicksort()
        }
    }
    return packets
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
	runningSum := 0
    comparePackets := newComparePackets(func(a, b int) compcode {
			if a == b {
				return cont
			}
			if a < b {
				return rightOrder
			} else {
				return wrongOrder
			}
		})
	for i := 0; i < len(packets); i += 2 {
		_, code := comparePackets(packets[i], packets[i+1])
		if code == rightOrder {
			runningSum += (i / 2) + 1
		}
	}
	fmt.Println("Part 1:", runningSum)

	// pick key last element
	// sort all elements according to key
	// swap key into place
	// divide array
	// repeat

	fmt.Println("Part 2:", quicksort(packets, comparePackets))
}
