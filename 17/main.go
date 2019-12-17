package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	mem, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(mem))
	fmt.Println(part2(mem))
}

func part1(mem []int64) int {
	g := readGrid(mem)
	var sum int
	for _, p := range intersections(g) {
		sum += p.x * p.y
	}
	return sum
}

// part2 is really not that nice of a read. I'd suggest against it. Part 2 and
// the IntCode computer should get a refactor.
func part2(mem []int64) int64 {
	g := readGrid(mem)
	var r robot
	for y, row := range g {
		for x, c := range row {
			if strings.ContainsRune("^v<>", c) {
				r = robot{point{x, y}, north}
				break
			}
		}
	}

	prev := r.pos.move(west)
	path := []direction{east}
	for {
		adj := adjacent(g, r.pos)
		if len(path) > 2 && len(adj) == 1 {
			break
		}
		if 0 < len(adj) && len(adj) <= 2 {
			next := adj[0]
			if len(adj) == 2 && adj[0] == prev {
				next = adj[1]
			}
			if r.pos.x < next.x {
				r.dir = east
			}
			if r.pos.x > next.x {
				r.dir = west
			}
			if r.pos.y < next.y {
				r.dir = south
			}
			if r.pos.y > next.y {
				r.dir = north
			}
		}
		path = append(path, r.dir)
		prev = r.pos
		r.move()
	}

	queue := []step{{0, right}}
	for i, d := range path[1:] {
		rot := path[i].rotation(d)
		if len(queue) > 0 && rot == none {
			queue[len(queue)-1].n++
			continue
		}
		queue = append(queue, step{1, rot})
	}

	var out string
	for _, v := range queue {
		out += v.String() + ","
	}
	fmt.Println(out)

	p := newProgram(mem)
	p.memory[0] = 2
	// I have to be honest. I used my editors "highlight other occurrences"
	// feature on the previous output.
	in := strings.Join([]string{
		"A,B,A,C,A,B,C,C,A,B",
		"R,8,L,10,R,8",
		"R,12,R,8,L,8,L,12",
		"L,12,L,10,L,8",
		"n",
		"",
	}, "\n")
	for _, c := range in {
		p.input = append(p.input, int64(c))
	}
	res, err := p.run()
	if err != nil {
		log.Fatal(err)
	}
	return res[len(res)-1]
}

type step struct {
	n   int
	rot rotation
}

func (s step) String() string { return fmt.Sprintf("%s,%d", s.rot, s.n) }

func readGrid(mem []int64) []string {
	p := newProgram(mem)
	var raw string
	for {
		out, err := p.take(1)
		if err != nil {
			break
		}
		raw += fmt.Sprintf("%c", out[0])
	}
	return strings.Split(raw, "\n")
}

// adjacent returns a slice of adjacent positions.
func adjacent(g []string, p point) []point {
	var ps []point
	qs := []point{
		point{p.x, p.y - 1},
		point{p.x, p.y + 1},
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
	}
	for _, q := range qs {
		if 0 > q.y || q.y >= len(g) {
			continue
		}
		if 0 > q.x || q.x >= len(g[q.y]) {
			continue
		}
		if g[q.y][q.x] == '#' {
			ps = append(ps, q)
		}
	}
	return ps
}

func intersections(g []string) []point {
	var ps []point
	for y := 1; y < len(g)-2; y++ {
		for x := 1; x < len(g[y])-2; x++ {
			if g[y][x] != '#' {
				continue
			}
			p := point{x, y}
			if len(adjacent(g, p)) < 4 {
				continue
			}
			ps = append(ps, p)
		}
	}
	return ps
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
