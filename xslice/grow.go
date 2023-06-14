package xslice

func GrowBools(b []bool, n int) []bool {
	if cap(b)-len(b) < n {
		buf := make([]bool, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowInts(b []int, n int) []int {
	if cap(b)-len(b) < n {
		buf := make([]int, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowInt32s(b []int32, n int) []int32 {
	if cap(b)-len(b) < n {
		buf := make([]int32, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowInt64s(b []int64, n int) []int64 {
	if cap(b)-len(b) < n {
		buf := make([]int64, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowFloat32s(b []float32, n int) []float32 {
	if cap(b)-len(b) < n {
		buf := make([]float32, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowFloat64s(b []float64, n int) []float64 {
	if cap(b)-len(b) < n {
		buf := make([]float64, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowStrings(b []string, n int) []string {
	if cap(b)-len(b) < n {
		buf := make([]string, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}

func GrowBytes(b []byte, n int) []byte {
	if cap(b)-len(b) < n {
		buf := make([]byte, len(b), 2*cap(b)+n)
		copy(buf, b)
		b = buf
	}
	return b
}
