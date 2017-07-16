package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLog1pKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Log1p) }

func BenchmarkLog1pGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Log1p) }

func BenchmarkLog1pChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Log1p) }

func BenchmarkLog1pBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Log1p) }
