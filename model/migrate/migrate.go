package migrate

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/pkg/log"
)

func MigrateTables(conf *config.Configuration) {
	if !conf.App.Debug { // 仅在开发模式执行migrate操作
		return
	}
	log.GetLogger().Infoln("Migrate tables start .....")

	err := model.GetDB().AutoMigrate(model.NewAccountUser())
	if err != nil {
		panic(err)
	}
	err = model.GetDB().AutoMigrate(model.NewAccountUserProfile())
	if err != nil {
		panic(err)
	}

	err = model.GetDB().AutoMigrate(model.NewAccountLoginHistory())
	if err != nil {
		panic(err)
	}

	log.GetLogger().Infoln("Migrate tables success .....")
}
