package main

import "fmt"

type point struct{ x, y int64 }

type robot struct {
	point
	dir direction
}

func (r *robot) rotate(d direction) {
	switch d {
	case left:
		r.dir = (r.dir - 1) % 4
	case right:
		r.dir = (r.dir + 1) % 4
	}
	if r.dir < 0 {
		r.dir = 3
	}
}

func (r *robot) step() {
	dx := map[direction]int64{left: -1, right: 1, up: 0, down: 0}
	dy := map[direction]int64{left: 0, right: 0, up: -1, down: 1}
	r.x += dx[r.dir]
	r.y += dy[r.dir]
}

type direction int64

const (
	up direction = iota
	right
	down
	left
)

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case right:
		return "right"
	case left:
		return "left"
	case down:
		return "down"
	default:
		return fmt.Sprintf("<unknown %d>", d)
	}
}
