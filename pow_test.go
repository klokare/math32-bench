package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkPowKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Pow) }

func BenchmarkPowGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Pow) }

func BenchmarkPowChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Pow) }

func BenchmarkPowBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Pow) }
