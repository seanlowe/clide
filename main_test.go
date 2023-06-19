package main

import (
	"testing"

	clide "github.com/TeddyRandby/clide/app"
)

func BenchmarkMainFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkLeaves(b *testing.B) {
	params := make(map[string]string)
	c := clide.New(params)

	for i := 0; i < b.N; i++ {
		getLeaves(c)
	}
}
