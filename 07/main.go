package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gitchander/permutation"
)

func main() {
	memory, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(memory))
	fmt.Println(part2(memory))
}

func part1(memory []int) int {
	phases := []int{0, 1, 2, 3, 4}
	perm := permutation.New(permutation.IntSlice(phases))
	var max int
	for perm.Next() {
		out := []int{0}
		for _, phase := range phases {
			p := newProgram(memory)
			p.input = append([]int{phase}, out...)
			p.run()
			out = p.output
		}
		p := newProgram(memory)
		p.input = append(p.input, 0)

		if len(out) > 0 && max < out[0] {
			max = out[0]
		}
	}
	return max
}

func part2(memory []int) (max int) {
	phases := []int{5, 6, 7, 8, 9}
	perm := permutation.New(permutation.IntSlice(phases))
	for perm.Next() {

		programs := make([]*program, len(phases))
		for i := range programs {
			programs[i] = newProgram(memory)
			programs[i].input = append(programs[i].input, phases[i])
		}
		programs[0].input = append(programs[0].input, 0)

		var halted int
	run:
		for halted < len(programs) {
			for i, p := range programs {
				op, err := p.run()
				if err != nil {
					halted++
					break run
				}
				switch op {
				case halt:
					halted++
				case output:
					q := programs[(i+1)%len(programs)]
					q.input = append(q.input, p.output...)
					p.output = nil
				}
			}
		}
		p := programs[0]
		if len(p.input) > 0 && max < p.input[0] {
			max = p.input[0]
		}
	}
	return max
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
