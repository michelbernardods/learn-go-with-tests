package main

import "testing"

/*
Run this is test with:
	Linux(go test -bench=. o)
	Windowns(go test-bench=".")
or go test -v
*/

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
