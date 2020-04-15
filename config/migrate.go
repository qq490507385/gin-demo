package config

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Self 数据库连接
var Self *gorm.DB

func init() {
	Self = InitDB()
}

// BaseModel 拥有自增id 时间
type BaseModel struct {
	ID        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
