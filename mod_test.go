package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkModKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Mod) }

func BenchmarkModGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Mod) }

func BenchmarkModChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Mod) }

func BenchmarkModBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Mod) }
