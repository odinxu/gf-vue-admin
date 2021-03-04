package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"gf-vue-admin/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type email struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewEmailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &email{router: router, response: &response.Handler{}}
}

func (e *email) Init() {
	group := e.router.Group("/email").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("emailTest", e.response.Handler()(api.Email.Test)) // 发送测试邮件
	}
}
