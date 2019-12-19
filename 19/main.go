package main

import (
	"bytes"
	"fmt"
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
	fmt.Println(part1(mem, 50, 50))
	fmt.Println(part2(mem, 100, 100))
}

func part1(mem []int64, w, h int64) int64 {
	var sum int64
	for y := int64(0); y < w; y++ {
		for x := int64(0); x < h; x++ {
			sum += at(mem, int64(x), int64(y))
		}
	}
	return sum
}

// part2 tl;dr: bruteforce.
func part2(mem []int64, w, h int64) int64 {
	for x := int64(0); ; x++ {
		for y := int64(0); ; y++ {
			if at(mem, x+w-1, y) == 0 {
				continue
			}
			if at(mem, x, y+h-1) == 0 {
				break
			}
			return x*10_000 + y
		}
	}
}

// at returns the output at a given position.
func at(mem []int64, x, y int64) int64 {
	p := newProgram(mem)
	p.input = append(p.input, x, y)
	res, _ := p.take(1)
	return res[0]
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
