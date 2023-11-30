package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/gm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AdminModuleControlSearch struct{
    gm.AdminModuleControl
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
