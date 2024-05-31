package xslice

import (
	"context"
)

// example:
//
//	var ls []string
//	_, err := PaginationOffset(context.TODO(), 10, func(limit int64, offset int64) (int64, error) {
//		var tmp []string
//		query(&tmp,"... limit ? offset ?", limit, offset)
//		ls = append(ls, tmp...)
//		return int64(len(tmp)), nil
//	})
func PaginationOffset(ctx context.Context, limit int64, fn func(limit int64, offset int64) (int64, error)) (total int64, err error) {
	return PaginationOffsetWithParams(ctx, limit /*limit*/, 0 /*offset*/, -1 /*maxTotal*/, fn)
}

func PaginationOffsetWithParams(ctx context.Context, limit int64, offset int64, maxTotal int64, fn func(limit int64, offset int64) (int64, error)) (total int64, err error) {
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

func PaginationOffsetInts(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int, error)) ([]int, error) {
	var ls []int
	_, err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationOffsetInt32s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int32, error)) ([]int32, error) {
	var ls []int32
	_, err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationOffsetInt64s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int64, error)) ([]int64, error) {
	var ls []int64
	_, err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}

func PaginationOffsetStrings(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]string, error)) ([]string, error) {
	var ls []string
	_, err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (int64, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return 0, err
		}
		ls = append(ls, tmp...)
		return int64(len(tmp)), nil
	})
	return ls, err
}
