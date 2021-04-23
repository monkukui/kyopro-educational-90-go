package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	io := NewIo()
	defer io.Flush()
	H, W := io.NextInt(), io.NextInt()
	uf := NewUnionFind(H * W)
	red := make([][]bool, H)
	for i := 0; i < H; i++ {
		red[i] = make([]bool, W)
	}
	Q := io.NextInt()
	for i := 0; i < Q; i++ {
		t := io.NextInt()
		if t == 1 {
			// マス (r, c) を赤く塗る
			r, c := io.NextInt(), io.NextInt()
			r--
			c--
			red[r][c] = true
			dr := [4]int {0, -1, 0, 1}
			dc := [4]int {1, 0, -1, 0}
			for dir := 0; dir < 4; dir++ {
				nr := r + dr[dir]
				nc := c + dc[dir]
				if nr < 0 || nr >= H || nc < 0 || nc >= W {
					continue
				}
				if red[nr][nc] {
					uf.Merge(r * W + c, nr * W + nc)
				}
			}
		} else {
			// 連結判定
			ra, ca, rb, cb := io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt()
			ra--
			ca--
			rb--
			cb--
			if red[ra][ca] && red[rb][cb] && uf.Same(ra * W + ca, rb * W + cb) {
				io.PrintLn("Yes")
			} else {
				io.PrintLn("No")
			}
		}
	}
}

type UnionFind struct {
	rank   []int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	d := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.rank[i] = 0
	}
	return d
}

func (d *UnionFind) Merge(a, b int) bool {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return false
	}
	if d.rank[x] < d.rank[y] {
		x, y = y, x
	}
	if d.rank[x] == d.rank[y] {
		d.rank[x]++
	}
	d.parent[y] = x
	return true
}

func (d *UnionFind) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

func (d *UnionFind) Leader(a int) int {
	for d.parent[a] != a {
		a = d.parent[a]
	}
	return a
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
