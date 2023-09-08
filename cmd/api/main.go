package main

import (
	"github.com/cnpythongo/goal/cmd/api/cmd"
	_ "github.com/cnpythongo/goal/docs/api"
)

// @title 前端应用接口文档
// @version 1.0

// @host localhost:8100
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
