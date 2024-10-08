package migrate

import (
	"goal-app/model"
	"goal-app/pkg/config"
	"goal-app/pkg/log"
)

func MigrateTables(conf *config.Configuration) {
	if !conf.App.Debug { // 仅在开发模式执行migrate操作
		return
	}
	log.GetLogger().Infoln("Migrate tables start .....")

	db := model.GetDB()
	err := db.AutoMigrate(
		model.NewUser(),
		model.NewUserProfile(),
		model.NewSystemConfig(),
		model.NewAttachment(),
		model.NewSystemOrg(),
		model.NewSystemMenu(),
		model.NewGenTable(),
		model.NewGenTableColumn(),
		model.NewSystemLog(),
	)
	if err != nil {
		panic(err)
	}

	log.GetLogger().Infoln("Migrate tables success .....")
}
