package models

import (
	"go-blog/common/models"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint          `gorm:"primary_key" json:"id"`
	CreatedAt models.XTime  `json:"created_at"`
	UpdatedAt models.XTime  `json:"updated_at"`
	DeletedAt *models.XTime `sql:"index"`
}

// Paginate 分页封装
func Paginate(index int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if index == 0 {
			index = 1
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}
		offset := (index - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

// GetCount 统计总数量
func GetCount(query *gorm.DB) (total int64, err error) {
	err = query.Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return
}
