package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

const (
	asteroid byte = '#'
	space    byte = '.'
)

func main() {
	grid, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// for _, row := range grid {
	// 	fmt.Printf("%s\n", row)
	// }

	fmt.Println(part1(grid))
}

func part1(grid [][]byte) (max int) {
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != asteroid {
				continue
			}
			n := asteroids(grid, x, y)
			if n > max {
				max = n
			}
		}
	}
	return max
}

func asteroids(grid [][]byte, x, y int) (n int) {
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

func parseInput(r io.Reader) ([][]byte, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	body = bytes.TrimSpace(body)
	return bytes.Split(body, []byte{'\n'}), nil
}
