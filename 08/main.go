package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type layer []color.Color

type spaceImageFormat struct {
	width, height int
	data          []color.Color
}

func (img spaceImageFormat) layer(n int) layer {
	size := img.width * img.height
	return img.data[n*size : n*size+size]
}

func (img spaceImageFormat) layers() int {
	return len(img.data) / (img.width * img.height)
}

func (img spaceImageFormat) at(x, y int) color.Color {
	for i := 0; i < img.layers(); i++ {
		l := img.layer(i)
		j := y*img.width + x
		switch l[j] {
		case color.Black, color.White:
			return l[j]
		}
	}
	return color.Transparent
}

func (img spaceImageFormat) decode() image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.width, img.height))
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			m.Set(x, y, img.at(x, y))
		}
	}
	return m
}

func main() {
	input, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	img := spaceImageFormat{width: 25, height: 6, data: input}

	fmt.Fprintln(os.Stderr, part1(img))

	if err := png.Encode(os.Stdout, img.decode()); err != nil {
		log.Fatal(err)
	}
}

func part1(img spaceImageFormat) int {
	var out int
	min := len(img.layer(0))
	for i := 0; i < img.layers(); i++ {
		l := img.layer(i)
		if numOfColor(l, color.Black) < min {
			min = numOfColor(l, color.Black)
			out = numOfColor(l, color.White) * numOfColor(l, color.Transparent)
		}
	}
	return out
}

func numOfColor(data []color.Color, v color.Color) (n int) {
	for _, w := range data {
		if v == w {
			n++
		}
	}
	return n
}

func parseInput(r io.Reader) ([]color.Color, error) {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	in = bytes.TrimSpace(in)
	out := make([]color.Color, len(in))
	for i := range in {
		d, err := strconv.Atoi(string(in[i]))
		if err != nil {
			return nil, err
		}
		switch d {
		case 0:
			out[i] = color.Black
		case 1:
			out[i] = color.White
		case 2:
			out[i] = color.Transparent
		}
	}
	return out, nil
}
