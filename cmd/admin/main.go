package main

import (
	"github.com/cnpythongo/goal/cmd/admin/cmd"
	_ "github.com/cnpythongo/goal/docs/admin"
)

// @title 后台管理系统接口文档
// @version 1.0
//
// @description http状态码是200，code为0时表示正常返回；code不为0时表示有业务错误。返回的JSON数据结构如下：
// @BasePath /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	cmd.Execute()
}
