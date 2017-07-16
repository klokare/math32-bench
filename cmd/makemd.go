package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

var libs = []string{"Golang64", "Klokare", "Chewxy", "Barnex"}
var funcs = []string{
	"Abs", "Acos", "Acosh", "Asin", "Asinh", "Atan", "Atan2", "Atanh", "Cbrt", "Ceil", "Copysign", "Cos",
	"Cosh", "Dim", "Erf", "Erfc", "Exp", "Exp2", "Expm1", "Floor", "Frexp", "Gamma", "Hypot", "Ilogb", "Inf",
	"IsInf", "IsNaN", "J0", "J1", "Jn", "Ldexp", "Lgamma", "Log", "Log10", "Log1p", "Log2", "Logb", "Max",
	"Min", "Mod", "Modf", "NaN", "Pow", "Pow10", "Remainder", "Signbit", "Sin", "Sincos", "Sinh", "Sqrt", "Tan",
	"Tanh", "Trunc", "Y0", "Y1", "Yn",
}

// As this is just a helper utility, some assumptions were made:
// 1. you are running this on a Mac or Linux box
// 2. you are executing this from within the test source directory (go run cmd/makemd.go)
// 3. error checking kept at a minimum.
func main() {

	// Create the names and map to libraries and functions
	lm := make(map[string]string, len(libs)*len(funcs))
	fm := make(map[string]string, len(libs)*len(funcs))
	for _, ln := range libs {
		for _, fn := range funcs {
			n := "Benchmark" + fn + ln
			lm[n] = ln
			fm[n] = fn
		}
	}

	// Execute the benchmarks N times to a temporary file
	var err error
	lines := make([]string, 0, 100*50*4)
	n := 2
	for i := 0; i < n; i++ {
		log.Println("Running benchmark set", i+1)
		var v []byte
		if v, err = exec.Command("go", "test", "-bench=.", "-benchmem").Output(); err != nil {
			log.Fatal(err)
		}

		// Transform output in to lines
		buf := bytes.NewBuffer(v)

		var line string
		for {
			line, err = buf.ReadString('\n')
			if err != nil {
				break
			}
			if strings.HasPrefix(line, "Benchmark") {
				lines = append(lines, line)
			}
		}
	}
	log.Println("Lines", len(lines))

	// Parse the output
	var marks []*parse.Benchmark
	if marks, err = parseBenchmarks(lines); err != nil {
		log.Fatal(err)
	}
	log.Println("Marks", len(marks))

	// Extract the ns/op
	var ok bool
	m := make(map[string]map[string][]float64, 5) // library:function:[]ns
	for _, bm := range marks {

		// Note the library's map
		var libm map[string][]float64
		ln := lm[bm.Name]
		if libm, ok = m[ln]; !ok {
			libm = make(map[string][]float64, 50)
		}

		// Note the values
		var vals []float64
		fn := fm[bm.Name]
		if vals, ok = libm[fn]; !ok {
			vals = make([]float64, 0, 100)
		}

		// Append the data and save the colletions
		libm[fn] = append(vals, bm.NsPerOp)
		m[ln] = libm
	}

	// Output the markdown table
	var f *os.File
	if f, err = os.Create("math32-bench.md"); err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	w.WriteString("Function")
	for _, ln := range libs {
		w.WriteString("|" + ln)
	}
	w.WriteString("\n-")
	for i := 0; i < len(libs); i++ {
		w.WriteString("|-")
	}
	w.WriteString("\n")

	for _, fn := range funcs {
		w.WriteString(fn)
		var base float64
		for i, ln := range libs {
			w.WriteString("|")
			vals := m[ln][fn]
			if len(vals) == 0 {
				continue
			}
			med := median(vals)
			if i == 0 {
				base = med
				w.WriteString(fmt.Sprintf("%0.2f", 1.0))
				//w.WriteString(fmt.Sprintf("%0.4f", base))
			} else {
				w.WriteString(fmt.Sprintf("%0.2f", med/base))
				//w.WriteString(fmt.Sprintf("%0.4f", med))
			}
		}
		w.WriteString("\n")
	}

}

func parseBenchmarks(lines []string) (marks []*parse.Benchmark, err error) {
	marks = make([]*parse.Benchmark, 0, len(lines))
	for _, line := range lines {
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}
		var bm *parse.Benchmark
		if bm, err = parse.ParseLine(line); err != nil {
			return
		}

		// Take the opportunity to strip the processor count from the name
		idx := strings.Index(bm.Name, "-")
		if idx > 0 {
			bm.Name = bm.Name[:idx]
		}

		marks = append(marks, bm)
	}
	return
}

func median(xs []float64) float64 {
	switch len(xs) {
	case 0:
		return 0
	case 1:
		return xs[0]
	default:
		sort.Float64s(xs)
		i := len(xs) / 2
		if len(xs)%2 == 0 {
			return (xs[i-1] + xs[i]) / 2.0
		}
		return xs[i]
	}
}
