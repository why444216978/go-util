package string

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, end int64) (int64, string) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, end)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	if err != nil {
		panic(err)
	}
	return n, buf.String()
}

// SubstrTarget 在字符串中查找指定子串，并返回left或right部分
func SubstrTarget(str string, target string, turn string, hasPos bool) string {
	pos := strings.Index(str, target)

	if pos == -1 {
		return ""
	}

	if turn == "left" {
		if hasPos == true {
			pos = pos + 1
		}
		return str[:pos]
	} else if turn == "right" {
		if hasPos == false {
			pos = pos + 1
		}
		return str[pos:]
	} else {
		panic("params 3 error")
	}
}

// GetStringUtf8Len 获得字符串按照uft8编码的长度
func GetStringUtf8Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// JoinStringAndOther 连接字符串和其他类型
// fmt.Println(JoinStringAndOther("why", 123))
func JoinStringAndOther(val ...interface{}) string {
	return fmt.Sprint(val...)
}

// CamelToSnake 驼峰转蛇形
func CamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// SnakeToCamel 蛇形转驼峰
func SnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
