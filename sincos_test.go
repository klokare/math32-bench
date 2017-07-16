package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkSincosKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = math32.Sincos(x)
		}
	}
}

func BenchmarkSincosGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy64, dummy64 = math.Sincos(x)
		}
	}
}

func BenchmarkSincosChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = chewxy.Sincos(x)
		}
	}
}

// barnex does not implmenet sincos
/*
func BenchmarkSincosBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = barnex.Sincos(x)
		}
	}
}
*/
