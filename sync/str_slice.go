package sync

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

// ConcurrentStrArray string类型并发安全切片
type ConcurrentStrArray interface {
	// Set 用于设置指定索引上的元素值。
	Set(index int, elem string) (old string, err error)
	// Get 用于获取指定索引上的元素值。
	Get(index int) (elem string, err error)
	// Len 用于获取数组的长度。
	Len() int
}

type segmentStr struct {
	val    atomic.Value
	length int    // 段内的元素数量。
	status uint32 // 0：可读可写；1：只读。
}

func (seg *segmentStr) init(length int) {
	seg.length = length
	seg.val.Store(make([]string, length))
}

func (seg *segmentStr) checkIndex(index int) error {
	if index < 0 || index >= seg.length {
		return fmt.Errorf("index out of range [0, %d) in segment", seg.length)
	}
	return nil
}

func (seg *segmentStr) set(index int, elem string) (old string, err error) {
	if err = seg.checkIndex(index); err != nil {
		return
	}
	point := 10 //TODO 此处是一个优化点，可以根据实际情况调整。
	count := 0
	for { // 简易的自旋锁。
		count++
		if !atomic.CompareAndSwapUint32(&seg.status, 0, 1) {
			if count%point == 0 {
				runtime.Gosched()
			}
			continue
		}
		defer atomic.StoreUint32(&seg.status, 0)
		newArray := make([]string, seg.length)
		copy(newArray, seg.val.Load().([]string))
		old = newArray[index]
		newArray[index] = elem
		seg.val.Store(newArray)
		return
	}
}

func (seg *segmentStr) get(index int) (elem string, err error) {
	if err = seg.checkIndex(index); err != nil {
		return
	}
	elem = seg.val.Load().([]string)[index]
	return
}

type myStrArray struct {
	length    int           // 元素总数量。
	segLenStd int           // 单个内部段的标准长度。
	segments  []*segmentStr // 内部段列表。
}

// NewConcurrentStrArray 工厂方法
func NewConcurrentStrArray(length int) ConcurrentStrArray {
	if length < 0 {
		length = 0
	}
	array := new(myStrArray)
	array.init(length)
	return array
}

func (array *myStrArray) init(length int) {
	array.length = length
	array.segLenStd = 10 //TODO 此处是一个优化点，可以根据参数值调整。
	segNum := length / array.segLenStd
	segLenTail := length % array.segLenStd
	if segLenTail > 0 {
		segNum = segNum + 1
	}
	array.segments = make([]*segmentStr, segNum)
	for i := 0; i < segNum; i++ {
		seg := segmentStr{}
		if i == segNum-1 && segLenTail > 0 {
			seg.init(segLenTail)
		} else {
			seg.init(array.segLenStd)
		}
		array.segments[i] = &seg
	}
}

func (array *myStrArray) Set(index int, elem string) (old string, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	seg := array.segments[index/array.segLenStd]
	return seg.set(index%array.segLenStd, elem)
}

func (array *myStrArray) Get(index int) (elem string, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	seg := array.segments[index/array.segLenStd]
	return seg.get(index % array.segLenStd)
}

func (array *myStrArray) Len() int {
	return array.length
}

func (array *myStrArray) checkIndex(index int) error {
	if index < 0 || index >= array.length {
		return fmt.Errorf("index out of range [0, %d)", array.length)
	}
	return nil
}
