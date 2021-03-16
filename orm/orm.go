package orm

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// Count 数量
func Count(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}) (count int64, err error) {
	db = db.Model(model)

	for k, v := range where {
		db.Where(k, v)
	}

	err = db.Count(&count).Error

	return
}

// Select 查询
func Select(ctx context.Context, db *gorm.DB, model interface{}, fields string, where map[string]interface{}, group []string, groupData interface{}, having map[string]interface{}, order []string, start, limit int) (err error) {
	db = db.Model(model).Select(fields)

	for k, v := range where {
		db.Where(k, v)
	}

	for _, v := range group {
		db.Group(v)
	}

	for k, v := range having {
		db.Having(k, v)
	}

	for _, v := range order {
		db.Order(v)
	}

	db = db.Offset(start).Limit(limit)

	if groupData != nil {
		err = db.Find(groupData).Error
	} else {
		err = db.Find(model).Error
	}

	return
}

// Update 更新
func Update(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}, updates map[string]interface{}) (int64, error) {
	db = db.Model(model)

	for k, v := range where {
		db.Where(k, v)
	}

	res := db.Updates(updates)

	return res.RowsAffected, res.Error
}

// Insert 插入
func Insert(ctx context.Context, db *gorm.DB, model interface{}) (int64, error) {
	res := db.Create(model)

	return res.RowsAffected, res.Error
}

// Delete 删除
func Delete(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}) (int64, error) {
	for k, v := range where {
		db.Where(k, v)
	}

	res := db.Delete(model)

	return res.RowsAffected, res.Error
}

// FormatEqString 格式化string等于
func FormatEqString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s = ?", k)] = v
	}
	return ret
}

// FormatEqInt 格式化int等于
func FormatEqInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s = ?", k)] = v
	}
	return ret
}

// FormatGtString 格式化string大于
func FormatGtString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s > ?", k)] = v
	}
	return ret
}

// FormatGtInt 格式化int大于
func FormatGtInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s > ?", k)] = v
	}
	return ret
}

// FormatLtString 格式化string小于
func FormatLtString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s < ?", k)] = v
	}
	return ret
}

// FormatLtInt 格式化int小于
func FormatLtInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s < ?", k)] = v
	}
	return ret
}

// FormatGteString 格式化string大于等于
func FormatGteString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s >= ?", k)] = v
	}
	return ret
}

// FormatGteInt 格式化int大于等于
func FormatGteInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s >=", k)] = v
	}
	return ret
}

// FormatLteString 格式化string小于等于
func FormatLteString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s <=", k)] = v
	}
	return ret
}

// FormatLteInt 格式化int小于等于
func FormatLteInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s <=", k)] = v
	}
	return ret
}

// FormatLike 格式化like
func FormatLike(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s Like ?", k)] = v
	}
	return ret
}

// FormatInString 格式化string切片类型in条件
func FormatInString(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s in (?)", k)] = v
	}
	return ret
}

// FormatInInt 格式化int切片类型in条件
func FormatInInt(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s in (?)", k)] = v
	}
	return ret
}
