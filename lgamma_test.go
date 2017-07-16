package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLgammaKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = math32.Lgamma(x)
		}
	}
}

func BenchmarkLgammaGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy64, dummyInt = math.Lgamma(x)
		}
	}
}

func BenchmarkLgammaChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = chewxy.Lgamma(x)
		}
	}
}

// barnex does not implement Lgamma
/*
func BenchmarkLgammaBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummyInt = barnex.Lgamma(x)
		}
	}
}
*/
