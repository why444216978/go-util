package page

import "math"

// Page 分页方法
func Page(list []map[string]interface{}, pageNum, pageSize, total int) map[string]interface{} {
	page := make(map[string]interface{})
	page["has_next"] = true
	page["current"] = pageNum
	page["next"] = pageSize + 1

	count := len(list)
	if count <= pageSize {
		page["has_next"] = false
	}
	page["total"] = math.Ceil(float64(count) / float64(pageSize))
	page["total_rows"] = total

	return page
}

// Limit 用于解决分页slice out of range 问题
func Limit(list []map[string]interface{}, pageNum, pageSize int) []map[string]interface{} {
	retList := make([]map[string]interface{}, 0, pageNum)

	if list != nil {
		last := pageSize
		count := len(list)
		if count < pageSize {
			last = count
		}
		retList = list[:last]
	}

	return retList
}
