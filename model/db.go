package model

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	"github.com/cnpythongo/goal/pkg/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		panic("DB not inited")
	}
	return db
}

func getDatabaseDsn(cfg *config.MysqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)", cfg.User, cfg.Password, cfg.Host, cfg.Port,
	)
}

func initDatabase(cfg *config.MysqlConfig) {
	dsn := fmt.Sprintf("%s/", getDatabaseDsn(cfg))
	db, err := gorm.Open(mysql.Open(dsn), nil)

	if err != nil {
		panic("Connect Database error >>> " + err.Error())
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';",
		cfg.DbName,
	)
	err = db.Exec(createSQL).Error
	if err != nil {
		panic("Create Database error >>> " + err.Error())
	}
}

func Init(conf *config.MysqlConfig) error {
	initDatabase(conf)

	var err error
	dsn := fmt.Sprintf("%s/%s?%s",
		getDatabaseDsn(conf), conf.DbName, conf.DbParams,
	)
	db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		return err
	}
	idb, _ := db.DB()
	idb.SetConnMaxIdleTime(120 * time.Second)
	idb.SetConnMaxLifetime(7200 * time.Second)
	idb.SetMaxOpenConns(200)
	idb.SetMaxIdleConns(10)
	if err := idb.Ping(); err != nil {
		return err
	}
	return nil
}

func Close() error {
	if db != nil {
		idb, err := db.DB()
		if err == nil {
			idb.Close()
		}
	}
	fmt.Println("Close mysql database connect done")
	return nil
}
