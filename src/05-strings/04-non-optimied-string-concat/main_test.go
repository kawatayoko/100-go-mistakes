package main

import (
	"strings"
	"testing"
)

var (
	thousandBytes = []string{strings.Repeat("a", 1000)}
	thousandChars = func() []string {
		ss := make([]string, 1000)
		for i := range ss {
			ss[i] = "a"
		}
		return ss
	}()
)

// ベンチマーク実行
// go test -bench=. -benchmem
// 残：入力のスライスには、1000個の文字列、各文字列には1000バイト
func BenchmarkConcat1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat1(thousandBytes)
	}
}

func BenchmarkConcat2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat2(thousandBytes)
	}
}

func BenchmarkConcat3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat3(thousandBytes)
	}
}

func BenchmarkConcat1ManySmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat1(thousandChars)
	}
}

func BenchmarkConcat2ManySmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat2(thousandChars)
	}
}

func BenchmarkConcat3ManySmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat3(thousandChars)
	}
}
