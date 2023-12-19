package xmap

import (
	"reflect"

	"github.com/go-comm/xcontainer/internal/transform"
)

func keys(m interface{}) []reflect.Value {
	rv := reflect.Indirect(reflect.ValueOf(m))
	return rv.MapKeys()
}

func KeysOfInterfaces(m interface{}) []interface{} {
	return transform.Values2Interfaces(keys(m))
}

func KeysOfStrings(m interface{}) []string {
	return transform.Values2Strings(keys(m))
}

func KeysOfInts(m interface{}) []int {
	return transform.Values2Ints(keys(m))
}

func KeysOfInt32s(m interface{}) []int32 {
	return transform.Values2Int32s(keys(m))
}

func KeysOfInt64s(m interface{}) []int64 {
	return transform.Values2Int64s(keys(m))
}

func KeysOfFloat32s(m interface{}) []float32 {
	return transform.Values2Float32s(keys(m))
}

func KeysOfFloat64s(m interface{}) []float64 {
	return transform.Values2Float64s(keys(m))
}
