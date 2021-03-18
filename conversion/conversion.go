package conversion

import (
	"github.com/pkg/errors"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// DeepCopy 深拷贝转换
//	 type User struct {
//		 A string
//	 }
//	 user1 := &User{A:"a"}
//	 user2 := new(User)
//	 conversion.DeepCopy(user2, user1)
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// JsonEncode interface格式化为json
func JsonEncode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// MapToJsonInt key为int的map转json
func MapToJsonInt(data map[int]interface{}) (string, error) {
	jsonStr, err := json.Marshal(data)
	return string(jsonStr), err
}

// MapToJson key为string的map转json
func MapToJson(data map[string]interface{}) (string, error) {
	jsonStr, err := json.Marshal(data)
	return string(jsonStr), err
}

// JsonToMapArray json转map数组
func JsonToMapArray(data string) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	if data == "" {
		return res, nil
	}
	err := json.Unmarshal([]byte(data), &res)

	return res, err
}

// JsonToMap json转map
func JsonToMap(data string) (map[string]interface{}, error) {
	var res map[string]interface{}
	if data == "" {
		return res, nil
	}
	err := json.Unmarshal([]byte(data), &res)

	return res, err
}

// StructToMap struct转map
func StructToMap(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	obj1 := reflect.TypeOf(obj)
	if obj1.Kind() != reflect.Struct {
		err = errors.New("type not Struct")
		return
	}
	obj2 := reflect.ValueOf(obj)
	if obj2.Kind() != reflect.Struct {
		err = errors.New("value not Struct")
		return
	}

	for i := 0; i < obj1.NumField(); i++ {
		k := obj1.Field(i).Tag.Get("json")
		if k == "" {
			k = obj1.Field(i).Name
		}
		data[k] = obj2.Field(i).Interface()
	}
	return
}

// GetStructJson 获得结构体的json切片
func GetStructJson(obj interface{}) []string {
	obj1 := reflect.TypeOf(obj)

	var data = make([]string, obj1.NumField())
	for i := 0; i < obj1.NumField(); i++ {
		k := obj1.Field(i).Tag.Get("json")
		if k == "" {
			k = obj1.Field(i).Name
		}
		data = append(data, k)
	}
	return data
}

// StructToJson 结构体转json
func StructToJson(v interface{}) (string, error) {
	jsons, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsons), nil
}

// StructToJsonByReflect 通过反射结构体转json
func StructToJsonByReflect(v interface{}) string {
	userValue := reflect.ValueOf(v)
	userType := reflect.TypeOf(v)

	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")

	num := userType.NumField()

	for i := 0; i < num; i++ {
		jsonTag := userType.Field(i).Tag.Get("json")

		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", userValue.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")

	return jsonBuilder.String()
}

// StructToMapByReflect 通过反射结构体转json
func StructToMapByReflect(v interface{}) string {
	userValue := reflect.ValueOf(v)
	userType := reflect.TypeOf(v)

	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")

	num := userType.NumField()

	for i := 0; i < num; i++ {
		jsonTag := userType.Field(i).Tag.Get("json")

		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", userValue.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")

	return jsonBuilder.String()
}

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
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return uint8(intTmp)
}

// InterfaceStringToUint16 interface转uint16
func InterfaceStringToUint16(str interface{}) uint16 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return uint16(intTmp)
}

// InterfaceStringToUint64 interface转uint64
func InterfaceStringToUint64(str interface{}) uint64 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return uint64(intTmp)
}

// InterfaceStringToInt8 interface转int8
func InterfaceStringToInt8(str interface{}) int8 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return int8(intTmp)
}

// InterfaceStringToInt16 interface转int16
func InterfaceStringToInt16(str interface{}) int16 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return int16(intTmp)
}

// InterfaceStringToInt32 interface转in32
func InterfaceStringToInt32(str interface{}) int32 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return int32(intTmp)
}

// InterfaceStringToInt64 interface转int64
func InterfaceStringToInt64(str interface{}) int64 {
	strTmp := str.(string)
	intTmp, _ := strconv.Atoi(strTmp)
	return int64(intTmp)
}

// InterfaceFloat64ToUint8 interface转uint8
func InterfaceFloat64ToUint8(str interface{}) uint8 {
	tmp := str.(float64)
	return uint8(tmp)
}

// InterfaceFloat64ToUint16 interface转uint16
func InterfaceFloat64ToUint16(str interface{}) uint16 {
	tmp := str.(float64)
	return uint16(tmp)
}

// InterfaceFloat64ToUint32 interface转uint32
func InterfaceFloat64ToUint32(str interface{}) uint32 {
	tmp := str.(float64)
	return uint32(tmp)
}

// InterfaceFloat64ToUint64 interface转uint64
func InterfaceFloat64ToUint64(str interface{}) uint64 {
	tmp := str.(float64)
	return uint64(tmp)
}

// InterfaceFloat64ToInt8 interface转int8
func InterfaceFloat64ToInt8(str interface{}) int8 {
	tmp := str.(float64)
	return int8(tmp)
}

// InterfaceFloat64ToInt16 interface转int16
func InterfaceFloat64ToInt16(str interface{}) int16 {
	tmp := str.(float64)
	return int16(tmp)
}

//InterfaceFloat64ToInt32 interface转int32
func InterfaceFloat64ToInt32(str interface{}) int32 {
	tmp := str.(float64)
	return int32(tmp)
}

// InterfaceFloat64ToInt64 interface转in64
func InterfaceFloat64ToInt64(str interface{}) int64 {
	tmp := str.(float64)
	return int64(tmp)
}
