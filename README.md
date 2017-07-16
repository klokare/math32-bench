# math32-bench
Benchmarking my math32 library against standard library and other math32 implementations

Benchmark initial run when the math32 library was just calling the standard library's math package and making the casts to/from float32 as required. The table uses the standard library's math package as the baseline. The values below show the relative performance of each 32-bit library compared to the baseline. The table shows 4 packages

* [Golang64](https://golang.org/pkg/math/) - This is the standard library using straight 64-bit floats. This is the target performance. 
* [Klokare](https://github.com/klokare/math32) - The implemenation I am working on
* [chewxy](https://github.com/chewxy/math32) - An API-complete library with some really good improvements and some room for work
* [barnex](https://github.com/barnex/fmath) - An incomplete library that mostly just casts to float64 and calls the standard library. I included this because it will be a good comparison as I work to improve the Klokare libary.


Function|Golang64|Klokare|Chewxy|Barnex
-|-|-|-|-
Abs|1.00|1.69|1.69|2.20
Acos|1.00|1.30|1.30|1.30
Acosh|1.00|1.10|1.83|1.85
Asin|1.00|1.88|1.88|1.88
Asinh|1.00|1.41|1.42|1.40
Atan|1.00|1.73|1.73|1.73
Atan2|1.00|1.63|1.63|1.64
Atanh|1.00|1.82|1.84|1.87
Cbrt|1.00|1.28|1.28|1.29
Ceil|1.00|2.54|13.28|2.54
Copysign|1.00|2.89|1.04|2.88
Cos|1.00|1.50|1.48|1.48
Cosh|1.00|1.28|1.28|1.28
Dim|1.00|3.10|4.25|3.10
Erf|1.00|1.70|1.70|1.70
Erfc|1.00|1.40|1.39|1.39
Exp|1.00|1.59|1.75|1.59
Exp2|1.00|1.20|2.96|1.20
Expm1|1.00|1.78|2.54|1.78
Floor|1.00|2.54|11.94|2.54
Frexp|1.00|1.42|3.22|
Gamma|1.00|1.43|1.43|1.43
Hypot|1.00|2.40|8.22|2.41
Ilogb|1.00|1.50|1.52|
Inf|1.00|1.00|1.07|
IsInf|1.00|2.76|3.40|
IsNaN|1.00|2.54|2.50|
J0|1.00|1.31|1.31|1.32
J1|1.00|1.38|1.39|1.39
Jn|1.00|1.12|1.12|
Ldexp|1.00|1.70|3.11|
Lgamma|1.00|1.26|1.26|
Log|1.00|2.01|1.22|1.97
Log10|1.00|1.50|1.51|1.49
Log1p|1.00|1.91|1.91|1.91
Log2|1.00|1.32|1.29|1.29
Logb|1.00|1.47|1.47|1.48
Max|1.00|4.27|7.28|4.28
Min|1.00|4.25|7.15|4.25
Mod|1.00|1.18|4.03|1.18
Modf|1.00|2.39|2.88|
NaN|1.00|1.00|1.00|
Pow|1.00|1.37|3.20|1.38
Pow10|1.00|1.82|1.81|
Remainder|1.00|1.16|4.22|1.16
Signbit|1.00|3.60|0.85|
Sin|1.00|1.80|1.81|1.81
Sincos|1.00|1.70|1.70|
Sinh|1.00|1.62|1.62|1.63
Sqrt|1.00|1.85|1.26|1.85
Tan|1.00|1.66|1.68|1.66
Tanh|1.00|1.84|2.47|1.88
Trunc|1.00|2.83|12.88|2.83
Y0|1.00|1.32|1.31|1.33
Y1|1.00|1.29|1.30|1.30
Yn|1.00|1.16|1.17|

