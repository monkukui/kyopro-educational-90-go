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
	N, K := io.NextInt(), io.NextInt()
	S := io.NextLine()
	ids := make([][]int, 26)
	for i := 0; i < 26; i++ {
		ids[i] = make([]int, N)
	}
	for j := N - 1; j >= 0; j-- {
		for c := 'a'; c <= 'z'; c++ {
			if uint8(c) == S[j] {
				ids[c-'a'][j] = j
			} else {
				ids[c-'a'][j] = N
			}
			if j != N-1 && ids[c-'a'][j] > ids[c-'a'][j+1] {
				ids[c-'a'][j] = ids[c-'a'][j+1]
			}
		}
	}
	i := 0
	ans := ""
	for K > 0 {
		for c := 'a'; c <= 'z'; c++ {
			id := ids[c-'a'][i]
			if id == N {
				continue
			}
			// これを取った時の、残りの文字数
			rest := N - id
			// 取れるなら、貪欲に取る
			if K <= rest {
				K--
				i = id + 1
				ans += string(c)
				break
			}
		}
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
