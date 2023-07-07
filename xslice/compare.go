package xslice

func CompareInt(a int, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareUint(a uint, b uint) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareInt32(a int32, b int32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareInt64(a int64, b int64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareUint32(a uint32, b uint32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareUint64(a uint64, b uint64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareFloat32(a float32, b float32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareFloat64(a float64, b float64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareString(a string, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
