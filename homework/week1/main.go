package main

import (
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	CapMinimum         = 1024 // 参考 GO GC 机制实现，小于 1024 不缩容
)

func DeleteAt[T interface{}](src []T, index int) ([]T, error) {
	if index < 0 || index >= len(src) {
		return nil, ErrIndexOutOfRange
	}
	for i := index; i < len(src)-1; i++ {
		src[i] = src[i+1]
	}
	return src, nil
}

// func Shrink[T interface{}](src []T, index int) []T {
// 	c, l := cap(src), len(src)
// 	if !shouldShrink(c, l) {
// 		return src
// 	}
// 	dst := make([]T, 0, c>>1)
// 	dst = append(dst, src...)
// 	return dst
// }

// func shouldShrink(c, l int) bool {
// 	// if c <= CapMinimum { // 原数组小于 1024，不缩容
// 	// 	return false
// 	// }
// 	// if (l << 2) < c { // 长度小于容量的 1/4，触发缩容，直接缩容为原数组的一半
// 	// 	return true
// 	// }
// 	// return false
// 	return c > CapMinimum && (l<<2) < c
// }

func Shrink[T interface{}](src []T, index int) []T {
	c, l := cap(src), len(src)
	if c <= CapMinimum || (l<<2) >= c {
		return src
	}
	dst := make([]T, 0, c>>1)
	dst = append(dst, src...)
	src = nil // 置空使得 src 被 gc 机制回收
	return dst
}
