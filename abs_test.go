package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAbsKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Abs) }

func BenchmarkAbsGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Abs) }

func BenchmarkAbsChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Abs) }

func BenchmarkAbsBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Abs) }
