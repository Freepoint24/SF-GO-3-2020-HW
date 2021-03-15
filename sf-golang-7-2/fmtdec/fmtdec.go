package fmtdec

import "strconv"

// F - форматирует вещественное или комплексное число в строку с десятичной точкой
// с минимально необходимым числом знаков после запятой.
//
// Поддерживаемые типы: float32, float64, complex64, complex128
func F(a interface{}) string {
	switch a.(type) {
	case float32:
		return strconv.FormatFloat(float64(a.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(a.(float64), 'f', -1, 64)
	case complex64:
		return strconv.FormatComplex(complex128(a.(complex64)), 'f', -1, 64)
	case complex128:
		return strconv.FormatComplex(a.(complex128), 'f', -1, 128)
	default:
		panic("unknown type")
	}
}
