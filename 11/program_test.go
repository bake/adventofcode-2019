package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	tt := []struct {
		memory        []int64
		input, output []int64
	}{
		{
			memory: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			output: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			memory: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			output: []int64{1219070632396864},
		},
		{
			memory: []int64{104, 1125899906842624, 99},
			output: []int64{1125899906842624},
		},
	}
	for _, tc := range tt {
		p := newProgram(tc.memory)
		out, err := p.run(tc.input...)
		if err != nil {
			t.Fatal(err)
		}
		if len(tc.output) != len(out) {
			t.Fatalf("expected output to have %d elements, got %d", len(tc.output), len(out))
		}
		for i := range tc.output {
			if tc.output[i] != out[i] {
				t.Fatalf("expected output at position %d to be %d, got %d", i, tc.output[i], out[i])
			}
		}
	}
}
