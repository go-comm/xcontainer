package transform

import "reflect"

func Values2Interfaces(vs []reflect.Value) []interface{} {
	var ls []interface{} = make([]interface{}, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.Interface())
	}
	return ls
}

func Values2Bools(vs []reflect.Value) []bool {
	var ls []bool = make([]bool, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.Bool())
	}
	return ls
}

func Values2Int8s(vs []reflect.Value) []int8 {
	var ls []int8 = make([]int8, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, int8(vv.Int()))
	}
	return ls
}

func Values2Int16s(vs []reflect.Value) []int16 {
	var ls []int16 = make([]int16, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, int16(vv.Int()))
	}
	return ls
}

func Values2Ints(vs []reflect.Value) []int {
	var ls []int = make([]int, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, int(vv.Int()))
	}
	return ls
}

func Values2Uints(vs []reflect.Value) []uint {
	var ls []uint = make([]uint, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, uint(vv.Uint()))
	}
	return ls
}

func Values2Int32s(vs []reflect.Value) []int32 {
	var ls []int32 = make([]int32, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, int32(vv.Int()))
	}
	return ls
}

func Values2Uint32s(vs []reflect.Value) []uint32 {
	var ls []uint32 = make([]uint32, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, uint32(vv.Uint()))
	}
	return ls
}

func Values2Int64s(vs []reflect.Value) []int64 {
	var ls []int64 = make([]int64, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.Int())
	}
	return ls
}

func Values2Uint64s(vs []reflect.Value) []uint64 {
	var ls []uint64 = make([]uint64, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.Uint())
	}
	return ls
}

func Values2Strings(vs []reflect.Value) []string {
	var ls []string = make([]string, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.String())
	}
	return ls
}

func Values2Float32s(vs []reflect.Value) []float32 {
	var ls []float32 = make([]float32, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, float32(vv.Float()))
	}
	return ls
}

func Values2Float64s(vs []reflect.Value) []float64 {
	var ls []float64 = make([]float64, 0, len(vs))
	for _, vv := range vs {
		ls = append(ls, vv.Float())
	}
	return ls
}
