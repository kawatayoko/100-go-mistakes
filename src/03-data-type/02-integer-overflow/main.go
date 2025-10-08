package main

import (
	"fmt"
	"math"
)

func main() {
	count := Inc32(math.MaxInt32)
	fmt.Printf("count: %d\n", count)
}

// 1を加算した時に整数のオーバーフローを検出する
func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("integer overflow")
	}
	return counter + 1
}

// 1を加算した時に整数のオーバーフローを検出する(int)
func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("integer overflow")
	}
	return counter + 1
}

// 1を加算した時に整数のオーバーフローを検出する(uint)
func IncUint(counter uint) uint {
	if counter == math.MaxUint {
		panic("integer overflow")
	}
	return counter + 1
}

// 加算時の整数オーバーフローの検出
func AddInt(a, b int) int {
	if b > 0 && a > (math.MaxInt-b) || b < 0 && a < (math.MinInt-b) {
		panic("int overflow")
	}
	return a + b
}

// 乗算時の整数オーバーフロー
func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}

	if a == math.MinInt || b == math.MinInt {
		panic("integer overflow")
	}

	if result/b != a { // 割った結果が乗算前と等しいかでオーバーフローをチェック
		panic("integer overflow")
	}

	return result
}
