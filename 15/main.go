package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	mem, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(mem))
	fmt.Println(part2(mem, 1_000_000))
}

// part1 is solved by spaceologies wishful thinking technique, in that it lets
// the droid move randomly until it stumbles upon the oxygen source. It might
// never find it. It might arrive there on a way thats more complex than the
// optional one. In either case, the breadth first search used to find the
// shortest path in this maze consisting of only one valid path, might not even
// be able to find the true shortest path. If that happens, restart the program.
// And hope it works.
func part1(mem []int64) int {
	p := newProgram(mem)
	g, r := grid{}, point{}
	for {
		d := direction(1 + rand.Intn(4))
		p.put(int64(d))
		res, err := p.take(1)
		if err != nil {
			log.Fatal(err)
		}
		next := r.move(d)
		switch res[0] {
		case 0:
			g[next] = wall
		case 1:
			g[next] = free
			r = next
		case 2:
			g[next] = target
			return bfs(g, point{})
		}
	}
}

func bfs(g grid, s point) int {
	queue := []point{s}
	distance := map[point]int{}
	var p point
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		for d := north; d <= east; d++ {
			q := p.move(d)
			if dist, ok := distance[q]; ok && dist < distance[p]+1 {
				continue
			}
			switch g[q] {
			case free:
				distance[q] = distance[p] + 1
				queue = append(queue, q)
			case target:
				return distance[p] + 1
			}
		}
	}
	return 0
}

// So a few things happend since I wrote part 1. I had a walk, watched a movie
// and remembered that there was a second part. So here I am, copying the
// previous solution and trying to make it even more stupid by ignoring the
// oxygen source completely and instead letting the droid take n steps and
// randomly wander through the maze.
// But no matter how bad I control the droid, there needs to be a working BFS
// this time. And someone should double check that there are no loops.
func part2(mem []int64, n int) int {
	p := newProgram(mem)
	g, r, o := grid{}, point{}, point{}
	for i := 0; i < n; i++ {
		d := direction(1 + rand.Intn(4))
		p.put(int64(d))
		res, err := p.take(1)
		if err != nil {
			log.Fatal(err)
		}
		next := r.move(d)
		switch res[0] {
		case 0:
			g[next] = wall
		case 1:
			g[next] = free
			r = next
		case 2:
			g[next] = target
			r, o = next, next
		}
	}
	return longestPath(g, o)
}

func longestPath(g grid, p point) int {
	queue := []point{p}
	distance := map[point]int{p: 0}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		// Never seen such awfully named variables in a for loop.
		for d := north; d <= east; d++ {
			q := p.move(d)
			if g[q] == wall {
				continue
			}
			// What? I'm not even comparing the distances I save? Why would I even
			// store integers instead of empty structs? As long as there really are no
			// loops ...
			if _, ok := distance[q]; ok {
				continue
			}
			distance[q] = distance[p] + 1
			queue = append(queue, q)
		}
	}
	var max int
	for _, d := range distance {
		if d > max {
			max = d
		}
	}
	return max
}

func draw(g grid, p point) {
	b := g.Bounds()
	fmt.Print("\033[H\033[2J")
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if x == 0 && y == 0 {
				fmt.Print("o")
				continue
			}
			if p.x == x && p.y == y {
				fmt.Printf("D")
				continue
			}
			fmt.Print(g[point{x, y}])
		}
		fmt.Println()
	}
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
