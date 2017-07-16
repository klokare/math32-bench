package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkDimKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Dim) }

func BenchmarkDimGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Dim) }

func BenchmarkDimChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Dim) }

func BenchmarkDimBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Dim) }

func BenchmarkMinKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Min) }

func BenchmarkMinGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Min) }

func BenchmarkMinChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Min) }

func BenchmarkMinBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Min) }

func BenchmarkMaxKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Max) }

func BenchmarkMaxGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Max) }

func BenchmarkMaxChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Max) }

func BenchmarkMaxBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Max) }
