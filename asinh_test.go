package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAsinhKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Asinh) }

func BenchmarkAsinhGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Asinh) }

func BenchmarkAsinhChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Asinh) }

func BenchmarkAsinhBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Asinh) }
