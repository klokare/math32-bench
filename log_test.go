package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkLogKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Log) }

func BenchmarkLogGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Log) }

func BenchmarkLogChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Log) }

func BenchmarkLogBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Log) }
