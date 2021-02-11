package sync

import (
	"runtime"
	"sync/atomic"
)

const POINT = 10 //TODO 可优化

func AddInt32(val, add int32) {
	count := 0
	for {
		count++
		if atomic.CompareAndSwapInt32(&val, val, (val + add)) {
			if count%POINT == 0 {
				runtime.Gosched()
			}
			continue
		}
	}
}

func AddInt64(val, add int64) {
	count := 0
	for {
		count++
		if atomic.CompareAndSwapInt64(&val, val, (val + add)) {
			if count%POINT == 0 {
				runtime.Gosched()
			}
			continue
		}
	}
}

func AddUint32(val, add uint32) {
	count := 0
	for {
		count++
		if atomic.CompareAndSwapUint32(&val, val, (val + add)) {
			if count%POINT == 0 {
				runtime.Gosched()
			}
			continue
		}
	}
}

func AddUint64(val, add uint64) {
	count := 0
	for {
		count++
		if atomic.CompareAndSwapUint64(&val, val, (val + add)) {
			if count%POINT == 0 {
				runtime.Gosched()
			}
			continue
		}
	}
}
