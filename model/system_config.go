package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type SystemConfig struct {
	BaseModel
	Scope   SystemConfigScope `json:"scope" gorm:"column:scope;type:varchar(50);not null;default:'';comment:作用域,global-全局,admin-管理后台,app-前台应用"`
	Name    string            `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:配置名称"`
	Value   string            `json:"value" gorm:"column:value;type:text;not null;comment:配置值"`
	Desc    string            `json:"desc" gorm:"column:desc;type:varchar(200);not null;default:'';comment:配置说明"`
	Enabled bool              `json:"enabled" gorm:"column:enabled;type:tinyint(1);not null;default:1;comment:是否启用,0-否,1-是"`
}

func (m *SystemConfig) TableName() string {
	return "system_config"
}

func NewSystemConfig() *SystemConfig {
	return &SystemConfig{}
}

func NewSystemConfigList() []*SystemConfig {
	return make([]*SystemConfig, 0)
}

func (m *SystemConfig) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func CreateSystemConfig(tx *gorm.DB, obj *SystemConfig) (*SystemConfig, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemConfig.CreateSystemConfig Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func UpdateSystemConfig(tx *gorm.DB, obj *SystemConfig) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemConfig.UpdateSystemConfig Error ==> %v", err)
	}
	return err
}

func DeleteSystemConfig(tx *gorm.DB, id int64) error {
	err := tx.Model(NewSystemConfig()).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemConfig.DeleteSystemConfig Error ==> %v", err)
	}
	return err
}

func GetSystemConfigInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemConfig, error) {
	result := NewSystemConfig()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemConfig.GetSystemConfigInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemConfig.GetSystemConfigInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemConfigList(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*SystemConfig, int64, error) {
	qs := tx.Model(NewSystemConfig()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemConfig.GetSystemConfigList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewSystemConfigList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemConfig.GetSystemConfigList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemConfig(tx *gorm.DB) ([]*SystemConfig, error) {
	result := NewSystemConfigList()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemConfig.GetAllSystemConfig Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}
