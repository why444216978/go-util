package sync

import (
	"runtime"
	"sync/atomic"
)

// POINT 自旋当前次数协程让出CPU
const POINT = 10

// AddInt32 CAS增加int32
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

// AddInt64 CAS增加int64
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

// AddUint32 CAS增加uint32
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

// AddUint64 CAS增加uint64
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
