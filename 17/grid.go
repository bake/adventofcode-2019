package main

import (
	"fmt"
	"image"
	"math"
)

type point struct{ x, y int }

func (p point) move(d direction) point {
	q := p
	switch d {
	case north:
		q.y--
	case south:
		q.y++
	case west:
		q.x--
	case east:
		q.x++
	}
	return q
}

func (p point) String() string { return fmt.Sprintf("(%d, %d)", p.x, p.y) }

type direction int

const (
	_ direction = iota
	north
	south
	west
	east
)

func (d direction) rotation(e direction) rotation {
	if d == e {
		return none
	}
	if d == north && e == west {
		return left
	}
	if d == south && e == east {
		return left
	}
	if d == west && e == south {
		return left
	}
	if d == east && e == north {
		return left
	}
	return right
}

func (d direction) String() string {
	switch d {
	case north:
		return "^"
	case south:
		return "v"
	case west:
		return "<"
	case east:
		return ">"
	default:
		return fmt.Sprintf("<unknown %d>", d)
	}
}

type robot struct {
	pos point
	dir direction
}

func (r *robot) move() { r.pos = r.pos.move(r.dir) }

func (r robot) String() string { return fmt.Sprintf("%s -> %s", r.pos, r.dir) }

type rotation int

const (
	none rotation = iota
	left
	right
)

func (r rotation) String() string {
	switch r {
	case left:
		return "L"
	case right:
		return "R"
	default:
		return fmt.Sprintf("<unknown %d>", r)
	}
}

type tile int

const (
	unknown tile = iota
	wall
	free
	target
)

// String returns a string representation of the tile. The characters for free
// and tile *in the unknown* are switched to increase readability.
func (t tile) String() string {
	switch t {
	case unknown:
		return "."
	case wall:
		return "#"
	case free:
		return " "
	case target:
		return "x"
	default:
		return fmt.Sprintf("<unknown %d>", t)
	}
}

type grid map[point]tile

func (g grid) Bounds() image.Rectangle {
	min := point{math.MaxInt64, math.MaxInt64}
	max := point{math.MinInt64, math.MinInt64}
	for p := range g {
		if p.x < min.x {
			min.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}
	}
	return image.Rect(int(min.x), int(min.y), int(max.x)+1, int(max.y)+1)
}

func (g grid) String() string {
	var str string
	b := g.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			str += fmt.Sprint(g[point{int(x), int(y)}])
		}
		str += fmt.Sprintln()
	}
	return str
}
