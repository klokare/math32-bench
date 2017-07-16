package bench

import (
	"math"
	"testing"

	barnex "github.com/barnex/fmath"
	chewxy "github.com/chewxy/math32"
	"github.com/klokare/math32"
)

func BenchmarkFloorKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Floor) }

func BenchmarkFloorGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Floor) }

func BenchmarkFloorChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Floor) }

func BenchmarkFloorBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Floor) }

func BenchmarkCeilKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Ceil) }

func BenchmarkCeilGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Ceil) }

func BenchmarkCeilChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Ceil) }

func BenchmarkCeilBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Ceil) }

func BenchmarkTruncKlokare(b *testing.B) { benchOneFloatToOneFloat32(b, math32.Trunc) }

func BenchmarkTruncGolang64(b *testing.B) { benchOneFloatToOneFloat64(b, math.Trunc) }

func BenchmarkTruncChewxy(b *testing.B) { benchOneFloatToOneFloat32(b, chewxy.Trunc) }

func BenchmarkTruncBarnex(b *testing.B) { benchOneFloatToOneFloat32(b, barnex.Trunc) }
