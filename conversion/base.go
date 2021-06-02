package conversion

import (
	"reflect"
	"strconv"
	"unsafe"
)

// StringToByte 自定义string转byte
// 解决强转内存分配效率低问题
//	 StringToByte("str")
func StringToByte(str string) []byte {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	sliceHeader.Cap = sliceHeader.Len
	b := *(*[]byte)(unsafe.Pointer(sliceHeader))
	return b
}

// ByteToString 自定义byte转string
// 解决强转效率低：先分配一个内存再复制内容
//	 ByteToString([]byte("why"))
func ByteToString(b []byte) string {
	ptr := (*string)(unsafe.Pointer(&b))
	return *ptr
}

// StringToUint8 string转uint
func StringToUint8(str string) uint8 {
	intTmp, _ := strconv.Atoi(str)
	return uint8(intTmp)
}

// StringToUint16 string转uint16
func StringToUint16(str string) uint16 {
	intTmp, _ := strconv.Atoi(str)
	return uint16(intTmp)
}

// StringToUint64 string转uint64
func StringToUint64(str string) uint64 {
	intTmp, _ := strconv.Atoi(str)
	return uint64(intTmp)
}

// StringToInt8 string转int8
func StringToInt8(str string) int8 {
	intTmp, _ := strconv.Atoi(str)
	return int8(intTmp)
}

// StringToInt16 string转int16
func StringToInt16(str string) int16 {
	intTmp, _ := strconv.Atoi(str)
	return int16(intTmp)
}

// StringToInt32 string转int32
func StringToInt32(str string) int32 {
	intTmp, _ := strconv.Atoi(str)
	return int32(intTmp)
}

// StringToInt64 string转int64
func StringToInt64(str string) int64 {
	intTmp, _ := strconv.Atoi(str)
	return int64(intTmp)
}

// InterfaceStringToUint8 interface转uint8
func InterfaceStringToUint8(str interface{}) uint8 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return uint8(intTmp)
}

// InterfaceStringToUint16 interface转uint16
func InterfaceStringToUint16(str interface{}) uint16 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return uint16(intTmp)
}

// InterfaceStringToUint64 interface转uint64
func InterfaceStringToUint64(str interface{}) uint64 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return uint64(intTmp)
}

// InterfaceStringToInt8 interface转int8
func InterfaceStringToInt8(str interface{}) int8 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return int8(intTmp)
}

// InterfaceStringToInt16 interface转int16
func InterfaceStringToInt16(str interface{}) int16 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return int16(intTmp)
}

// InterfaceStringToInt32 interface转in32
func InterfaceStringToInt32(str interface{}) int32 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return int32(intTmp)
}

// InterfaceStringToInt64 interface转int64
func InterfaceStringToInt64(str interface{}) int64 {
	strTmp, ok := str.(string)
	if !ok {
		return 0
	}
	intTmp, _ := strconv.Atoi(strTmp)
	return int64(intTmp)
}

// InterfaceFloat64ToUint8 interface转uint8
func InterfaceFloat64ToUint8(str interface{}) uint8 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return uint8(tmp)
}

// InterfaceFloat64ToUint16 interface转uint16
func InterfaceFloat64ToUint16(str interface{}) uint16 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return uint16(tmp)
}

// InterfaceFloat64ToUint32 interface转uint32
func InterfaceFloat64ToUint32(str interface{}) uint32 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return uint32(tmp)
}

// InterfaceFloat64ToUint64 interface转uint64
func InterfaceFloat64ToUint64(str interface{}) uint64 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return uint64(tmp)
}

// InterfaceFloat64ToInt8 interface转int8
func InterfaceFloat64ToInt8(str interface{}) int8 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return int8(tmp)
}

// InterfaceFloat64ToInt16 interface转int16
func InterfaceFloat64ToInt16(str interface{}) int16 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return int16(tmp)
}

//InterfaceFloat64ToInt32 interface转int32
func InterfaceFloat64ToInt32(str interface{}) int32 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return int32(tmp)
}

// InterfaceFloat64ToInt64 interface转in64
func InterfaceFloat64ToInt64(str interface{}) int64 {
	tmp, ok := str.(float64)
	if !ok {
		return 0
	}
	return int64(tmp)
}
