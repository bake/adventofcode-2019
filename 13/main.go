package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	render := flag.Bool("render", false, "Render the game")
	flag.Parse()
	mem, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(mem))
	fmt.Println(part2(mem, *render))
}

func part1(mem []int64) int {
	p := newProgram(mem)
	g := grid{}
	for {
		out, err := p.take(3)
		if err != nil {
			break
		}
		g[point{out[0], out[1]}] = tile(out[2])
	}

	var n int
	for _, t := range g {
		if t == block {
			n++
		}
	}
	return n
}

func part2(mem []int64, render bool) int64 {
	mem[0] = 2
	p := newProgram(mem)
	g := grid{}
	var s int64
	var px, bx int64
	for {
		for {
			out, err := p.take(3)
			if err != nil {
				break
			}
			p := point{out[0], out[1]}
			if p.x == -1 && p.y == 0 {
				s = out[2]
				continue
			}
			switch tile(out[2]) {
			case ball:
				bx = p.x
			case paddle:
				px = p.x
			}
			g[p] = tile(out[2])
		}
		if render {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Score:", s)
			fmt.Println(g)
			time.Sleep(50 * time.Millisecond)
		}

		var input int64
		if bx < px {
			input = -1
		}
		if bx > px {
			input = 1
		}
		p.input = append(p.input, input)

		var blocks int
		for _, t := range g {
			if t == block {
				blocks++
			}
		}
		if blocks == 0 {
			break
		}
	}
	return s
}

func parseInput(r io.Reader) ([]int64, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	bs := bytes.Split(bytes.TrimSpace(input), []byte(","))
	memory := make([]int64, len(bs))
	for i, b := range bs {
		mem, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			return nil, err
		}
		memory[i] = mem
	}
	return memory, nil
}
