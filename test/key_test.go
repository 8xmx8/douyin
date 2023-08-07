package test

import (
	"testing"
	"time"
)

/*
go test -benchmem -benchtime=3s -bench=.

cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkByteA-12               46484457                98.90 ns/op           24 B/op          1 allocs/op
BenchmarkByteB-12               44854060                83.33 ns/op           16 B/op          1 allocs/op
BenchmarkByteC-12               38626191                97.36 ns/op           16 B/op          1 allocs/op
BenchmarkBuilder-12             28206279               128.5 ns/op            88 B/op          2 allocs/op
BenchmarkPrintf-12              26774776               166.6 ns/op            56 B/op          2 allocs/op
BenchmarkBuffer-12              24778485               156.5 ns/op           136 B/op          3 allocs/op
BenchmarkByteAParallel-12       298355262               12.21 ns/op           24 B/op          1 allocs/op
BenchmarkByteBParallel-12       328606353               11.07 ns/op           16 B/op          1 allocs/op
BenchmarkByteCParallel-12       313040088               11.64 ns/op           16 B/op          1 allocs/op
BenchmarkBuilderParallel-12     180520376               19.96 ns/op           88 B/op          2 allocs/op
BenchmarkPrintfParallel-12      148170528               24.34 ns/op           56 B/op          2 allocs/op
BenchmarkBufferParallel-12      124574508               28.74 ns/op          136 B/op          3 allocs/op
PASS
*/
func BenchmarkByteA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteA(time.Now().UnixNano())
	}
}

func BenchmarkByteB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteB(time.Now().UnixNano())
	}
}

func BenchmarkByteC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteC(time.Now().UnixNano())
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder(time.Now().UnixNano())
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Printf(time.Now().UnixNano())
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buffer(time.Now().UnixNano())
	}
}

func BenchmarkByteAParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ByteA(time.Now().UnixNano())
		}
	})
}

func BenchmarkByteBParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ByteB(time.Now().UnixNano())
		}
	})
}

func BenchmarkByteCParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ByteC(time.Now().UnixNano())
		}
	})
}

func BenchmarkBuilderParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Builder(time.Now().UnixNano())
		}
	})
}

func BenchmarkPrintfParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Printf(time.Now().UnixNano())
		}
	})
}

func BenchmarkBufferParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Buffer(time.Now().UnixNano())
		}
	})
}
