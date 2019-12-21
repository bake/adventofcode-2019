package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	g, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(g))
}

// 626 is too low.
func part1(g grid) int {
	labels, dst, src := g.labels("ZZ", "AA")
	var p point
	queue := []point{src}
	depth := map[point]int{}
	seen := map[point]struct{}{}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		if _, ok := seen[p]; ok {
			continue
		}
		if q, ok := labels[p]; ok {
			depth[q] = depth[p] + 1
			queue = append(queue, q)
		}
		seen[p] = struct{}{}
		if p == dst {
			return depth[p]
		}
		for _, q := range g.adjacent(p) {
			depth[q] = depth[p] + 1
			queue = append(queue, q)
		}
	}

	return 0
}

func parseInput(r io.Reader) (grid, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bs), "\n"), nil
}
