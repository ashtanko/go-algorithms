package main

import (
	"strings"
	"testing"
)

func Benchmark_run_line_short(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_line_short(data, false)
	}
}

func Benchmark_run_line_long(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_line_long(data, false)
	}
}
func Benchmark_run_line_short_workers(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_line_short_workers(data, false)
	}
}
func Benchmark_run_line_long_workers(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_line_long_workers(data, false)
	}
}
func Benchmark_run_bulk_short(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_bulk_short(data, false)
	}
}
func Benchmark_run_bulk_long(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_bulk_long(data, false)
	}
}
func Benchmark_run_seq_short(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_seq_short(data, false)
	}
}
func Benchmark_run_seq_long(b *testing.B) {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	for i := 0; i < b.N; i++ {
		run_seq_long(data, false)
	}
}
