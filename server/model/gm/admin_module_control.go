// 自动生成模板AdminModuleControl
package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// module control 结构体  AdminModuleControl
type AdminModuleControl struct {
      global.GVA_MODEL
      ForbiddenLogin  *bool `json:"forbiddenLogin" form:"forbiddenLogin" gorm:"column:forbidden_login;comment:forbidden login;"`  //forbidden login 
      ForbiddenInGameHeroExport  *bool `json:"forbiddenInGameHeroExport" form:"forbiddenInGameHeroExport" gorm:"column:forbidden_in_game_hero_export;comment:forbidden in game hero export;"`  //forbidden in game hero export 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName module control AdminModuleControl自定义表名 admin_module_control
func (AdminModuleControl) TableName() string {
  return "admin_module_control"
}

