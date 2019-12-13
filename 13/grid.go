package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type point struct{ x, y int64 }

type tile int64

const (
	empty tile = iota
	wall
	block
	paddle
	ball
)

func (t tile) String() string {
	switch t {
	case empty:
		return " "
	case wall:
		return "#"
	case block:
		return "="
	case paddle:
		return "-"
	case ball:
		return "o"
	default:
		return fmt.Sprintf("<unknown %d>", t)
	}
}

type grid map[point]tile

func (g grid) ColorModel() color.Model {
	return color.RGBAModel
}

func (g grid) Bounds() image.Rectangle {
	min := point{math.MaxInt64, math.MaxInt64}
	max := point{-math.MaxInt64, -math.MaxInt64}
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

func (g grid) At(x, y int) color.Color {
	return color.Transparent
}

func (g grid) String() string {
	var str string
	b := g.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			str += fmt.Sprint(g[point{int64(x), int64(y)}])
		}
		str += fmt.Sprintln()
	}
	return str
}
