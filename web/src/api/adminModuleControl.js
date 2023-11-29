import service from '@/utils/request'

// @Tags AdminModuleControl
// @Summary 创建模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdminModuleControl true "创建模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /adminModuleControl/createAdminModuleControl [post]
export const createAdminModuleControl = (data) => {
  return service({
    url: '/adminModuleControl/createAdminModuleControl',
    method: 'post',
    data
  })
}

// @Tags AdminModuleControl
// @Summary 删除模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdminModuleControl true "删除模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adminModuleControl/deleteAdminModuleControl [delete]
export const deleteAdminModuleControl = (data) => {
  return service({
    url: '/adminModuleControl/deleteAdminModuleControl',
    method: 'delete',
    data
  })
}

// @Tags AdminModuleControl
// @Summary 批量删除模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adminModuleControl/deleteAdminModuleControl [delete]
export const deleteAdminModuleControlByIds = (data) => {
  return service({
    url: '/adminModuleControl/deleteAdminModuleControlByIds',
    method: 'delete',
    data
  })
}

// @Tags AdminModuleControl
// @Summary 更新模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdminModuleControl true "更新模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adminModuleControl/updateAdminModuleControl [put]
export const updateAdminModuleControl = (data) => {
  return service({
    url: '/adminModuleControl/updateAdminModuleControl',
    method: 'put',
    data
  })
}

// @Tags AdminModuleControl
// @Summary 用id查询模块控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AdminModuleControl true "用id查询模块控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adminModuleControl/findAdminModuleControl [get]
export const findAdminModuleControl = (params) => {
  return service({
    url: '/adminModuleControl/findAdminModuleControl',
    method: 'get',
    params
  })
}

// @Tags AdminModuleControl
// @Summary 分页获取模块控制列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取模块控制列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adminModuleControl/getAdminModuleControlList [get]
export const getAdminModuleControlList = (params) => {
  return service({
    url: '/adminModuleControl/getAdminModuleControlList',
    method: 'get',
    params
  })
}
