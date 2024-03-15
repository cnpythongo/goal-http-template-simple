package main

import (
	"goal-app/cmd/api/cmd"
	_ "goal-app/docs/api"
)

// @title 前台接口文档
// @version 1.0
// @description http状态码是200，code为0时表示正常返回；code不为0时表示有业务错误。
// @BasePath /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey APIAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	cmd.Execute()
}
