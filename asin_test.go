package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkAsinKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Asin) }

func BenchmarkAsinGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Asin) }

func BenchmarkAsinChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Asin) }

func BenchmarkAsinBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Asin) }

func BenchmarkAcosKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Acos) }

func BenchmarkAcosGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Acos) }

func BenchmarkAcosChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Acos) }

func BenchmarkAcosBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Acos) }
