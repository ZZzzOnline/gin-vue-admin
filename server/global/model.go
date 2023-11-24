package global

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type GVA_MONGO_MODEL struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // 主键ID
	CreatedAt time.Time          `bson:"CreatedAt" json:"created_at"`        // 创建时间
	UpdatedAt time.Time          `bson:"UpdatedAt" json:"updated_at"`        // 更新时间
	DeletedAt *time.Time         `bson:"DeletedAt" json:"deleted_at"`        // 删除时间
}
