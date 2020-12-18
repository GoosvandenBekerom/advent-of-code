package main

import (
	"testing"
)

func Benchmark_parse1(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		{name: "1 + 2 * 3 + 4 * 5 + 6"},
		{name: "2 * 3 + (4 * 5)"},
		{name: "5 + (8 * 3 + 9 + 3 * 4 * 3)"},
		{name: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"},
		{name: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parse1(bm.name)
			}
		})
	}
}
