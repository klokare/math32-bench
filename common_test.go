package bench

import (
	"testing"

	"github.com/klokare/math32"
)

var dummy32 float32
var dummy64 float64
var dummyBool bool
var dummyInt int

// Test values
var values = []float32{
	// Not a number
	math32.NaN(),
	// ±Infinity
	math32.Inf(-1),
	math32.Inf(1),
	// Really big
	math32.MaxFloat32,
	-math32.MaxFloat32,
	// Really small but non-zero
	math32.SmallestNonzeroFloat32,
	-math32.SmallestNonzeroFloat32,
	// Numbers bigger than 1
	5, 2, 1,
	-5, -2, -1,
	// Numbers smaller than 1
	0.9, 0.5, 0.2, 0.1,
	-0.9, -0.5, -0.2, -0.1,
	// Zero
	0, -0,
}

func values64() []float64 {
	v64 := make([]float64, len(values))
	for i, x := range values {
		v64[i] = float64(x)
	}
	return v64
}

func benchOneFloatToOneFloat32(b *testing.B, fx func(float32) float32) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy32 = fx(x)
		}
	}
}

func benchOneFloatToOneFloat64(b *testing.B, fx func(float64) float64) {
	values := values64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			dummy64 = fx(x)
		}
	}
}

func benchTwoFloatToOneFloat32(b *testing.B, fx func(float32, float32) float32) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			for _, y := range values {
				dummy32 = fx(x, y)
			}
		}
	}
}

func benchTwoFloatToOneFloat64(b *testing.B, fx func(float64, float64) float64) {
	values := values64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range values {
			for _, y := range values {
				dummy64 = fx(x, y)
			}
		}
	}
}
