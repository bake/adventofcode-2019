package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
)

const (
	asteroid byte = '#'
	space    byte = '.'
)

type point struct{ x, y int }

func (p point) distance(q point) float64 {
	return math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y))
}

type asteroids struct {
	center    point
	asteroids []point
}

func (a *asteroids) add(p point) { a.asteroids = append(a.asteroids, p) }

func (a asteroids) Len() int { return len(a.asteroids) }
func (a asteroids) Less(i, j int) bool {
	return a.center.distance(a.asteroids[i]) < a.center.distance(a.asteroids[j])
}
func (a asteroids) Swap(i, j int) {
	a.asteroids[i], a.asteroids[j] = a.asteroids[j], a.asteroids[i]
}

func main() {
	grid, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := part1(grid)
	fmt.Println(n)
	fmt.Println(part2(grid, 200))
}

func part1(grid [][]byte) (max int, pt point) {
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != asteroid {
				continue
			}
			n := len(asteroidsInView(grid, point{x, y}))
			if n > max {
				max = n
				pt = point{x, y}
			}
		}
	}
	return max, pt
}

// part2 is solved in an awful way. I'm sorry.
func part2(grid [][]byte, n int) int {
	_, pt := part1(grid)
	grid[pt.y][pt.x] = 'x'
	inView := asteroidsInView(grid, pt)

	var j int
	for {
		var angles []float64
		for a, asteroids := range inView {
			if len(asteroids.asteroids) == 0 {
				continue
			}
			angles = append(angles, a)
			sort.Sort(asteroids)
		}
		if len(angles) == 0 {
			break
		}
		sort.Float64s(angles)

		var start int
		for i, a := range angles {
			if a >= -90 {
				start = i
				break
			}
		}

		for i := start; i != start+1; i-- {
			j++
			a := angles[i]
			pt := inView[a].asteroids[0]
			inView[a].asteroids = inView[a].asteroids[1:]
			grid[pt.y][pt.x] = '_'

			if i == 0 {
				i = len(angles) - 1
			}

			if j == 200-1 {
				return pt.x*100 + pt.y
			}
		}
	}
	return 0
}

func print(grid [][]byte) {
	for _, row := range grid {
		fmt.Printf("%s\n", row)
	}
}

func numAsteroidsInView(grid [][]byte, x, y int) (n int) {
	angles := map[float64]bool{}
	for y2, row := range grid {
		for x2, cell := range row {
			a := math.Atan2(float64(y2-y), float64(x2-x))
			if x == x2 && y == y2 {
				continue
			}
			if cell != asteroid {
				continue
			}
			angles[a] = true
		}
	}
	return len(angles)
}

func asteroidsInView(grid [][]byte, p point) map[float64]*asteroids {
	angles := map[float64]*asteroids{}
	for y, row := range grid {
		for x, cell := range row {
			if cell != asteroid {
				continue
			}
			q := point{x, y}
			if p == q {
				continue
			}
			x, y := float64(p.x-q.x), float64(q.y-p.y)
			a := math.Atan2(y, x) * 180 / math.Pi
			if _, ok := angles[a]; !ok {
				angles[a] = &asteroids{center: p}
			}
			angles[a].add(q)
		}
	}
	return angles
}

func parseInput(r io.Reader) ([][]byte, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	body = bytes.TrimSpace(body)
	return bytes.Split(body, []byte{'\n'}), nil
}
