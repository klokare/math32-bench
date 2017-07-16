package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkCopysignKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Copysign) }

func BenchmarkCopysignGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Copysign) }

func BenchmarkCopysignChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Copysign) }

func BenchmarkCopysignBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Copysign) }
