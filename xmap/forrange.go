package xmap

import "reflect"

func forrange(m interface{}, f func(k reflect.Value, v reflect.Value) (continued bool)) {
	rv := reflect.Indirect(reflect.ValueOf(m))
	keys := rv.MapKeys()
	for _, k := range keys {
		if !f(k, rv.MapIndex(k)) {
			break
		}
	}
}
