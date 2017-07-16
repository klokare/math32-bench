package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkExpm1Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Expm1) }

func BenchmarkExpm1Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Expm1) }

func BenchmarkExpm1Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Expm1) }

func BenchmarkExpm1Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Expm1) }
