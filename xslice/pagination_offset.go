package xslice

import (
	"context"
)

// example:
//
//	var ls []string
//	err := PaginationOffset(context.TODO(), 10, func(limit int64, offset int64) (int64, error) {
//		var tmp []string
//		query(&tmp,"... limit ? offset ?", limit, offset)
//		ls = append(ls, tmp...)
//		return len(tmp)<=0, nil
//	})
func PaginationOffset(ctx context.Context, limit int64, fn func(limit int64, offset int64) (done bool, err error)) (err error) {
	return PaginationOffsetWithParams(ctx, limit /*limit*/, 0 /*offset*/, fn)
}

func PaginationOffsetWithParams(ctx context.Context, limit int64, offset int64, fn func(limit int64, offset int64) (done bool, err error)) (err error) {
	if limit < 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	var done bool
	for ; ; offset += limit {
		done, err = fn(limit, offset)
		if err != nil {
			break
		}
		if done {
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
	err := PaginationOffset(ctx, limit, func(limit, offset int64) (bool, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return false, err
		}
		ls = append(ls, tmp...)
		return len(tmp) <= 0, nil
	})
	return ls, err
}

func PaginationOffsetInt32s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int32, error)) ([]int32, error) {
	var ls []int32
	err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (bool, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return false, err
		}
		ls = append(ls, tmp...)
		return len(tmp) <= 0, nil
	})
	return ls, err
}

func PaginationOffsetInt64s(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]int64, error)) ([]int64, error) {
	var ls []int64
	err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (bool, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return false, err
		}
		ls = append(ls, tmp...)
		return len(tmp) <= 0, nil
	})
	return ls, err
}

func PaginationOffsetStrings(ctx context.Context, limit int64, fn func(limit int64, offset int64) ([]string, error)) ([]string, error) {
	var ls []string
	err := PaginationOffset(ctx, limit, func(limit int64, offset int64) (bool, error) {
		tmp, err := fn(limit, offset)
		if err != nil {
			return false, err
		}
		ls = append(ls, tmp...)
		return len(tmp) <= 0, nil
	})
	return ls, err
}
