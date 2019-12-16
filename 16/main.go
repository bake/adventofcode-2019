package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	signal := strings.TrimSpace(string(bs))
	fmt.Println(part1(signal))
	fmt.Println(part2(signal))
}

func part1(s string) string {
	for i := 0; i < 100; i++ {
		s = fft(s)
	}
	return s[:8]
}

func part2(s string) string {
	s = strings.Repeat(s, 10_000)
	off, _ := strconv.Atoi(s[:7])

	ds := make([]int, off)
	for i := off; i < len(s); i++ {
		ds[i-off], _ = strconv.Atoi(string(s[i%len(s)]))
	}

	for i := 0; i < 100; i++ {
		for j := len(ds) - 2; j >= 0; j-- {
			ds[j] += ds[j+1]
			ds[j] %= 10
		}
	}

	var out string
	for _, d := range ds[:8] {
		out += strconv.Itoa(d)
	}
	return out
}

func fft(signal string) string {
	p := []int{0, 1, 0, -1}
	var out string
	for i := range signal {
		var sum int
		// Start from i since all phase multiplicators before it are always 0.
		for j := i; j < len(signal); j++ {
			d, _ := strconv.Atoi(string(signal[j]))
			sum += d * p[(j+1)/(i+1)%len(p)]
		}
		if sum < 0 {
			sum *= -1
		}
		out += strconv.Itoa(sum % 10)
	}
	return out
}
