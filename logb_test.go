package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLogbKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Logb) }

func BenchmarkLogbGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Logb) }

func BenchmarkLogbChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Logb) }

func BenchmarkLogbBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Logb) }

func BenchmarkIlogbKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyInt = math32.Ilogb(x)
		}
	}
}

func BenchmarkIlogbGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyInt = math.Ilogb(x)
		}
	}
}

func BenchmarkIlogbChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyInt = chewxy.Ilogb(x)
		}
	}
}

// barnex does not implement Ilogb
/*
func BenchmarkIlogbBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyInt = barnex.Ilogb(x)
		}
	}
}
*/
