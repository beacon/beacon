package dao

import "github.com/jinzhu/gorm"

// Options for database
type Options struct {
}

// TimeField 通用时间字段
type TimeField struct {
	CreatedAt int64
	UpdatedAt int64
}

// BeforeCreate gorm 钩子方法，设置CreateTime和ModifyTime
func (t *TimeField) BeforeCreate() error {
	now := gorm.NowFunc()
	t.CreatedAt = now.Unix()
	t.UpdatedAt = now.Unix()
	return nil
}

// BeforeUpdate gorm 钩子方法，设置ModifyTime
func (t *TimeField) BeforeUpdate() error {
	now := gorm.NowFunc()
	t.UpdatedAt = now.Unix()
	return nil
}
