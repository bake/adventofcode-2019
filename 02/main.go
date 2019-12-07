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

func main() {
	memory, err := input(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(memory))
	n, v := part2(memory, 19690720)
	fmt.Println(100*n + v)
}

func part1(memory []int) int {
	cpy := make([]int, len(memory))
	copy(cpy, memory)
	cpy[1] = 12
	cpy[2] = 2
	return run(cpy)[0]
}

func part2(in []int, target int) (int, int) {
	out := make([]int, len(in))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(out, in)
			out[1] = noun
			out[2] = verb
			if run(out)[0] == target {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func run(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	for i := 0; i < len(out); i++ {
		switch out[i] {
		case 1:
			out[out[i+3]] = out[out[i+1]] + out[out[i+2]]
			i += 3
		case 2:
			out[out[i+3]] = out[out[i+1]] * out[out[i+2]]
			i += 3
		case 99:
			return out
		}
	}
	return out
}

func input(r io.Reader) ([]int, error) {
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
