package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLog10Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Log10) }

func BenchmarkLog10Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Log10) }

func BenchmarkLog10Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Log10) }

func BenchmarkLog10Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Log10) }

func BenchmarkLog2Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Log2) }

func BenchmarkLog2Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Log2) }

func BenchmarkLog2Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Log2) }

func BenchmarkLog2Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Log2) }
