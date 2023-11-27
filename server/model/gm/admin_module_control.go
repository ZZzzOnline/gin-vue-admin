// 自动生成模板AdminModuleControl
package gm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 模块控制 结构体  AdminModuleControl
type AdminModuleControl struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`                                                                                                                                          // 主键ID
	CreatedAt                 time.Time          `bson:"CreatedAt" json:"CreatedAt"`                                                                                                                                                  // 创建时间
	UpdatedAt                 time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`                                                                                                                                                  // 更新时间
	DeletedAt                 *time.Time         `bson:"DeletedAt" json:"DeletedAt"`                                                                                                                                                  // 删除时间
	Email                     string             `bson:"Email" json:"email" form:"email" gorm:"column:email;comment:邮箱地址;"`                                                                                                           //邮箱地址
	AccountId                 *uint              `bson:"AccountId" json:"accountId" form:"accountId" gorm:"column:account_id;comment:账号唯一标识;"`                                                                                        //账号唯一标识
	PlayerId                  string             `bson:"PlayerId" json:"playerId" form:"playerId" gorm:"column:player_id;comment:玩家标识;"`                                                                                              //玩家标识
	ForbiddenLogin            *bool              `bson:"ForbiddenLogin" json:"forbiddenLogin" form:"forbiddenLogin" gorm:"column:forbidden_login;comment:玩家是否被禁止登录：禁止登录-true;未禁止登录-false/空;"`                                         //禁止登录
	ForbiddenInGameHeroExport *bool              `bson:"ForbiddenInGameHeroExport" json:"forbiddenInGameHeroExport" form:"forbiddenInGameHeroExport" gorm:"column:forbidden_in_game_hero_export;comment:是否禁止游戏内英雄导出：禁止-true;不禁止-其它;"` //禁止游戏内英雄导出
	CreatedBy                 uint               `bson:"CreatedBy" gorm:"column:created_by;comment:创建者"`
	UpdatedBy                 uint               `bson:"UpdatedBy" gorm:"column:updated_by;comment:更新者"`
	DeletedBy                 uint               `bson:"DeletedBy" gorm:"column:deleted_by;comment:删除者"`
}

// TableName 模块控制 AdminModuleControl自定义表名 admin_module_control
func (AdminModuleControl) TableName() string {
	return "admin_module_control"
}
