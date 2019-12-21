package main

import "fmt"

type point struct{ x, y int }

func (p point) String() string { return fmt.Sprintf("(%d,%d)", p.x, p.y) }

type grid []string

func (g grid) label(name string) point {
	for y, row := range g {
		for x := range row {
			if g.at(point{x, y}) == name {
				return point{x, y}
			}
		}
	}
	return point{}
}

func (g grid) labels(dst, src string) (map[point]point, point, point) {
	ps := map[point]point{}
	ls := map[string]point{}
	var a, b point
	for y, row := range g {
		for x := range row {
			if !isLetter(g[y][x]) {
				continue
			}
			p := point{x, y}
			for _, q := range g.adjacent(p) {
				if g[q.y][q.x] != '.' {
					continue
				}
				l := g.at(p)
				if l == dst {
					b = q
				}
				if l == src {
					a = q
				}
				if d, ok := ls[l]; ok {
					ps[d] = q
					ps[q] = d
				}
				ls[l] = q
				break
			}
		}
	}
	return ps, b, a
}

func isLetter(c byte) bool { return 'A' <= c && c <= 'Z' }

func (g grid) at(p point) string {
	if !isLetter(g[p.y][p.x]) {
		return string(g[p.y][p.x])
	}

	for _, q := range g.adjacent(p) {
		if !isLetter(g[q.y][q.x]) {
			continue
		}
		if q.x < p.x || q.y < p.y {
			return fmt.Sprintf("%c%c", g[q.y][q.x], g[p.y][p.x])
		}
		return fmt.Sprintf("%c%c", g[p.y][p.x], g[q.y][q.x])
	}

	return ""
}

func (g grid) adjacent(p point) []point {
	var ps []point
	qs := []point{
		point{p.x, p.y - 1},
		point{p.x, p.y + 1},
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
	}
	for _, q := range qs {
		if 0 > q.y || q.y >= len(g) {
			continue
		}
		if 0 > q.x || q.x >= len(g[q.y]) {
			continue
		}
		c := g[q.y][q.x]
		if '.' == c {
			ps = append(ps, q)
		}
		if isLetter(c) {
			ps = append(ps, q)
		}
	}
	return ps
}

func (g grid) String() string {
	var str string
	for _, row := range g {
		str += fmt.Sprintf("%s\n", row)
	}
	return str[:len(str)-1]
}
