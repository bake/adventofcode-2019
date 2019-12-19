package main

import (
	"fmt"
	"image"
	"math"
	"strings"
)

type point struct{ x, y int }

func (p point) String() string { return fmt.Sprintf("(%d,%d)", p.x, p.y) }

type tile int

const (
	space tile = iota
	beam
)

func (t tile) String() string {
	switch t {
	case space:
		return "."
	case beam:
		return "#"
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
	return strings.TrimSpace(str)
}
