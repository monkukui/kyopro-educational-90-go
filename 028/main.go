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
	imos := make([][]int, 1002)
	for i := 0; i < 1002; i++ {
		imos[i] = make([]int, 1002)
	}
	for i := 0; i < N; i++ {
		lx, ly, rx, ry := io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt()
		imos[lx][ly]++
		imos[rx][ry]++
		imos[lx][ry]--
		imos[rx][ly]--
	}
	for x := 0; x <= 1001; x++ {
		for y := 1; y <= 1001; y++ {
			imos[x][y] += imos[x][y - 1]
		}
	}
	for y := 0; y <= 1001; y++ {
		for x := 1; x <= 1001; x++ {
			imos[x][y] += imos[x - 1][y]
		}
	}
	ans := make([]int, N + 1)
	for x := 0; x <= 1001; x++ {
		for y := 0; y <= 1001; y++ {
			ans[imos[x][y]]++
		}
	}
	for i := 1; i <= N; i++ {
		io.PrintLn(ans[i])
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
