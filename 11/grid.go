package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type spaceColor int64

const (
	black spaceColor = iota
	white
)

func (c spaceColor) String() string {
	switch c {
	case black:
		return "black"
	case white:
		return "white"
	default:
		return fmt.Sprintf("<unknown %d>", c)
	}
}

type grid map[point]spaceColor

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
	switch g[point{int64(x), int64(y)}] {
	case black:
		return color.Black
	case white:
		return color.White
	default:
		return color.Transparent
	}
}
