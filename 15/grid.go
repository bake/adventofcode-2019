package main

import (
	"fmt"
	"image"
	"math"
)

type point struct{ x, y int64 }

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

type direction int64

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

type tile int64

const (
	unknown tile = iota
	wall
	free
	target
)

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
