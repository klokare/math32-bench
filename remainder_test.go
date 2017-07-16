package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkRemainderKlokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Remainder) }

func BenchmarkRemainderGolang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Remainder) }

func BenchmarkRemainderChewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Remainder) }

func BenchmarkRemainderBarnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Remainder) }
