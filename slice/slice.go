package slice

import (
	"errors"
	"reflect"
)

// 切片快捷操作汇总：
// a := []int{1, 2, 3}
// b := []int{4, 5, 6}
// i := 1
// j := 3
// 1.将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
// 2.删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
// 3.切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
// 4.为切片 a 扩展 j 个元素长度：a = append(a, make([]int, j)...)
// 5.在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
// 6.在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]int, j), a[i:]...)...)
// 7.在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
// 8.取出位于切片 a 最末尾的元素 x：x, a := a[len(a)-1:], a[:len(a)-1]

// DeleteSliceByPos 删除切片指定位置元素
func DeleteSliceByPos(slice interface{}, index int) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return slice, errors.New("index error")
	}
	return reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface(), nil
}

// InsertSliceByIndex 在指定位置插入元素
func InsertSliceByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if index < 0 || index > v.Len() || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("index error")
	}
	if index == v.Len() {
		return reflect.Append(v, reflect.ValueOf(value)).Interface(), nil
	}
	v = reflect.AppendSlice(v.Slice(0, index+1), v.Slice(index, v.Len()))
	v.Index(index).Set(reflect.ValueOf(value))
	return v.Interface(), nil
}

// UpdateSliceByIndex 更新指定位置元素
func UpdateSliceByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return slice, errors.New("not slice")
	}
	if index > v.Len()-1 || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("index error")
	}
	v.Index(index).Set(reflect.ValueOf(value))

	return v.Interface(), nil
}

// ContainsInterface 是否包含指定interface
func ContainsInterface(sl []interface{}, v interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsInt 是否包含指定int
func ContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsInt64 是否包含指定int64
func ContainsInt64(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsString 是否包含指定string
func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// UniqueInt64 int64切片去重
func UniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	m := make(map[int64]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int64, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// UniqueInt int切片去重
func UniqueInt(s []int) []int {
	size := len(s)
	if size == 0 {
		return []int{}
	}

	m := make(map[int]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// UniqueString string切片去重
func UniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	m := make(map[string]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]string, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}

// SumInt64 int64切片求和
func SumInt64(intSlice []int64) (sum int64) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// SumInt int切片求和
func SumInt(intSlice []int) (sum int) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// SumFloat64 float64切片求和
func SumFloat64(intSlice []float64) (sum float64) {
	for _, v := range intSlice {
		sum += v
	}
	return
}
