package defined_identifiers

import (
	"math"
)

var Functions = map[string]func(x float64) float64{
	"sqrt":  math.Sqrt,
	"abs":   math.Abs,
	"exp":   math.Exp,
	"ln":    math.Log,
	"log":   math.Log,
	"sin":   math.Sin,
	"cos":   math.Cos,
	"tan":   math.Tan,
	"asin":  math.Asin,
	"acos":  math.Acos,
	"atan":  math.Atan,
	"sinh":  math.Sinh,
	"cosh":  math.Cosh,
	"tanh":  math.Tanh,
	"asinh": math.Asinh,
	"acosh": math.Acosh,
	"atanh": math.Atanh,
}
