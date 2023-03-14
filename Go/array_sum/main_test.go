package main

import "testing"

func BenchmarkSumForSmall(b *testing.B) {
	arr := make([]int, 10)
	for i := 0; i < b.N; i++ {
		sumFor(arr)
	}
}

func BenchmarkSumReduceSmall(b *testing.B) {
	arr := make([]int, 10)
	for i := 0; i < b.N; i++ {
		sumReduce(arr)
	}
}

func BenchmarkSumForMedium(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		sumFor(arr)
	}
}

func BenchmarkSumReduceMedium(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		sumReduce(arr)
	}
}

func BenchmarkSumForLarge(b *testing.B) {
	arr := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		sumFor(arr)
	}
}

func BenchmarkSumReduceLarge(b *testing.B) {
	arr := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		sumReduce(arr)
	}
}
