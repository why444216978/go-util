package binary

// SetBit 将指定的二进制位设置为1
func SetBit(n int, pos int) int {
	return n | (1 << pos)
}

// CheckBit 判断指定二进制位是否为1
func CheckBit(n int, pos int) bool {
	return n&(1<<pos) == (1 << pos)
}

// ByteList2Int 字节列表转整型，list需要满足遍历低位->高位
func ByteList2Int(list []uint8) int {
	size := len(list)
	if size == 0 {
		return 0
	}
	n := 0
	for pos := size - 1; pos >= 0; pos-- {
		if list[pos] > 0 {
			n = SetBit(n, pos)
		}
	}
	return n
}

// IntToByteList 整型转字节列表，list遍历低位->高位
func IntToByteList(n int, max int) []uint8 {
	res := make([]uint8, max)
	for pos := 0; pos < max; pos++ {
		if CheckBit(n, pos) {
			res[pos] = 1
		}
	}
	return res
}
