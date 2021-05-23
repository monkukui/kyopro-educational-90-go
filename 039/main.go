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
	N := io.NextInt()
	a := make([]int, N - 1)
	b := make([]int, N - 1)
	g := make([][]int, N)
	for i := 0; i < N - 1; i++ {
		a[i], b[i] = io.NextInt(), io.NextInt()
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}
	size := make([]int, N)
	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		childSizeSum := 0
		for _, nxt := range g[cur] {
			if nxt == par {
				continue
			}
			dfs(nxt, cur)
			childSizeSum += size[nxt]
		}
		size[cur] = childSizeSum + 1
	}
	dfs(0, -1)

	ans := 0
	for i := 0; i < N - 1; i++ {
		u, v := a[i], b[i]
		if size[v] < size[u] {
			u, v = v, u
		}
		c := size[u]
		ans += c * (N - c)
	}
	io.PrintLn(ans)
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
