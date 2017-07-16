package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAtan2Klokare(b *testing.B) { benchTwoFloatToOneFloat32(b, math32.Atan2) }

func BenchmarkAtan2Golang64(b *testing.B) { benchTwoFloatToOneFloat64(b, math.Atan2) }

func BenchmarkAtan2Chewxy(b *testing.B) { benchTwoFloatToOneFloat32(b, chewxy.Atan2) }

func BenchmarkAtan2Barnex(b *testing.B) { benchTwoFloatToOneFloat32(b, barnex.Atan2) }
