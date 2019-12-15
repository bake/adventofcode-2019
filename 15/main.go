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
	// render := flag.Bool("render", false, "Render the game")
	// flag.Parse()
	mem, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(mem))
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
		next := point(r).move(d)
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
			switch g[q] {
			case free:
				if dist, ok := distance[q]; ok && dist < distance[p]+1 {
					continue
				}
				distance[q] = distance[p] + 1
				queue = append(queue, q)
			case target:
				return distance[p] + 1
			}
		}
	}
	return 0
}

func draw(g grid, r point) {
	b := g.Bounds()
	fmt.Print("\033[H\033[2J")
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if x == 0 && y == 0 {
				fmt.Print("o")
				continue
			}
			if r.x == int64(x) && r.y == int64(y) {
				fmt.Printf("D")
				continue
			}
			fmt.Print(g[point{int64(x), int64(y)}])
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
