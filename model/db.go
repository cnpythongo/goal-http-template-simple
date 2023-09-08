package model

import (
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

type BaseModel struct {
	ID        int64            `gorm:"primary_key;comment:流水ID" json:"-"`
	CreatedAt *utils.LocalTime `gorm:"column:created_at;autoCreateTime;comment:数据创建时间" json:"-"`
	UpdatedAt *utils.LocalTime `gorm:"column:updated_at;autoUpdateTime;comment:数据更新时间" json:"-"`
	DeletedAt *utils.LocalTime `gorm:"column:deleted_at;default:null;comment:数据删除时间" json:"-"`
}

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
			_ = idb.Close()
		}
	}
	log.GetLogger().Info("Close mysql database connect done")
	return nil
}
