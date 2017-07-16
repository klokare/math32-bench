package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkNaNKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy32 = math32.NaN()
	}
}

func BenchmarkNaNGolang64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy64 = math.NaN()
	}
}

func BenchmarkNaNChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy32 = chewxy.NaN()
	}
}

// barnex does not implement IsNaN
/*
func BenchmarkIsNaNBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
			dummy32 = barnex.NaN()
		}
	}
*/

func BenchmarkIsNaNKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = math32.IsNaN(x)
		}
	}
}

func BenchmarkIsNaNGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = math.IsNaN(x)
		}
	}
}

func BenchmarkIsNaNChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = chewxy.IsNaN(x)
		}
	}
}

// barnex does not implement IsNaN
/*
func BenchmarkIsNaNBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummyBool = barnex.IsNaN(x)
		}
	}
}
*/

func BenchmarkInfKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			dummy32 = math32.Inf(sign)
		}
	}
}

func BenchmarkInfGolang64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			dummy64 = math.Inf(sign)
		}
	}
}

func BenchmarkInfChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			dummy32 = chewxy.Inf(sign)
		}
	}
}

// barnex does not implement IsInf
/*
func BenchmarkIsInfBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			dummy32 = barnex.Inf(sign)
		}
	}
}
*/

func BenchmarkIsInfKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			for _, x := range values {
				dummyBool = math32.IsInf(x, sign)
			}
		}
	}
}

func BenchmarkIsInfGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for sign := -1; sign <= 1; sign++ {
			for _, x := range values {
				dummyBool = math.IsInf(x, sign)
			}
		}
	}
}

func BenchmarkIsInfChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			for sign := -1; sign <= 1; sign++ {
				dummyBool = chewxy.IsInf(x, sign)
			}
		}
	}
}

// barnex does not implement IsInf
/*
func BenchmarkIsInfBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			for sign := -1; sign <= 1; sign++ {
				dummyBool = barnex.IsInf(x, sign)
			}
		}
	}
}
**/
