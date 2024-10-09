package cmp

import (
	"github.com/samber/lo"
)

type Comparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func CompareLT[T Comparable](v1 T, v2 T) bool {
	return v1 < v2
}

func CompareLE[T Comparable](v1 T, v2 T) bool {
	return v1 <= v2
}

func CompareEQ[T comparable](v1 T, v2 T) bool {
	return v1 == v2
}

func CompareNE[T comparable](v1 T, v2 T) bool {
	return v1 != v2
}

func CompareGT[T Comparable](v1 T, v2 T) bool {
	return v1 > v2
}

func CompareGE[T Comparable](v1 T, v2 T) bool {
	return v1 >= v2
}

func CompareIN[T comparable](v1 []T, v2 T) bool {
	return lo.Count(v1, v2) > 0
}
