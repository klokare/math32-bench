package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkSignbitKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = math32.Signbit(x)
		}
	}
}

func BenchmarkSignbitGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = math.Signbit(x)
		}
	}
}

func BenchmarkSignbitChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = chewxy.Signbit(x)
		}
	}
}

// barnex does not implement Signbit
/*
func BenchmarkSignbitBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = barnex.Signbit(x)
		}
	}
}
*/
