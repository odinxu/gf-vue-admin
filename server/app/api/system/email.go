package api

import (
	"gf-vue-admin/app/api/response"
	service "gf-vue-admin/app/service/system"
	"github.com/gogf/gf/net/ghttp"
)

var Email = new(email)

type email struct{}

func (e *email) Test(r *ghttp.Request) *response.Response {
	if err := service.Email.Test(); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "发送测试邮件失败!"}
	}
	return &response.Response{Code: 0, Message: "发送测试邮件成功!"}
}
