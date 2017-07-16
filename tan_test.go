package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkTanKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Tan) }

func BenchmarkTanGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Tan) }

func BenchmarkTanChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Tan) }

func BenchmarkTanBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Tan) }
