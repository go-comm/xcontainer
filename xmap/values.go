package xmap

import (
	"reflect"
)

func Values(m interface{}, rv interface{}) {
	rdst := reflect.Indirect(reflect.ValueOf(rv))
	forrange(m, func(k, v reflect.Value) (continued bool) {
		rdst.Set(reflect.Append(rdst, v))
		return true
	})
}

func ValuesOfInterfaces(m interface{}) []interface{} {
	var ls []interface{}
	forrange(m, func(k, v reflect.Value) (continued bool) {
		ls = append(ls, v.Interface())
		return true
	})
	return ls
}

func ValuesOfStrings(m interface{}) []string {
	var ls []string
	forrange(m, func(k, v reflect.Value) (continued bool) {
		ls = append(ls, v.String())
		return true
	})
	return ls
}

func ValuesOfInts(m interface{}) []int {
	var ls []int
	forrange(m, func(k, v reflect.Value) (continued bool) {
		ls = append(ls, int(v.Int()))
		return true
	})
	return ls
}

func ValuesOfInt32s(m interface{}) []int32 {
	var ls []int32
	forrange(m, func(k, v reflect.Value) (continued bool) {
		ls = append(ls, int32(v.Int()))
		return true
	})
	return ls
}

func ValuesOfInt64s(m interface{}) []int64 {
	var ls []int64
	forrange(m, func(k, v reflect.Value) (continued bool) {
		ls = append(ls, int64(v.Int()))
		return true
	})
	return ls
}
