// Package router 路由
package router

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/pkg/app"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/server/user"
	"github.com/thoohv5/template/pkg/hpx"
)

// ProviderSet is app providers.
var ProviderSet = wire.NewSet(
	New,
)

// New 创建
func New(cf config.IConfig, server user.IServer) hpx.RegisterRouter {
	return func(r *gin.Engine, trans ut.Translator) error {

		if cf.GetHttp().SwaggerEnabled {
			app.InitSwagRouter(r, cf.GetHttp().LocalAddr)
		}

		u := r.Group("/user")

		// 注册
		u.POST("/register", server.Register)
		// 详情
		u.GET("/detail", server.Detail)
		// 列表
		u.GET("/list", server.List)
		// 编辑
		u.POST("/edit", server.Edit)

		return nil
	}
}
