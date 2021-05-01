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
	g := make([][]int, N)
	for i := 0; i < N - 1; i++ {
		a, b := io.NextInt(), io.NextInt()
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	color := io.NextInts(N, -1)
	var dfs func(cur, par, c int)
	dfs = func(cur, par, c int) {
		color[cur] = c
		for _, nxt := range g[cur] {
			if nxt == par {
				continue
			}
			dfs(nxt, cur, 1 - c)
		}
	}
	dfs(0, -1, 0)
	red := make([]int, 0, N)
	blue := make([]int, 0, N)

	for i := 0; i < N; i++ {
		if color[i] == 0 {
			red = append(red, i)
		} else if color[i] == 1 {
			blue = append(blue, i)
		} else {
			panic("")
		}
	}
	if len(red) < len(blue) {
		red, blue = blue, red
	}
	ans := make([]int, N / 2)
	for i := 0; i < N / 2; i++ {
		ans[i] = red[i] + 1
	}
	io.PrintIntLn(ans)
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
