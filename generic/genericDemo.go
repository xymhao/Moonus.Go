package generic

//non-generic functions

// SumIntNumber demo
func SumIntNumber(m map[string]int) int {
	var s int

	for _, v := range m {
		s += v
	}

	return s
}

// SumFloatNumber demo
func SumFloatNumber(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// generic

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
//func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//	s += v
//	}
//	return s
//}
