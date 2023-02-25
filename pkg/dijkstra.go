package algos

const Infinity uint32 = 1<<32 - 1

type Noder interface {
	Neighbors(n Node) []Node
	UpdateOrEnd(n Node) bool
}

type Node interface {
	Index() int
}

type HeapItem struct {
	Node
	index int
}

func dijkstra(searchSpace []Node, start Node, noder Noder) {
}
