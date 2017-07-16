package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkPow10Klokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, e := range []int{-325, -324, -100, -5, -2, -1, -0, 0, 1, 2, 5, 100, 309, 310} {
			dummy32 = math32.Pow10(e)
		}
	}
}

func BenchmarkPow10Golang64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, e := range []int{-325, -324, -100, -5, -2, -1, -0, 0, 1, 2, 5, 100, 309, 310} {
			dummy64 = math.Pow10(e)
		}
	}
}

func BenchmarkPow10Chewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, e := range []int{-325, -324, -100, -5, -2, -1, -0, 0, 1, 2, 5, 100, 309, 310} {
			dummy32 = chewxy.Pow10(e)
		}
	}
}

// barnex does not implement pow10
/*
func BenchmarkPow10Barnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, e := range []int{-325, -324, -100, -5, -2, -1, -0, 0, 1, 2, 5, 100, 309, 310} {
				dummy32 = barnex.Pow10(e)
			}
		}
}
*/
