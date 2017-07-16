package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAcoshKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Abs) }

func BenchmarkAcoshGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Acosh) }

func BenchmarkAcoshChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Acosh) }

func BenchmarkAcoshBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Acosh) }
