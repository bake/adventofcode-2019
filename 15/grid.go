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
	case 1:
		q.y--
	case 2:
		q.y++
	case 3:
		q.x--
	case 4:
		q.x++
	}
	return q
}

type direction int

const (
	_ direction = iota
	north
	south
	west
	east
)

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case south:
		return "south"
	case west:
		return "west"
	case east:
		return "east"
	default:
		return fmt.Sprintf("<unknown %d>", d)
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
