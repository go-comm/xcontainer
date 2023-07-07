package xslice

func SumToInt(length int, m func(i int) int) int {
	var s int
	for i := 0; i < length; i++ {
		s += m(i)
	}
	return s
}

func SumToInt32(length int, m func(i int) int32) int32 {
	var s int32
	for i := 0; i < length; i++ {
		s += m(i)
	}
	return s
}

func SumToInt64(length int, m func(i int) int64) int64 {
	var s int64
	for i := 0; i < length; i++ {
		s += m(i)
	}
	return s
}

func SumToFloat32(length int, m func(i int) float32) float32 {
	var s float32
	for i := 0; i < length; i++ {
		s += m(i)
	}
	return s
}

func SumToFloat64(length int, m func(i int) float64) float64 {
	var s float64
	for i := 0; i < length; i++ {
		s += m(i)
	}
	return s
}
