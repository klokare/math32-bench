package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkExpKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Exp) }

func BenchmarkExpGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Exp) }

func BenchmarkExpChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Exp) }

func BenchmarkExpBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Exp) }

func BenchmarkExp2Klokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Exp2) }

func BenchmarkExp2Golang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Exp2) }

func BenchmarkExp2Chewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Exp2) }

func BenchmarkExp2Barnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Exp2) }
