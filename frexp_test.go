package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkFrexpKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = math32.Frexp(x)
		}
	}
}

func BenchmarkFrexpGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy64, dummyInt = math.Frexp(x)
		}
	}
}

func BenchmarkFrexpChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = chewxy.Frexp(x)
		}
	}
}

// barnex does not implement Frepx
/*
func BenchmarkFrexpBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = barnex.Frexp(x)
		}
	}
}
*/
