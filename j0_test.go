package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkJ0Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.J0) }

func BenchmarkJ0Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.J0) }

func BenchmarkJ0Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.J0) }

func BenchmarkJ0Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.J0) }

func BenchmarkY0Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Y0) }

func BenchmarkY0Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Y0) }

func BenchmarkY0Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Y0) }

func BenchmarkY0Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Y0) }
