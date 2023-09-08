package main

import (
	"github.com/cnpythongo/goal/cmd/admin/cmd"
	_ "github.com/cnpythongo/goal/docs/admin"
)

// @title 后台管理系统接口文档
// @version 1.0

// @host localhost
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
