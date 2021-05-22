package maps

import (
	"fmt"
	"sort"
)

// MergeInterface 合并interface类型的map
func MergeInterface(list ...map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for _, val := range list {
		for k, v := range val {
			ret[k] = v
		}
	}
	return ret
}

// MergeString 合并string类型的map
func MergeString(list ...map[string]string) map[string]string {
	ret := make(map[string]string)
	for _, val := range list {
		for k, v := range val {
			ret[k] = v
		}
	}
	return ret
}

// MergeInt 合并int类型的map
func MergeInt(list ...map[string]int) map[string]int {
	ret := make(map[string]int)
	for _, val := range list {
		for k, v := range val {
			ret[k] = v
		}
	}
	return ret
}

// KeySort map按照key排序，解决map随机迭代问题
func KeySort(m map[string]interface{}) []map[string]string {
	ret := make([]map[string]string, len(m))

	kvMap := make(map[string]string)
	sortedKeys := make([]string, 0)

	for k, v := range m {
		kvMap[k] = fmt.Sprintf("%v", v)
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for idx, key := range sortedKeys {
		ret[idx] = map[string]string{
			"key": key,
			"val": kvMap[key],
		}
	}

	return ret
}

// GetKeyList 获得map的key列表
func GetKeyList(m map[string]interface{}) []string {
	list := make([]string, 0)
	for k, _ := range m {
		list = append(list, k)
	}
	return list
}

// GetStringValueList 获得字符串map值列表
func GetStringValueList(m map[string]string) []string {
	list := make([]string, 0)
	for k, _ := range m {
		list = append(list, k)
	}
	return list
}

// GetMapList 获得map某个key对应值列表
func GetMapList(m []map[string]interface{}, fields string) []interface{} {
	list := make([]interface{}, 0)
	for _, v := range m {
		list = append(list, v[fields])
	}
	return list
}
