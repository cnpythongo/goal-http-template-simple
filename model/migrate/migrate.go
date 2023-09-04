package migrate

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/service/account"
)

func MigrateTables(conf *config.Configuration) {
	if !conf.App.Debug { // 仅在开发模式执行migrate操作
		return
	}
	log.GetLogger().Infoln("Migrate tables start .....")

	err := model.GetDB().AutoMigrate(account.NewUser())
	if err != nil {
		panic(err)
	}
	err = model.GetDB().AutoMigrate(account.NewUserProfile())
	if err != nil {
		panic(err)
	}

	err = model.GetDB().AutoMigrate(account.NewLoginHistory())
	if err != nil {
		panic(err)
	}

	log.GetLogger().Infoln("Migrate tables success .....")
}
