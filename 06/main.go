package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	m, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(m))
	fmt.Println(part2(m, "YOU", "SAN"))
}

func part1(m map[string]string) int {
	var d int
	for n := range m {
		d += depth(m, n)
	}
	return d
}

func part2(m map[string]string, a, b string) int {
	a, b = m[a], m[b]
	min := len(m)
	for n := range m {
		da, err := distance(m, n, a)
		if err != nil {
			continue
		}
		db, err := distance(m, n, b)
		if err != nil {
			continue
		}
		if min > da+db {
			min = da + db
		}
	}
	return min
}

// depth calculates the depth of a given key to another key b or the trees root.
func depth(m map[string]string, a string) (d int) {
	for {
		o, ok := m[a]
		if !ok {
			break
		}
		d++
		a = o
	}
	return d
}

// distance uses bruteforce to find the closest path between two objects.
func distance(m map[string]string, b, a string) (d int, err error) {
	for {
		o, ok := m[a]
		if !ok {
			return 0, fmt.Errorf("%s and %s are not connected", a, b)
		}
		d++
		if o == b {
			break
		}
		a = o
	}
	return d, nil
}

func parseInput(r io.Reader) (map[string]string, error) {
	m := map[string]string{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		ab := strings.Split(s.Text(), ")")
		m[ab[1]] = ab[0]
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}
