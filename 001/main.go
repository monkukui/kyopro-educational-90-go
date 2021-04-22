package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := NewScanner(os.Stdin)
	N := scanner.Int()
	L := scanner.Int()
	K := scanner.Int()
	A := scanner.Ints(N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	A = append(A, L)
	N++
	for i := N - 1; i >= 1; i-- {
		A[i] -= A[i-1]
	}

	ok := 1
	ng := L + 1
	for ng-ok > 1 {
		x := (ok + ng) / 2
		cnt := 0
		sum := 0
		for i := 0; i < N; i++ {
			sum += A[i]
			if sum >= x {
				sum = 0
				cnt++
			}
		}
		if cnt >= K+1 {
			ok = x
		} else {
			ng = x
		}
	}
	fmt.Println(ok)
}

type Scanner struct {
	r     io.Reader
	buf   []byte
	split func(byte) bool
	size, eof,
	start, end int
}

func NewScanner(r io.Reader) *Scanner { return NewScannerSize(r, 4096) }
func NewScannerSize(r io.Reader, size int) *Scanner {
	return &Scanner{
		r:     r,
		buf:   make([]byte, size),
		split: func(b byte) bool { return b == ' ' || b == '\n' },
		size:  size, start: size, end: size,
	}
}

func (s *Scanner) Split(f func(byte) bool) { s.split = f }
func (s *Scanner) Bytes() (b []byte) {
	unsafe, size := s.next()
	b = make([]byte, size)
	copy(b, unsafe)
	return
}
func (s *Scanner) UnsafeBytes() []byte { return s.bytes() }
func (s *Scanner) Text() string        { return string(s.bytes()) }
func (s *Scanner) Int() int            { return ParseInt(s.bytes()) }
func (s *Scanner) Float() float64      { return ParseFloat(s.bytes()) }

func (s *Scanner) Texts(len int) []string {
	a := make([]string, len)
	for i := 0; i < len; i++ {
		a[i] = s.Text()
	}
	return a
}

func (s *Scanner) Ints(len int) []int {
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = s.Int()
	}
	return a
}

func (s *Scanner) IntsMin(len int) ([]int, int) {
	a, min := make([]int, len), 0
	for i := 0; i < len; i++ {
		a[i] = s.Int()
		if i == 0 || a[i] < min {
			min = a[i]
		}
	}
	return a, min
}

func (s *Scanner) IntsMax(len int) ([]int, int) {
	a, max := make([]int, len), 0
	for i := 0; i < len; i++ {
		a[i] = s.Int()
		if i == 0 || a[i] > max {
			max = a[i]
		}
	}
	return a, max
}

func (s *Scanner) Floats(len int) []float64 {
	a := make([]float64, len)
	for i := 0; i < len; i++ {
		a[i] = s.Float()
	}
	return a
}

func (s *Scanner) bytes() []byte { b, _ := s.next(); return b }
func (s *Scanner) next() (b []byte, size int) {
	for n := 0; ; {
		for i := s.end; i < s.size; i++ {
			if s.split(s.buf[i]) || i == s.eof {
				b, size = s.buf[s.start:i], i-s.start
				s.start, s.end = i+1, i+1
				return
			}
		}
		s.end = s.size
		if n++; n == 2 { // 無限ループ対策
			panic("lack of buffer size")
		}
		s.read()
	}
}
func (s *Scanner) read() {
	copy(s.buf, s.buf[s.start:s.end])
	s.end -= s.start
	s.start = 0
	n, _ := s.r.Read(s.buf[s.end:])
	s.eof = s.end + n
}

func ParseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func ParseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}
