package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type point struct{ x, y, z int }

func (p point) String() string {
	return fmt.Sprintf("<x=%3d, y=%3d, z=%3d>", p.x, p.y, p.z)
}

func (p point) compare(q point) point {
	var r point
	switch {
	case p.x < q.x:
		r.x = 1
	case p.x > q.x:
		r.x = -1
	}
	switch {
	case p.y < q.y:
		r.y = 1
	case p.y > q.y:
		r.y = -1
	}
	switch {
	case p.z < q.z:
		r.z = 1
	case p.z > q.z:
		r.z = -1
	}
	return r
}

func (p *point) add(q point) {
	p.x += q.x
	p.y += q.y
	p.z += q.z
}

func (p point) abs() int {
	var e int
	e += int(math.Abs(float64(p.x)))
	e += int(math.Abs(float64(p.y)))
	e += int(math.Abs(float64(p.z)))
	return e
}

type object struct{ pos, vel point }

func (o object) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", o.pos, o.vel)
}

func (o *object) gravity(os []*object, i int) {
	for j, p := range os {
		if i == j {
			continue
		}
		o.vel.add(o.pos.compare(p.pos))
	}
}

func main() {
	moons, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(moons, 1000))
}

func part1(moons []*object, n int) int {
	for i := 0; i < n; i++ {
		for j, m := range moons {
			m.gravity(moons, j)
		}
		for _, m := range moons {
			m.pos.add(m.vel)
		}
	}

	var e int
	for _, m := range moons {
		e += m.pos.abs() * m.vel.abs()
	}
	return e
}

func parseInput(r io.Reader) (objs []*object, err error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		var o object
		fmt.Sscanf(s.Text(), "<x=%d, y=%d, z=%d>", &o.pos.x, &o.pos.y, &o.pos.z)
		objs = append(objs, &o)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return objs, nil
}
