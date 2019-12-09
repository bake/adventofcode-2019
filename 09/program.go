package main

import (
	"fmt"
)

type mode int

const (
	position mode = iota
	immediate
	relative
)

func (m mode) String() string {
	switch m {
	case position:
		return "position"
	case immediate:
		return "immediate"
	case relative:
		return "relative"
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
	adjustRelativeBase
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
	case adjustRelativeBase:
		return "adjustRelativeBase"
	case halt:
		return "halt"
	default:
		return fmt.Sprintf("<unknown %d>", o)
	}
}

type program struct {
	ip            int64
	relativeBase  int64
	memory        []int64
	input, output []int64
}

func newProgram(memory []int64) *program {
	var p program
	p.memory = make([]int64, len(memory))
	copy(p.memory, memory)
	return &p
}

func (p *program) at(i int64) int64 {
	if 0 > i || i >= int64(len(p.memory)) {
		return 0
	}
	return p.memory[i]
}

func (p *program) set(i, j int64) {
	for k := int64(len(p.memory)); k <= i; k++ {
		p.memory = append(p.memory, 0)
	}
	p.memory[i] = j
}

func (p *program) run(input ...int64) (output []int64, err error) {
	p.input = append(p.input, input...)
	for i := 0; ; i++ {
		op, err := p.step()
		if err != nil {
			return p.output, err
		}
		if op == halt {
			break
		}
	}
	return p.output, nil
}

func (p *program) step() (operation, error) {
	op := operation(p.memory[p.ip] % 100)
	a := mode(p.memory[p.ip] / 100 % 10)
	b := mode(p.memory[p.ip] / 1000 % 10)
	c := mode(p.memory[p.ip] / 10000 % 10)
	d, e, f := p.ip+1, p.ip+2, p.ip+3
	if a == position {
		d = p.at(d)
	}
	if a == relative {
		d = p.at(d) + p.relativeBase
	}
	if b == position {
		e = p.at(e)
	}
	if b == relative {
		e = p.at(e) + p.relativeBase
	}
	if c == position {
		f = p.at(f)
	}
	if c == relative {
		f = p.at(f) + p.relativeBase
	}
	// fmt.Printf("%12s %04d: %d,%d,%d %d,%d,%d %d,%d,%d\n", op, p.ip, a, d, p.at(d), b, e, p.at(e), c, f, p.at(f))
	switch op {
	case add:
		p.set(f, p.at(d)+p.at(e))
		p.ip += 4
	case multiply:
		p.set(f, p.at(d)*p.at(e))
		p.ip += 4
	case input:
		if len(p.input) == 0 {
			return 0, fmt.Errorf("empty input")
		}
		p.set(d, p.input[0])
		p.input = p.input[1:]
		p.ip += 2
	case output:
		p.output = append(p.output, p.at(d))
		p.ip += 2
	case jumpIfTrue:
		p.ip += 3
		if p.at(d) != 0 {
			p.ip = p.at(e)
		}
	case jumpIfFalse:
		p.ip += 3
		if p.at(d) == 0 {
			p.ip = p.at(e)
		}
	case lessThan:
		var v int64
		if p.at(d) < p.at(e) {
			v = 1
		}
		p.set(f, v)
		p.ip += 4
	case equal:
		var v int64
		if p.at(d) == p.at(e) {
			v = 1
		}
		p.set(f, v)
		p.ip += 4
	case adjustRelativeBase:
		p.relativeBase += p.at(d)
		p.ip += 2
	case halt:
	}
	return op, nil
}
