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
	memory, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(run(memory, 1)[0])
	fmt.Println(run(memory, 2)[0])
}

func run(memory []int64, input ...int64) []int64 {
	p := newProgram(memory)
	out, _ := p.run(input...)
	return out
}

func parseInput(r io.Reader) (memory []int64, err error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, value := range strings.Split(strings.TrimSpace(string(input)), ",") {
		mem, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		memory = append(memory, mem)
	}
	return memory, nil
}
