package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAtanKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Atan) }

func BenchmarkAtanGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Atan) }

func BenchmarkAtanChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Atan) }

func BenchmarkAtanBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Atan) }
