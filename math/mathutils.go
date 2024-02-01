package mathutils

// Max 返回 a 和 b 中较大的值
func Max[T uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 返回 a 和 b 中较小的值
func Min[T uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}
