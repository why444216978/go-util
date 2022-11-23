package orm

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var ErrDBNil = errors.New("db is nil")

// Count 数量
func Count(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}) (count int64, err error) {
	if db == nil {
		err = ErrDBNil
		return
	}

	db = db.WithContext(ctx).Model(model)

	for k, v := range where {
		db = db.Where(k, v)
	}

	err = db.Count(&count).Error

	return
}

// Select 查询
func Select(ctx context.Context, db *gorm.DB, model interface{}, fields string, where map[string]interface{}, group []string, groupData interface{}, having map[string]interface{}, order []string, start, limit int) (err error) {
	if db == nil {
		err = ErrDBNil
		return
	}

	db = db.WithContext(ctx).Model(model).Select(fields)

	for k, v := range where {
		db = db.Where(k, v)
	}

	for _, v := range group {
		db = db.Group(v)
	}

	for k, v := range having {
		db = db.Having(k, v)
	}

	for _, v := range order {
		db = db.Order(v)
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
func Update(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}, updates map[string]interface{}) (affectRows int64, err error) {
	if db == nil {
		return 0, ErrDBNil
	}

	db = db.WithContext(ctx).Model(model)

	for k, v := range where {
		db = db.Where(k, v)
	}

	res := db.Updates(updates)

	return res.RowsAffected, res.Error
}

// Insert 插入
func Insert(ctx context.Context, db *gorm.DB, model interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDBNil
	}

	res := db.WithContext(ctx).Create(model)

	return res.RowsAffected, res.Error
}

// Delete 删除
func Delete(ctx context.Context, db *gorm.DB, model interface{}, where map[string]interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDBNil
	}

	db = db.WithContext(ctx)

	for k, v := range where {
		db = db.Where(k, v)
	}

	res := db.Delete(model)

	return res.RowsAffected, res.Error
}

func WithWhere(ctx context.Context, db *gorm.DB, where map[string]interface{}) *gorm.DB {
	if db == nil {
		return nil
	}

	db = db.WithContext(ctx)

	for k, v := range where {
		db = db.Where(k, v)
	}

	return db
}

func WithOrder(ctx context.Context, db *gorm.DB, order []string) *gorm.DB {
	if db == nil {
		return nil
	}

	db = db.WithContext(ctx)

	for k, v := range order {
		db = db.Having(k, v)
	}

	return db
}

func WithHaving(ctx context.Context, db *gorm.DB, having map[string]interface{}) *gorm.DB {
	if db == nil {
		return nil
	}

	db = db.WithContext(ctx)

	for k, v := range having {
		db = db.Having(k, v)
	}

	return db
}

// FormatEq
func FormatEq(v interface{}) string {
	return fmt.Sprintf("%s = ?", v)
}

// FormatEqList
func FormatEqList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s = ?", k)] = v
	}
	return ret
}

// FormatGt
func FormatGt(v interface{}) string {
	return fmt.Sprintf("%s > ?", v)
}

// FormatGtList
func FormatGtList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s > ?", k)] = v
	}
	return ret
}

// FormatLt
func FormatLt(v interface{}) string {
	return fmt.Sprintf("%s < ?", v)
}

// FormatLtList
func FormatLtList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s < ?", k)] = v
	}
	return ret
}

// FormatGte
func FormatGte(v interface{}) string {
	return fmt.Sprintf("%s >= ?", v)
}

// FormatGteList
func FormatGteList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s >= ?", k)] = v
	}
	return ret
}

// FormatLte
func FormatLte(v interface{}) string {
	return fmt.Sprintf("%s <= ?", v)
}

// FormatLteList
func FormatLteList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s <=", k)] = v
	}
	return ret
}

// FormatLike
func FormatLike(v interface{}) string {
	return fmt.Sprintf("%s like ?", v)
}

// FormatLikeList
func FormatLikeList(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s Like ?", k)] = v
	}
	return ret
}

// FormatIn
func FormatIn(v interface{}) string {
	return fmt.Sprintf("%s in (?)", v)
}

// FormatInList
func FormatInList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s in (?)", k)] = v
	}
	return ret
}

// FormatNotIn
func FormatNotIn(v interface{}) string {
	return fmt.Sprintf("%s not in (?)", v)
}

// FormatNotInList
func FormatNotInList(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s not in (?)", k)] = v
	}
	return ret
}

// ExtractError extract gorm error to judge db error
func ExtractError(err error) (mysqlErr *mysql.MySQLError, ok bool) {
	mysqlErr = &mysql.MySQLError{}
	ok = errors.As(err, &mysqlErr)
	return
}

func CheckDuplicate(err error) bool {
	e, ok := ExtractError(err)
	if ok && e.Number == uint16(1062) {
		return true
	}
	return false
}
