package xslice

import "context"

// example:
//
//	var ls []string
//	_, err := PaginationStartID(context.TODO(), 10, func(startID int64, limit int64) (lastID int64, done bool, err error) {
//		var tmp []string
//		query(&tmp,"... where id >= ? limit ?", startID, limit)
//		if len(tmp) == 0 {
//			return startID, true, nil
//		}
//		lastID = max(tmp)
//		ls = append(ls, tmp...)
//		return lastID, len(tmp)<=0, nil
//	})
func PaginationStartID(ctx context.Context, limit int64, fn func(startID int64, limit int64) (lastID int64, done bool, err error)) (err error) {
	return PaginationStartIDWithParams(ctx, 0 /*startID*/, limit /*limit*/, true /*asc*/, fn)
}

func PaginationStartIDWithParams(ctx context.Context, startID int64, limit int64, asc bool, fn func(startID int64, limit int64) (lastID int64, done bool, err error)) (err error) {
	if limit < 0 {
		limit = 10
	}
	var step int64 = +1
	if !asc {
		step = -1
	}
	var done bool
	var lastID int64
	for {
		lastID, done, err = fn(startID, limit)
		if err != nil {
			break
		}
		if done {
			break
		}
		if startID == lastID {
			break
		}
		startID = lastID + step
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}
	}
	return
}
