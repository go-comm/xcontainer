package xslice

import "context"

func Partition(ctx context.Context, length int, limit int, fn func(i int, j int) error) error {
	offset := 0
	for {
		if offset >= length {
			break
		}
		end := offset + limit
		if end > length {
			end = length
		}
		if err := fn(offset, end); err != nil {
			return err
		}
		offset = end

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	return nil
}

func PartitionInts(ctx context.Context, arr []int, limit int, fn func(sub []int) error) error {
	return Partition(ctx, len(arr), limit, func(i, j int) error { return fn(arr[i:j]) })
}

func PartitionInt32s(ctx context.Context, arr []int32, limit int, fn func(sub []int32) error) error {
	return Partition(ctx, len(arr), limit, func(i, j int) error { return fn(arr[i:j]) })
}

func PartitionInt64s(ctx context.Context, arr []int64, limit int, fn func(sub []int64) error) error {
	return Partition(ctx, len(arr), limit, func(i, j int) error { return fn(arr[i:j]) })
}

func PartitionStrings(ctx context.Context, arr []string, limit int, fn func(sub []string) error) error {
	return Partition(ctx, len(arr), limit, func(i, j int) error { return fn(arr[i:j]) })
}
