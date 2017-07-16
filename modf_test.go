package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkModfKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = math32.Modf(x)
		}
	}
}

func BenchmarkModfGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy64, dummy64 = math.Modf(x)
		}
	}
}

func BenchmarkModfChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = chewxy.Modf(x)
		}
	}
}

// barnex does not implement modf
/*
func BenchmarkModfBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32, dummy32 = barnex.Modf(x)
		}
	}
}
*/
