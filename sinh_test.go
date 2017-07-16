package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkSinhKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Sinh) }

func BenchmarkSinhGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Sinh) }

func BenchmarkSinhChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Sinh) }

func BenchmarkSinhBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Sinh) }

func BenchmarkCoshKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Cosh) }

func BenchmarkCoshGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Cosh) }

func BenchmarkCoshChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Cosh) }

func BenchmarkCoshBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Cosh) }
