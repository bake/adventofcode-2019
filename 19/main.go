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
}

func part1(mem []int64, w, h int) int {
	g := grid{}
	var sum int
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			p := newProgram(mem)
			p.input = append(p.input, int64(x), int64(y))
			res, _ := p.take(1)
			sum += int(res[0])
			if res[0] == 1 {
				g[point{x, y}] = beam
			}
		}
	}
	return sum
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
