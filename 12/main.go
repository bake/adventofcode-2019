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

func (p point) zero() bool {
	return p.x == 0 && p.y == 0 && p.z == 0
}

type points []point

func (ps points) filter(fn func(point) bool) bool {
	for _, p := range ps {
		if !fn(p) {
			return false
		}
	}
	return true
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
	// fmt.Println(part1(moons, 1000))
	fmt.Println(part2(moons))
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

func part2(moons []*object) int {
	loop := make([]int, 3)
	for i := 1; ; i++ {
		for j, m := range moons {
			m.gravity(moons, j)
		}
		for _, m := range moons {
			m.pos.add(m.vel)
		}

		// I hate this. Someone should refactor it.
		if loop[0] == 0 && moons[0].vel.x == 0 && moons[1].vel.x == 0 && moons[2].vel.x == 0 && moons[3].vel.x == 0 {
			loop[0] = i * 2
		}
		if loop[1] == 0 && moons[0].vel.y == 0 && moons[1].vel.y == 0 && moons[2].vel.y == 0 && moons[3].vel.y == 0 {
			loop[1] = i * 2
		}
		if loop[2] == 0 && moons[0].vel.z == 0 && moons[1].vel.z == 0 && moons[2].vel.z == 0 && moons[3].vel.z == 0 {
			loop[2] = i * 2
		}

		done := true
		for _, v := range loop {
			if v == 0 {
				done = false
				break
			}
		}
		if done {
			break
		}
	}

	return lcm(loop[0], loop[1], loop[2])
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
