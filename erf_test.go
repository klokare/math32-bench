package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkErfKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Erf) }

func BenchmarkErfGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Erf) }

func BenchmarkErfChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Erf) }

func BenchmarkErfBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Erf) }

func BenchmarkErfcKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Erfc) }

func BenchmarkErfcGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Erfc) }

func BenchmarkErfcChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Erfc) }

func BenchmarkErfcBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Erfc) }
