package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	rs, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(rs, 1))
	fmt.Println(part2(rs, 1_000_000_000_000))
}

func part1(rs reactions, fuel int64) int64 {
	return rs.reduce(fuel)
}

func part2(rs reactions, ore int64) int64 {
	return search(rs, ore, 0, math.MaxInt32)
}

func search(rs reactions, ore, min, max int64) int64 {
	if min >= max-1 {
		return min
	}
	m := min/2 + max/2
	if rs.reduce(m) > ore {
		return search(rs, ore, min, m)
	}
	return search(rs, ore, m, max)
}

type reaction struct {
	quantity  int64
	chemicals map[string]int64
}

type reactions map[string]reaction

func (rs reactions) reduce(fuel int64) int64 {
	bank := map[string]int64{}
	list := map[string]int64{"FUEL": fuel}
	for len(list) > 1 || list["FUEL"] != 0 {
		for name, q := range list {
			if name == "ORE" {
				continue
			}
			r := rs[name]
			n := int64(math.Ceil(float64(q-bank[name]) / float64(r.quantity)))
			for name, q := range r.chemicals {
				list[name] += n * q
			}
			bank[name] += n*r.quantity - q
			delete(list, name)
		}
	}
	return list["ORE"]
}

func parseInput(r io.Reader) (reactions, error) {
	regex := regexp.MustCompile(`([0-9]+) ([A-Z]+)`)
	s := bufio.NewScanner(r)
	rs := reactions{}
	for s.Scan() {
		sides := strings.Split(s.Text(), " => ")
		var name string
		var quantity int64
		if _, err := fmt.Sscanf(sides[1], "%d %s", &quantity, &name); err != nil {
			return nil, err
		}
		pairs := regex.FindAllString(sides[0], -1)
		chemicals := map[string]int64{}
		for _, p := range pairs {
			var n string
			var q int64
			if _, err := fmt.Sscanf(p, "%d %s", &q, &n); err != nil {
				return nil, err
			}
			chemicals[n] += q
		}
		rs[name] = reaction{quantity, chemicals}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}
