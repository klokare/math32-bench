package bench

import (
	"math"
	"testing"

	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkJnKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = math32.Jn(n, x)
			}
		}
	}
}

func BenchmarkJnGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy64 = math.Jn(n, x)
			}
		}
	}
}

func BenchmarkJnChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = chewxy.Jn(n, x)
			}
		}
	}
}

// barnex doese not implement Jn
/*
func BenchmarkJnBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = barnex.Jn(n, x)
			}
		}
	}
}
*/

func BenchmarkYnKlokare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = math32.Yn(n, x)
			}
		}
	}
}

func BenchmarkYnGolang64(b *testing.B) {
	values := values64()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy64 = math.Yn(n, x)
			}
		}
	}
}

func BenchmarkYnChewxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = chewxy.Yn(n, x)
			}
		}
	}
}

// barnex doese not implement Yn
/*
func BenchmarkYnBarnex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := -5; n <= 5; n++ {
			for _, x := range values {
				dummy32 = barnex.Yn(n, x)
			}
		}
	}
}
*/
