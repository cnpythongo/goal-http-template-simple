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

	db := model.GetDB()
	err := db.AutoMigrate(model.NewUser())
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(model.NewUserProfile())
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(model.NewHistory())
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(model.NewConfig())
	if err != nil {
		panic(err)
	}

	log.GetLogger().Infoln("Migrate tables success .....")
}
