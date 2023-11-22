package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/gm"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gmReq "github.com/flipped-aurora/gin-vue-admin/server/model/gm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AdminModuleControlApi struct {
}

var adminModuleControlService = service.ServiceGroupApp.GmServiceGroup.AdminModuleControlService


// CreateAdminModuleControl 创建模块控制
// @Tags AdminModuleControl
// @Summary 创建模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "创建模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /adminModuleControl/createAdminModuleControl [post]
func (adminModuleControlApi *AdminModuleControlApi) CreateAdminModuleControl(c *gin.Context) {
	var adminModuleControl gm.AdminModuleControl
	err := c.ShouldBindJSON(&adminModuleControl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    adminModuleControl.CreatedBy = utils.GetUserID(c)
	if err := adminModuleControlService.CreateAdminModuleControl(&adminModuleControl); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAdminModuleControl 删除模块控制
// @Tags AdminModuleControl
// @Summary 删除模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "删除模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adminModuleControl/deleteAdminModuleControl [delete]
func (adminModuleControlApi *AdminModuleControlApi) DeleteAdminModuleControl(c *gin.Context) {
	var adminModuleControl gm.AdminModuleControl
	err := c.ShouldBindJSON(&adminModuleControl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    adminModuleControl.DeletedBy = utils.GetUserID(c)
	if err := adminModuleControlService.DeleteAdminModuleControl(adminModuleControl); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAdminModuleControlByIds 批量删除模块控制
// @Tags AdminModuleControl
// @Summary 批量删除模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /adminModuleControl/deleteAdminModuleControlByIds [delete]
func (adminModuleControlApi *AdminModuleControlApi) DeleteAdminModuleControlByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := adminModuleControlService.DeleteAdminModuleControlByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAdminModuleControl 更新模块控制
// @Tags AdminModuleControl
// @Summary 更新模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "更新模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adminModuleControl/updateAdminModuleControl [put]
func (adminModuleControlApi *AdminModuleControlApi) UpdateAdminModuleControl(c *gin.Context) {
	var adminModuleControl gm.AdminModuleControl
	err := c.ShouldBindJSON(&adminModuleControl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    adminModuleControl.UpdatedBy = utils.GetUserID(c)
	if err := adminModuleControlService.UpdateAdminModuleControl(adminModuleControl); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAdminModuleControl 用id查询模块控制
// @Tags AdminModuleControl
// @Summary 用id查询模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gm.AdminModuleControl true "用id查询模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adminModuleControl/findAdminModuleControl [get]
func (adminModuleControlApi *AdminModuleControlApi) FindAdminModuleControl(c *gin.Context) {
	var adminModuleControl gm.AdminModuleControl
	err := c.ShouldBindQuery(&adminModuleControl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if readminModuleControl, err := adminModuleControlService.GetAdminModuleControl(adminModuleControl.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"readminModuleControl": readminModuleControl}, c)
	}
}

// GetAdminModuleControlList 分页获取模块控制列表
// @Tags AdminModuleControl
// @Summary 分页获取模块控制列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gmReq.AdminModuleControlSearch true "分页获取模块控制列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adminModuleControl/getAdminModuleControlList [get]
func (adminModuleControlApi *AdminModuleControlApi) GetAdminModuleControlList(c *gin.Context) {
	var pageInfo gmReq.AdminModuleControlSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := adminModuleControlService.GetAdminModuleControlInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
