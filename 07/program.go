package main

import (
	"fmt"
)

type mode int

const (
	position mode = iota
	immediate
)

func (m mode) String() string {
	switch m {
	case position:
		return "position"
	case immediate:
		return "immediate"
	default:
		return fmt.Sprintf("<unknown %d>", m)
	}
}

type operation int

const (
	_ operation = iota
	add
	multiply
	input
	output
	jumpIfTrue
	jumpIfFalse
	lessThan
	equal
	halt = 99
)

func (o operation) String() string {
	switch o {
	case add:
		return "add"
	case multiply:
		return "multiply"
	case input:
		return "input"
	case output:
		return "output"
	case jumpIfTrue:
		return "jump if true"
	case jumpIfFalse:
		return "jump if false"
	case lessThan:
		return "less than"
	case equal:
		return "equal"
	case halt:
		return "halt"
	default:
		return fmt.Sprintf("<unknown %d>", o)
	}
}

type program struct {
	ip            int
	memory        []int
	input, output []int
}

func newProgram(memory []int) *program {
	var p program
	p.memory = make([]int, len(memory))
	copy(p.memory, memory)
	return &p
}

func (p *program) run() (operation, error) {
	for p.ip < len(p.memory)-1 {
		op, err := p.step()
		if err != nil {
			return 0, err
		}
		switch op {
		case halt:
			return halt, nil
		case output:
			return output, nil
		}
	}
	return 0, nil
}

func (p *program) step() (operation, error) {
	op := operation(p.memory[p.ip] % 100)
	a := mode(p.memory[p.ip] / 100 % 10)
	b := mode(p.memory[p.ip] / 1000 % 10)
	c := mode(p.memory[p.ip] / 10000 % 10)
	d, e, f := p.ip+1, p.ip+2, p.ip+3
	if a == position && d < len(p.memory) {
		d = p.memory[d]
	}
	if b == position && e < len(p.memory) {
		e = p.memory[e]
	}
	if c == position && f < len(p.memory) {
		f = p.memory[f]
	}
	switch op {
	case add:
		p.memory[f] = p.memory[d] + p.memory[e]
		p.ip += 4
	case multiply:
		p.memory[f] = p.memory[d] * p.memory[e]
		p.ip += 4
	case input:
		if len(p.input) == 0 {
			return 0, fmt.Errorf("empty input")
		}
		p.memory[d], p.input = p.input[0], p.input[1:]
		p.ip += 2
	case output:
		if p.memory[d] != 0 {
			p.output = append(p.output, p.memory[d])
		}
		p.ip += 2
	case jumpIfTrue:
		p.ip += 3
		if p.memory[d] != 0 {
			p.ip = p.memory[e]
		}
	case jumpIfFalse:
		p.ip += 3
		if p.memory[d] == 0 {
			p.ip = p.memory[e]
		}
	case lessThan:
		p.memory[f] = 0
		if p.memory[d] < p.memory[e] {
			p.memory[f] = 1
		}
		p.ip += 4
	case equal:
		p.memory[f] = 0
		if p.memory[d] == p.memory[e] {
			p.memory[f] = 1
		}
		p.ip += 4
	case halt:
	}
	return op, nil
}
