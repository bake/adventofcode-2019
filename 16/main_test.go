package main

import (
	"fmt"
	"testing"
)

func TestFFT(t *testing.T) {
	tt := []struct {
		input   string
		outputs []string
	}{
		{
			input:   "12345678",
			outputs: []string{"48226158", "34040438", "03415518", "01029498"},
		},
	}
	for _, tc := range tt {
		out := tc.input
		for i, signal := range tc.outputs {
			out = fft(out)
			fmt.Println(out)
			if out != signal {
				t.Fatalf("expected output of phase %d to be %s, got %s", i+1, signal, out)
			}
		}
	}
}

func BenchmarkFTT(b *testing.B) {
	out := "12345678"
	for n := 0; n < b.N; n++ {
		out = fft(out)
	}
}
