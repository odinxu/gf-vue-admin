package router

import (
	api "gf-vue-admin/app/api/extra"
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type excel struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewExcelRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &excel{router: router, response: &response.Handler{}}
}

func (e *excel) Init() {
	group := e.router.Group("/excel")
	{
		group.GET("/loadExcel", e.response.Handler()(api.Excel.Load))            // 加载Excel数据
		group.POST("/importExcel", e.response.Handler()(api.Excel.Import))       // 导入Excel
		group.POST("/exportExcel", e.response.Handler()(api.Excel.Export))       // 导出Excel
		group.GET("/downloadTemplate", e.response.Handler()(api.Excel.Download)) // 下载模板文件
	}
}
