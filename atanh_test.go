package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAtanhKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Atanh) }

func BenchmarkAtanhGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Atanh) }

func BenchmarkAtanhChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Atanh) }

func BenchmarkAtanhBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Atanh) }
