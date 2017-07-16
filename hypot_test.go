package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkHypotKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Hypot) }

func BenchmarkHypotGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Hypot) }

func BenchmarkHypotChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Hypot) }

func BenchmarkHypotBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Hypot) }
