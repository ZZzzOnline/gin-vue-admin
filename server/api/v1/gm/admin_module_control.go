package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gm"
	gmReq "github.com/flipped-aurora/gin-vue-admin/server/model/gm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdminModuleControlApi struct {
}

var adminModuleControlService = service.ServiceGroupApp.GmServiceGroup.AdminModuleControlService

// CreateAdminModuleControl 创建module control
// @Tags AdminModuleControl
// @Summary 创建module control
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "创建module control"
// @Success 200 {string} string "{"success":true,"data":{},"msg":global.Translate("general.getDataSuccess")}"
// @Router /adminModuleControl/createAdminModuleControl [post]
func (adminModuleControlApi *AdminModuleControlApi) CreateAdminModuleControl(c *gin.Context) {
	var adminModuleControl gm.AdminModuleControl
	err := c.ShouldBindJSON(&adminModuleControl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	adminModuleControl.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Email": {utils.NotEmpty()},
	}
	if err := utils.Verify(adminModuleControl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := adminModuleControlService.CreateAdminModuleControl(&adminModuleControl); err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccss"), c)
	}
}

// DeleteAdminModuleControl 删除module control
// @Tags AdminModuleControl
// @Summary 删除module control
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "删除module control"
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
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
	}
}

// DeleteAdminModuleControlByIds 批量删除module control
// @Tags AdminModuleControl
// @Summary 批量删除module control
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除module control"
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
	if err := adminModuleControlService.DeleteAdminModuleControlByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
	}
}

// UpdateAdminModuleControl 更新module control
// @Tags AdminModuleControl
// @Summary 更新module control
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gm.AdminModuleControl true "更新module control"
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
	verify := utils.Rules{
		"Email": {utils.NotEmpty()},
	}
	if err := utils.Verify(adminModuleControl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := adminModuleControlService.UpdateAdminModuleControl(adminModuleControl); err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("general.updateSuccess"), c)
	}
}

// FindAdminModuleControl 用id查询module control
// @Tags AdminModuleControl
// @Summary 用id查询module control
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gm.AdminModuleControl true "用id查询module control"
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
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
	} else {
		response.OkWithData(gin.H{"readminModuleControl": readminModuleControl}, c)
	}
}

// GetAdminModuleControlList 分页获取module control列表
// @Tags AdminModuleControl
// @Summary 分页获取module control列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gmReq.AdminModuleControlSearch true "分页获取module control列表"
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
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, global.Translate("general.getDataSuccess"), c)
	}
}
