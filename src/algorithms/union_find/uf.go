package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func BuildUF(path string) *UF {
	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\r\n")
	var uf *UF
	var n int
	for i, line := range lines {
		if i == 0 {
			n, _ = strconv.Atoi(line)
			uf = NewUF(n)
		} else {
			p, q := getPairOfInts(line)
			if !uf.connected(p, q) {
				uf.union(p, q)
				//fmt.Println("p + q", p, q)
			}
		}
	}
	return uf
}

func getPairOfInts(input string) (int, int) {
	s := strings.Split(input, " ")
	p, _ := strconv.Atoi(s[0])
	q, _ := strconv.Atoi(s[1])
	return p, q
}

func GetFileReader(path string) io.Reader {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return reader
}

func main() {
	// API p. 219
	// final 229
	if len(os.Args) < 2 {
		panic("Need file path")
	}
	path := os.Args[1]
	uf := BuildUF(path)
	fmt.Println(uf.count)
	/*
	fmt.Println(uf.ids)
	fmt.Println(uf.sizes)
	*/
	panic("Take notes from page 233")
}

func NewUF(n int) *UF {
	ids := make([]int, n, n)
	sizes := make([]int, n, n)
	for i := range ids {
		ids[i] = i
	}
	for i := range sizes {
		sizes[i] = 1
	}
	return &UF{
		ids,
		sizes,
		n,
	}
}

type UF struct {
	ids			[]int
	sizes		[]int
	count		int
}

func (self *UF) connected(p, q int) bool {
	return self.find(p) == self.find(q)
}

func (self *UF) find(p int) int {
	for p != self.ids[p] {
		p = self.ids[p]
	}
	return p
}

func (self *UF) union(p, q int) {
	i := self.find(p)
	j := self.find(q)
	if i == j {
		return
	}
	if self.sizes[i] < self.sizes[j] {
		self.ids[i] = j
		self.sizes[j] += self.sizes[i]
	} else {
		self.ids[j] = i
		self.sizes[i] += self.sizes[j]
	}
	self.count -= 1
}


// Exercises 235