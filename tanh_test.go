package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkTanhKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Tanh) }

func BenchmarkTanhGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Tanh) }

func BenchmarkTanhChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Tanh) }

func BenchmarkTanhBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Tanh) }
