// 自动生成模板AdminModuleControl
package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 模块控制 结构体  AdminModuleControl
type AdminModuleControl struct {
      global.GVA_MODEL
      ForbiddenLogin  *bool `json:"forbiddenLogin" form:"forbiddenLogin" gorm:"column:forbidden_login;comment:玩家是否被禁止登录：禁止登录-true;未禁止登录-false/空;"`  //禁止登录 
      ForbiddenInGameHeroExport  *bool `json:"forbiddenInGameHeroExport" form:"forbiddenInGameHeroExport" gorm:"column:forbidden_in_game_hero_export;comment:是否禁止游戏内英雄导出：禁止-true;不禁止-其它;"`  //禁止游戏内英雄导出 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 模块控制 AdminModuleControl自定义表名 admin_module_control
func (AdminModuleControl) TableName() string {
  return "admin_module_control"
}

