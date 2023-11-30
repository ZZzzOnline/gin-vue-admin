package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gmReq "github.com/flipped-aurora/gin-vue-admin/server/model/gm/request"
    "gorm.io/gorm"
)

type AdminModuleControlService struct {
}

// CreateAdminModuleControl 创建module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) CreateAdminModuleControl(adminModuleControl *gm.AdminModuleControl) (err error) {
	err = global.GVA_DB.Create(adminModuleControl).Error
	return err
}

// DeleteAdminModuleControl 删除module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService)DeleteAdminModuleControl(adminModuleControl gm.AdminModuleControl) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gm.AdminModuleControl{}).Where("id = ?", adminModuleControl.ID).Update("deleted_by", adminModuleControl.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&adminModuleControl).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAdminModuleControlByIds 批量删除module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService)DeleteAdminModuleControlByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gm.AdminModuleControl{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gm.AdminModuleControl{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAdminModuleControl 更新module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService)UpdateAdminModuleControl(adminModuleControl gm.AdminModuleControl) (err error) {
	err = global.GVA_DB.Save(&adminModuleControl).Error
	return err
}

// GetAdminModuleControl 根据id获取module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService)GetAdminModuleControl(id uint) (adminModuleControl gm.AdminModuleControl, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&adminModuleControl).Error
	return
}

// GetAdminModuleControlInfoList 分页获取module control记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService)GetAdminModuleControlInfoList(info gmReq.AdminModuleControlSearch) (list []gm.AdminModuleControl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gm.AdminModuleControl{})
    var adminModuleControls []gm.AdminModuleControl
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ForbiddenLogin != nil {
        db = db.Where("forbidden_login = ?",info.ForbiddenLogin)
    }
    if info.ForbiddenInGameHeroExport != nil {
        db = db.Where("forbidden_in_game_hero_export = ?",info.ForbiddenInGameHeroExport)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&adminModuleControls).Error
	return  adminModuleControls, total, err
}
