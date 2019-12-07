package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type mode int

const (
	position mode = iota
	immediate
)

type opcode int

const (
	_ opcode = iota
	add
	multiply
	input
	output
	jumpIfTrue
	jumpIfFalse
	lessThan
	equal
	halt opcode = 99
)

func main() {
	memory, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	cpy := make([]int, len(memory))

	copy(cpy, memory)
	_, code := run(cpy, []int{1})
	fmt.Println(code)

	copy(cpy, memory)
	_, code = run(cpy, []int{5})
	fmt.Println(code)
}

func run(in, args []int) ([]int, string) {
	out := make([]int, len(in))
	var str string
	copy(out, in)
	var i int
	for i < len(out)-1 {
		a, b, c, op := parseOpcode(out[i])
		d, e, f := i+1, i+2, i+3
		if a == position && int(a) < len(out) {
			d = out[d]
		}
		if b == position && int(e) < len(out) {
			e = out[e]
		}
		if c == position && int(f) < len(out) {
			f = out[f]
		}
		switch op {
		case add:
			out[f] = out[d] + out[e]
			i += 4
		case multiply:
			out[f] = out[d] * out[e]
			i += 4
		case input:
			out[d], args = args[0], args[1:]
			i += 2
		case output:
			if out[d] != 0 {
				str += strconv.Itoa(out[d])
			}
			i += 2
		case jumpIfTrue:
			i += 3
			if out[d] != 0 {
				i = out[e]
			}
		case jumpIfFalse:
			i += 3
			if out[d] == 0 {
				i = out[e]
			}
		case lessThan:
			out[f] = 0
			if out[d] < out[e] {
				out[f] = 1
			}
			i += 4
		case equal:
			out[f] = 0
			if out[d] == out[e] {
				out[f] = 1
			}
			i += 4
		case halt:
			return out, str
		}
	}
	return out, str
}

func parseOpcode(n int) (a, b, c mode, op opcode) {
	op = opcode(n % 100)
	a = mode(n / 100 % 10)
	b = mode(n / 1000 % 10)
	c = mode(n / 10000 % 10)
	return
}

func parseInput(r io.Reader) ([]int, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var memory []int
	for _, value := range strings.Split(strings.TrimSpace(string(input)), ",") {
		mem, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		memory = append(memory, mem)
	}
	return memory, nil
}
