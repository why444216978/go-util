package conversion

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	util_err "github.com/why444216978/go-util/error"
)

// DeepCopy 深拷贝转换
// type User struct {
// 	A string
// }
// user1 := &User{A:"a"}
// user2 := new(User)
// conversion.DeepCopy(user2, user1)
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// JsonEncode interface转json
func JsonEncode(v interface{}) string {
	b, err := json.Marshal(v)
	util_err.Must(err)
	return string(b)
}

// MapToJsonInt 索引为int的map转json
func MapToJsonInt(data map[int]interface{}) string {
	jsonStr, err := json.Marshal(data)
	util_err.Must(err)
	return string(jsonStr)
}

// MapToJson 索引为string的map转json
func MapToJson(data map[string]interface{}) string {
	jsonStr, err := json.Marshal(data)
	util_err.Must(err)
	return string(jsonStr)
}

// JsonToMapArray json转map数组
func JsonToMapArray(data string) []map[string]interface{} {
	var res []map[string]interface{}
	if data == "" {
		return res
	}
	err := json.Unmarshal([]byte(data), &res)
	util_err.Must(err)

	return res
}

// JsonToMap json转map
func JsonToMap(data string) map[string]interface{} {
	var res map[string]interface{}
	if data == "" {
		return res
	}
	err := json.Unmarshal([]byte(data), &res)
	util_err.Must(err)
	return res
}

// StructToMap 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[strings.ToLower(obj1.Field(i).Name)] = obj2.Field(i).Interface()
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
// 解决强转效率低：先分配一个内存再复制内容
// StringToByte("str"
func StringToByte(str string) []byte {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	sliceHeader.Cap = sliceHeader.Len
	b := *(*[]byte)(unsafe.Pointer(sliceHeader))
	return b
}

// ByteToString 自定义byte转string
// 解决强转效率低：先分配一个内存再复制内容
// ByteToString([]byte("why"))
func ByteToString(b []byte) string {
	ptr := (*string)(unsafe.Pointer(&b))
	return *ptr
}

//json转struct
//jsonStr := `{"name":"why","age":18}`
//user := User{}
//json.Unmarshal([]byte(jsonStr), &user)
//fmt.Println(user)
