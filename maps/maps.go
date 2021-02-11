package maps

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
