package xslice

import "time"

func CompareInt8(a int8, b int8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareInt16(a int16, b int16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareInt(a int, b int) int {
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

func CompareUint8(a uint8, b uint8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func CompareUint16(a uint16, b uint16) int {
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

func CompareByte(a byte, b byte) int {
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

func CompareTime(a time.Time, b time.Time) int {
	d := a.Sub(b)
	if d == 0 {
		return 0
	}
	if d < 0 {
		return -1
	}
	return +1
}

func CompareDuration(a time.Duration, b time.Duration) int {
	d := a - b
	if d == 0 {
		return 0
	}
	if d < 0 {
		return -1
	}
	return +1
}
