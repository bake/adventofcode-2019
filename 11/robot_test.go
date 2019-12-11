package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	r := &robot{dir: up}
	tt := []struct{ prev, next direction }{
		{left, left},
		{left, down},
		{left, right},
		{right, down},
	}
	for _, tc := range tt {
		r.rotate(tc.prev)
		if r.dir != tc.next {
			t.Fatalf("expected robot to look %s, instead it looks %s", tc.next, r.dir)
		}
	}
}

func TestStep(t *testing.T) {
	r := &robot{dir: up}
	tt := []robot{
		{point{0, -1}, left},
		{point{-1, -1}, right},
		{point{-1, -2}, right},
		{point{0, -2}, right},
	}
	for _, tc := range tt {
		r.step()
		r.rotate(tc.dir)
		if r.point != tc.point {
			t.Fatalf("expected robot to stand at %v, instead it stands at %v", tc.point, r.point)
		}
	}
}
