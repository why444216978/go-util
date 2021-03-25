package runtime

import (
	"runtime"
	"strconv"
)

// GetGoroutineId 获得当前 goroutine id
func GetGoroutineId() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := string(buf[:n])

	str := stk[10:11]

	return strconv.Atoi(str)
}
