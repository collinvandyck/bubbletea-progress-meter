package main

import "testing"

func TestProgram(t *testing.T) {
}

func BenchmarkProgram(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
