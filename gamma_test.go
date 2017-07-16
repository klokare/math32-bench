package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkGammaKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Gamma) }

func BenchmarkGammaGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Gamma) }

func BenchmarkGammaChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Gamma) }

func BenchmarkGammaBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Gamma) }
