package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkSinKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Sin) }

func BenchmarkSinGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Sin) }

func BenchmarkSinChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Sin) }

func BenchmarkSinBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Sin) }

func BenchmarkCosKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Cos) }

func BenchmarkCosGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Cos) }

func BenchmarkCosChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Cos) }

func BenchmarkCosBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Cos) }
