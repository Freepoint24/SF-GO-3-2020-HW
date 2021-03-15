package fmtnum

import "strconv"

func F64(a float64) string  {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func C128(a complex128) string  {
	return strconv.FormatComplex(a, 'f', -1, 128)
}
