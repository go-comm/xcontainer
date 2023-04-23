package xslice

func AppendJoinBytes(b []byte, length int, f func(i int) []byte, sep []byte) []byte {
	if length == 0 {
		return b
	}
	if length > 0 {
		b = append(b, f(0)...)
	}
	if length > 1 {
		for i := 1; i < length; i++ {
			b = append(b, sep...)
			b = append(b, f(i)...)
		}
	}
	return b
}

func JoinBytes(length int, f func(i int) []byte, sep []byte) []byte {
	var b []byte
	b = AppendJoinBytes(b, length, f, sep)
	return b
}

func AppendJoinStr(b []byte, length int, f func(i int) string, sep string) []byte {
	if length == 0 {
		return b
	}
	if length > 0 {
		b = append(b, f(0)...)
	}
	if length > 1 {
		for i := 1; i < length; i++ {
			b = append(b, sep...)
			b = append(b, f(i)...)
		}
	}
	return b
}

func JoinStr(length int, f func(i int) string, sep string) string {
	var b []byte
	b = AppendJoinStr(b, length, f, sep)
	return string(b)
}
