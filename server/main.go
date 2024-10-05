package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.4
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

/*
package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // 模拟用户身份验证中间件
    router.Use(func(c *gin.Context) {
        // 假设通过请求头中的 Token 进行身份验证
        token := c.GetHeader("Authorization")
        if token != "valid-token" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Next()
    })

    // 动态路由处理图片请求
    router.GET("/images/:imageID", func(c *gin.Context) {
        imageID := c.Param("imageID")
        // 根据 imageID 和用户身份验证结果返回图片数据
        // 这里假设所有用户都可以访问所有图片
        c.File("./public/images/" + imageID + ".jpg")
    })

    // API 路由
    router.GET("/api/images", func(c *gin.Context) {
        images := []map[string]string{
            {"url": "/images/image1", "description": "Image 1", "hash": "abc123"},
            {"url": "/images/image2", "description": "Image 2", "hash": "def456"},
            // 更多图片信息...
        }
        c.JSON(http.StatusOK, images)
    })

    // 启动服务器
    router.Run(":8080")
}






package main

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/events", func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "text/event-stream")
        c.Writer.Header().Set("Cache-Control", "no-cache")
        c.Writer.Header().Set("Connection", "keep-alive")

        ticker := time.NewTicker(2 * time.Second)
        defer ticker.Stop()

        for {
            select {
            case t := <-ticker.C:
                c.SSEvent("message", t.Format("2006-01-02 15:04:05"))
                c.Writer.Flush()
            case <-c.Request.Context().Done():
                return
            }
        }
    })
    r.Run(":8080")
}

*/
