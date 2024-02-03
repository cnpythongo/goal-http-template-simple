package admin

import (
	"github.com/gin-gonic/gin"
	"goal-app/cmd/admin/cmd"
	"goal-app/test/utils"
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
