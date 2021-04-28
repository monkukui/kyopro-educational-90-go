package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	io := NewIo()
	defer io.Flush()
	N := io.NextInt()
	type Task struct {
		D int
		C int
		S int
	}
	tasks := make([]Task, N)
	for i := 0; i < N; i++ {
		tasks[i].D, tasks[i].C, tasks[i].S = io.NextInt(), io.NextInt(), io.NextInt()
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].D < tasks[j].D
	})

	dp := make([]int, 5001)
	for i := 0; i < N; i++ {
		ndp := make([]int, 5001)
		D, C, S := tasks[i].D, tasks[i].C, tasks[i].S
		for j := 0; j <= 5000; j++ {

			// とらない
			ndp[j] = max(ndp[j], dp[j])

			// とる
			if j+C <= D {
				ndp[j+C] = max(ndp[j+C], dp[j]+S)
			}
		}
		dp, ndp = ndp, dp
	}

	ans := 0
	for j := 0; j <= 5000; j++ {
		ans = max(ans, dp[j])
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
