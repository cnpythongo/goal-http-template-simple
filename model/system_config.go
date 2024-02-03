package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	BaseModel
	Scope   SystemConfigScope `json:"scope" gorm:"column:scope;type:varchar(50);not null;default:'';comment:作用域,global-全局,admin-管理后台,app-前台应用"`
	Name    string            `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:配置名称"`
	Value   string            `json:"value" gorm:"column:value;type:text;not null;comment:配置值"`
	Desc    string            `json:"desc" gorm:"column:desc;type:varchar(200);not null;default:'';comment:配置说明"`
	Enabled bool              `json:"enabled" gorm:"column:enabled;type:tinyint(1);not null;default:1;comment:是否启用,0-否,1-是"`
}

func (c *Config) TableName() string {
	return "system_config"
}

func NewConfig() *Config {
	return &Config{}
}

func NewConfigList() []*Config {
	return make([]*Config, 0)
}

// GetConfigList 获取配置列表
func GetConfigList(db *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*Config, int, error) {
	qs := db.Where("deleted_at = null")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var count int64
	err := qs.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if err != nil {
		log.GetLogger().Errorf("model.system_config.GetConfigList Count Error ==> %s", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewConfigList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.system_config.GetConfigList Query Error ==> %s", err)
		return nil, 0, err
	}
	return result, int(count), nil
}

// CreateConfig 创建配置
func CreateConfig(db *gorm.DB, cfg *Config) (*Config, error) {
	err := db.Create(&cfg).Error
	if err != nil {
		log.GetLogger().Errorf("model.system_config.CreateConfig Error ==> %v", err)
		return nil, err
	}
	return cfg, nil
}

// UpdateConfig 更新配置
func UpdateConfig(db *gorm.DB, id int, data map[string]interface{}) error {
	err := db.Model(NewConfig()).Where("id = ?", id).UpdateColumns(data).Error
	if err != nil {
		log.GetLogger().Errorf("model.system_config.UpdateConfig Error ==> %v", err)
	}
	return err
}

// DeleteConfig 删除配置，逻辑删除
func DeleteConfig(db *gorm.DB, id int) error {
	return UpdateConfig(db, id, map[string]interface{}{"deleted_at": time.Now()})
}

// GetConfigById 根据ID获取单个配置数据
func GetConfigById(db *gorm.DB, id int) (*Config, error) {
	result := NewConfig()
	err := db.Model(NewConfig()).Where("id = ?", id).Limit(1).First(&result).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.GetLogger().Errorf("model.system_config.CreateConfig Error ==> %v", err)
		return nil, err
	}
	return result, nil
}
