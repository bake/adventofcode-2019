package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	mem, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, part1(mem))

	g := part2(mem)
	if err := png.Encode(os.Stdout, g); err != nil {
		log.Fatal(err)
	}
}

func part1(mem []int64) int {
	p := newProgram(mem)
	g := grid{}
	run(p, g)
	return len(g)
}

func part2(mem []int64) grid {
	p := newProgram(mem)
	g := grid{point{0, 0}: white}
	run(p, g)
	return g
}

func run(p *program, g grid) {
	r := &robot{}
	for {
		c, d, err := step(p, g[r.point])
		if err != nil {
			break
		}
		g[r.point] = c
		r.rotate(d)
		r.step()
	}
	return
}

func step(p *program, c spaceColor) (spaceColor, direction, error) {
	p.input = append(p.input, int64(c))
	for i := 0; i < 2; {
		op, err := p.step()
		if err != nil {
			return 0, 0, err
		}
		switch op {
		case output:
			i++
		case halt:
			return 0, 0, fmt.Errorf("program halted")
		}
	}
	c, d := spaceColor(p.output[0]), direction(p.output[1])
	// I use the the order up, right, down, left to simplify rotation.
	switch d {
	case 0:
		d = left
	case 1:
		d = right
	}
	p.output = nil
	return c, d, nil
}

func parseInput(r io.Reader) (memory []int64, err error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, value := range bytes.Split(bytes.TrimSpace(input), []byte(",")) {
		mem, err := strconv.ParseInt(string(value), 10, 64)
		if err != nil {
			return nil, err
		}
		memory = append(memory, mem)
	}
	return memory, nil
}
