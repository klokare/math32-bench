package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkCbrtKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Cbrt) }

func BenchmarkCbrtGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Cbrt) }

func BenchmarkCbrtChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Cbrt) }

func BenchmarkCbrtBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Cbrt) }
