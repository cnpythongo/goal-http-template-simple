package migrate

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/cnpythongo/goal/pkg/common/log"
)

func MigrateTables(conf *config.Configuration) {
	if !conf.App.Debug { // 仅在开发模式执行migrate操作
		return
	}
	log.GetLogger().Infoln("migrate tables start .....")
	err := model.GetDB().AutoMigrate(model.NewUser())
	if err != nil {
		panic(err)
	}
	err = model.GetDB().AutoMigrate(model.NewUserProfile())
	if err != nil {
		panic(err)
	}
	err = model.GetDB().AutoMigrate(model.NewLoginHistory())
	if err != nil {
		panic(err)
	}
	log.GetLogger().Infoln("migrate tables success .....")
}
