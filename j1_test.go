package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkJ1Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.J1) }

func BenchmarkJ1Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.J1) }

func BenchmarkJ1Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.J1) }

func BenchmarkJ1Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.J1) }

func BenchmarkY1Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Y1) }

func BenchmarkY1Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Y1) }

func BenchmarkY1Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Y1) }

func BenchmarkY1Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Y1) }
