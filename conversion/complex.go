package conversion

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	json "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

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

// JsonToMapArray json转map数组
func JsonToMapArray(data string) ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0)
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

// ReaderToStruct 解析reader到结构体
func ReaderToStruct(reader io.Reader, val interface{}) error {
	return json.NewDecoder(reader).Decode(val)
}

// StructToMap struct转map
func StructToMap(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	if obj == nil {
		err = errors.New("obj is nil")
		return
	}

	obj1 := reflect.TypeOf(obj)
	if obj1.Kind() != reflect.Struct {
		err = errors.New("type not Struct")
		return
	}
	obj2 := reflect.ValueOf(obj)

	for i := 0; i < obj1.NumField(); i++ {
		k := obj1.Field(i).Tag.Get("json")
		if k == "" {
			k = obj1.Field(i).Name
		}
		data[k] = obj2.Field(i).Interface()
	}
	return
}

// StructToJson 结构体转json
func StructToJson(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	obj := reflect.ValueOf(v)
	if obj.Kind() != reflect.Struct {
		return "", errors.New("value not Struct")
	}

	jsons, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsons), nil
}

// StructToJsonByReflect 通过反射结构体转json
func StructToJsonByReflect(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	userValue := reflect.ValueOf(v)
	if userValue.Kind() != reflect.Struct {
		return "", errors.New("type not Struct")
	}
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

	return jsonBuilder.String(), nil
}
