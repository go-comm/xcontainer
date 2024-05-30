package xslice

import (
	"context"
)

// example:
//
//	var ls []string
//	_, err := PaginationDeep(context.TODO(), 10, func(limit int64, offset int64) (int64, error) {
//		var tmp []string
//		query(&tmp,"... limit ? offset ?", limit, offset)
//		ls = append(ls, tmp...)
//		return int64(len(tmp)), nil
//	})
func PaginationDeep(ctx context.Context, limit int64, fn func(limit int64, offset int64) (int64, error)) (total int64, err error) {
	return PaginationDeepWithParams(ctx, limit /*limit*/, 0 /*offset*/, -1 /*maxTotal*/, fn)
}

func PaginationDeepWithParams(ctx context.Context, limit int64, offset int64, maxTotal int64, fn func(limit int64, offset int64) (int64, error)) (total int64, err error) {
	if limit < 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	var n int64
	for ; ; offset += limit {
		n, err = fn(limit, offset)
		if err != nil {
			break
		}
		if n <= 0 {
			break
		}
		total += n
		if maxTotal > 0 && total >= maxTotal {
			break
		}
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}
	}
	return
}

func PaginationDeepInts(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int, error)) ([]int, error) {
	var ls []int
	_, err := PaginationDeep(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationDeepInt32s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int32, error)) ([]int32, error) {
	var ls []int32
	_, err := PaginationDeep(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationDeepInt64s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int64, error)) ([]int64, error) {
	var ls []int64
	_, err := PaginationDeep(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationDeepStrings(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]string, error)) ([]string, error) {
	var ls []string
	_, err := PaginationDeep(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}
