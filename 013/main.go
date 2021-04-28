package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = math.MaxInt64 / 4

type Edge struct {
	from int
	to int
	weight int
}

type Node struct {
	node int
	dist int
	index int
}

// PriorityQueue は heap.Interface を実装し，Item のリストを保持します。
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Pop が最小ではなく最大の優先度を持つ項目を返して欲しいので，ここでは > を使っています。
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // メモリリークを避ける
	node.index = -1 // 安全のため
	*pq = old[0 : n-1]
	return node
}

func dijkstra(g [][]Edge, s int) []int {
	n := len(g)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[s] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{
		node: s,
		dist: 0,
	})
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if dist[node.node] < node.dist {
			continue
		}
		for _, e := range(g[node.node]) {
			nextNode := e.to
			nextDist := node.dist + e.weight
			if dist[nextNode] <= nextDist {
				continue
			}
			dist[nextNode] = nextDist
			heap.Push(&pq, &Node{
				node: nextNode,
				dist: nextDist,
			})
		}
	}
	return dist
}

func main() {
	io := NewIo()
	defer io.Flush()
	N, M := io.NextInt(), io.NextInt()
	g := make([][]Edge, N)
	for i := 0; i < M; i++ {
		a, b, c := io.NextInt(), io.NextInt(), io.NextInt()
		a--
		b--
		g[a] = append(g[a], Edge{
			from: a,
			to: b,
			weight: c,
		})
		g[b] = append(g[b], Edge{
			from: b,
			to: a,
			weight: c,
		})
	}

	dist0 := dijkstra(g, 0)
	dist1 := dijkstra(g, N - 1)
	for i := 0; i < N; i++ {
		io.PrintLn(dist0[i] + dist1[i])
	}
}

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextInts(n, v int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = v
	}
	return a
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *Io) PrintIntLn(a []int) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func (io *Io) PrintStringLn(a []string) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func Log(name string, value interface{}) {
	fmt.Fprintf(os.Stderr, "%s=%+v\n", name, value)
}
