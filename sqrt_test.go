package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkSqrtKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Sqrt) }

func BenchmarkSqrtGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Sqrt) }

func BenchmarkSqrtChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Sqrt) }

func BenchmarkSqrtBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Sqrt) }
