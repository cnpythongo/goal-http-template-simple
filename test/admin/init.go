package admin

import (
	"github.com/cnpythongo/goal/cmd/admin/cmd"
	"github.com/cnpythongo/goal/test/utils"
	"github.com/gin-gonic/gin"
)

var testAppInst *cmd.Application

func init() {
	cfgFile := utils.ConfigFilePath
	cmd.CfgFile = &cfgFile
	testAppInst = &cmd.Application{}
	err := testAppInst.Init(nil)
	if err != nil {
		panic(err)
	}
}

func GetRouter() *gin.Engine {
	return testAppInst.GetGinEngine()
}

func GetApp() *cmd.Application {
	return testAppInst
}
