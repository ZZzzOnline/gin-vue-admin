package gm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdminModuleControlRouter struct {
}

// InitAdminModuleControlRouter 初始化 模块控制 路由信息
func (s *AdminModuleControlRouter) InitAdminModuleControlRouter(Router *gin.RouterGroup) {
	adminModuleControlRouter := Router.Group("adminModuleControl").Use(middleware.OperationRecord())
	adminModuleControlRouterWithoutRecord := Router.Group("adminModuleControl")
	var adminModuleControlApi = v1.ApiGroupApp.GmApiGroup.AdminModuleControlApi
	{
		adminModuleControlRouter.POST("createAdminModuleControl", adminModuleControlApi.CreateAdminModuleControl)   // 新建模块控制
		adminModuleControlRouter.DELETE("deleteAdminModuleControl", adminModuleControlApi.DeleteAdminModuleControl) // 删除模块控制
		adminModuleControlRouter.DELETE("deleteAdminModuleControlByIds", adminModuleControlApi.DeleteAdminModuleControlByIds) // 批量删除模块控制
		adminModuleControlRouter.PUT("updateAdminModuleControl", adminModuleControlApi.UpdateAdminModuleControl)    // 更新模块控制
	}
	{
		adminModuleControlRouterWithoutRecord.GET("findAdminModuleControl", adminModuleControlApi.FindAdminModuleControl)        // 根据ID获取模块控制
		adminModuleControlRouterWithoutRecord.GET("getAdminModuleControlList", adminModuleControlApi.GetAdminModuleControlList)  // 获取模块控制列表
	}
}
