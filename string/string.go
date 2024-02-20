package string

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, length int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, length)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

// SubStrTarget 在字符串中查找指定子串，并返回left或right部分
func SubStrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos == true {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if hasPos == false {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
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
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
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

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

var splitRegexp = regexp.MustCompile(`([\n。！!？?.])`)

func SplitPunctuation(content string, length int) []string {
	arr := splitRegexp.Split(content, -1)
	submatch := splitRegexp.FindAllStringSubmatch(content, -1)

	contents := []string{}
	for i, s := range submatch {
		contents = append(contents, arr[i]+s[0])
	}

	result := []string{}
	current := ""
	for _, s := range contents {
		if current == "" {
			current = s
			continue
		}
		tmp := current + s
		if utf8.RuneCountInString(tmp) > length {
			result = append(result, current)
			current = s
			continue
		}
		current = tmp
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}

func ReverseSplitPunctuation(content string, length int) []string {
	arr := splitRegexp.Split(content, -1)
	submatch := splitRegexp.FindAllStringSubmatch(content, -1)

	contents := []string{}
	for i, s := range submatch {
		contents = append(contents, arr[i]+s[0])
	}

	l := list.New()
	current := ""
	for i := len(contents) - 1; i >= 0; i-- {
		s := contents[i]
		if current == "" {
			current = s
			continue
		}
		tmp := s + current
		if utf8.RuneCountInString(tmp) > length {
			l.PushFront(current)
			current = s
			continue
		}
		current = tmp
	}

	if current != "" {
		l.PushFront(current)
	}

	result := []string{}
	for i := l.Front(); i != nil; i = i.Next() {
		result = append(result, i.Value.(string))
	}

	return result
}
