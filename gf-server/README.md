gf-vue-admin

## 开发计划

- [x] gf-vue-admin服务端
  - [x] 基于[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)的server与[goframe](https://goframe.org/start/index)的官方推荐架构
  - [x] 架构具体看文档
- [x] 路由层
- [x] jwt鉴权
- [x] jwt黑名单
- [ ] 管理员用户模块
- [x] 基础功模块
- [ ] menu模块
- [ ] 角色管理模块
- [ ] 功能api模块
- [ ] 文件上传下载功能模块
- [ ] 工作流
- [ ] casbin权限模块
- [ ] system模块
- [ ] 客户模块
- [ ] 自动化代码模块
- [x] 字典详情管理模块
- [x] 字典管理模块

## 注意问题
- Mysql的tinyint(1) 使用gf gen model 生成model文件夹里面的Entity.go文件中的字段类型是int而不是bool