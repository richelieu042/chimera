package sliceKit

// Split 将 buf 按照每组最多 limit 个元素进行分割。
/*
	Split 只创建新的切片头，不会复制 buf 中的元素，因此返回的各个子切片与 buf 共享底层数组。
	修改子切片中的元素，会同时修改 buf 中对应的元素。
	对子切片执行 append 时，如果它仍有剩余容量，还可能覆盖 buf 或后续分组中的元素。
	如果需要完全独立的分组，调用方应复制子切片。
	limit 小于等于 0 时返回 nil。
*/
func Split[T any](buf []T, limit int) [][]T {
	if limit <= 0 {
		return nil
	}

	var chunk []T
	chunks := make([][]T, 0, len(buf)/limit+1)

	for len(buf) >= limit {
		chunk, buf = buf[:limit], buf[limit:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
