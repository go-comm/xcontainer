package xslice

import "context"

// example:
//
//	var ls []string
//	_, err := PaginationStartID(context.TODO(), 10, func(startID int64, limit int64) (lastID int64, n int64, err error) {
//		var tmp []string
//		query(&tmp,"... where id >= ? limit ?", startID, limit)
//		if len(tmp) == 0 {
//			return startID, 0, nil
//		}
//		lastID = max(tmp)
//		ls = append(ls, tmp...)
//		return lastID, int64(len(tmp)), nil
//	})
func PaginationStartID(ctx context.Context, limit int64, fn func(startID int64, limit int64) (lastID int64, n int64, err error)) (total int64, err error) {
	return PaginationStartIDWithParams(ctx, 0 /*startID*/, limit /*limit*/, -1 /*maxTotal*/, fn)
}

func PaginationStartIDWithParams(ctx context.Context, startID int64, limit int64, maxTotal int64, fn func(startID int64, limit int64) (lastID int64, n int64, err error)) (total int64, err error) {
	if limit < 0 {
		limit = 10
	}
	var n int64
	var lastID int64 = startID
	for {
		lastID, n, err = fn(lastID, limit)
		if err != nil {
			break
		}
		if n <= 0 {
			break
		}
		if lastID <= startID {
			break
		}
		lastID++
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
