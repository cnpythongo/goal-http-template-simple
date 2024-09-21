package model

import (
	"fmt"
	"goal-app/pkg/config"
	"goal-app/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB
var dbWrite *gorm.DB

type BaseModel struct {
	ID         uint64 `gorm:"primary_key;comment:流水ID" json:"-"`
	CreateTime uint64 `gorm:"column:create_time;autoCreateTime;not null;comment:数据创建时间" json:"-"`
	UpdateTime uint64 `gorm:"column:update_time;autoUpdateTime;not null;comment:数据更新时间" json:"-"`
	DeleteTime uint64 `gorm:"column:delete_time;default:0;comment:数据删除时间" json:"-"`
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("DB Reader not inited")
	}
	return db
}

func GetDBWrite() *gorm.DB {
	if dbWrite == nil {
		panic("DBWrite not inited")
	}
	return dbWrite
}

func getDatabaseDsn(cfg *config.MysqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)", cfg.User, cfg.Password, cfg.Host, cfg.Port,
	)
}

func initDatabase(cfg *config.MysqlConfig) {
	var err error
	dsn := fmt.Sprintf("%s/", getDatabaseDsn(cfg))
	db, err = gorm.Open(mysql.Open(dsn), nil)

	if err != nil {
		panic("Connect Database error >>> " + err.Error())
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';",
		cfg.DbName,
	)
	err = db.Exec(createSQL).Error
	if err != nil {
		panic("CreateDatabase error: " + err.Error())
	}
}

func Init(conf *config.MysqlConfig) error {
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
	if err = idb.Ping(); err != nil {
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

	if dbWrite != nil {
		wdb, err := dbWrite.DB()
		if err == nil {
			_ = wdb.Close()
		}
	}
	log.GetLogger().Info("Close mysql write database connect done")
	return nil
}

func InitWrite(conf *config.MysqlConfig) error {
	initDatabase(conf)

	var err error
	dsn := fmt.Sprintf("%s/%s?%s",
		getDatabaseDsn(conf), conf.DbName, conf.DbParams,
	)
	dbWrite, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		return err
	}
	idb, _ := dbWrite.DB()
	idb.SetConnMaxIdleTime(120 * time.Second)
	idb.SetConnMaxLifetime(7200 * time.Second)
	idb.SetMaxOpenConns(200)
	idb.SetMaxIdleConns(10)
	if err = idb.Ping(); err != nil {
		return err
	}
	return nil
}
