package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLdexpKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = math32.Ldexp(x, n)
			}
		}
	}
}

func BenchmarkLdexpGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy64 = math.Ldexp(x, n)
			}
		}
	}
}

func BenchmarkLdexpChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = chewxy.Ldexp(x, n)
			}
		}
	}
}

// barnex does not implement Ldexp
/*
func BenchmarkLdexpBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = barnex.Ldexp(x, n)
			}
		}
	}
}
*/
