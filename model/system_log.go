package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type SystemLog struct {
	BaseModel
	UserUUID     string `json:"user_uuid" gorm:"column:user_uuid;type:varchar(128);not null;default:'';comment:用户ID"`
	Cellphone    string `json:"cellphone" gorm:"column:cellphone;type:varchar(32);not null;default:'';comment:手机号"`
	ClientIP     string `json:"client_ip" gorm:"column:client_ip;type:varchar(128);not null;default:'';comment:客户端IP"`
	Path         string `json:"path" gorm:"column:path;type:varchar(1024);not null;default:'';comment:请求路径"`
	Body         string `json:"body" gorm:"column:body;type:text;not null;comment:请求Body"`
	Method       string `json:"method" gorm:"column:method;type:varchar(64);not null;default:'';comment:请求方法"`
	Status       int64  `json:"status" gorm:"column:status;type:int(11);not null;default:0;comment:状态码"`
	UserAgent    string `json:"user_agent" gorm:"column:user_agent;type:varchar(1024);not null;default:'';comment:请求UA"`
	Referer      string `json:"referer" gorm:"column:referer;type:varchar(1024);not null;default:'';comment:请求携带的Referer"`
	StartTime    int64  `json:"start_time" gorm:"column:start_time;type:int(11);not null;default:0;comment:请求开始时间"`
	EndTime      int64  `json:"end_time" gorm:"column:end_time;type:int(11);not null;default:0;comment:请求结束时间"`
	LatencyTime  string `json:"latency_time" gorm:"column:latency_time;type:varchar(128);not null;default:'';comment:请求耗时"`
	MemberName   string `json:"member_name" gorm:"column:member_name;type:varchar(64);not null;default:'';comment:用户名"`
	OperateTitle string `json:"operate_title" gorm:"column:operate_title;type:varchar(64);not null;default:'';comment:操作标题"`
	Page         string `json:"page" gorm:"column:page;type:varchar(64);not null;default:'';comment:请求触发页面"`
}

func (m *SystemLog) TableName() string {
	return "system_logs"
}

func NewSystemLog() *SystemLog {
	return &SystemLog{}
}

func NewSystemLogList() []*SystemLog {
	return make([]*SystemLog, 0)
}

func (m *SystemLog) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func CreateSystemLog(tx *gorm.DB, obj *SystemLog) (*SystemLog, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemLog.CreateSystemLog Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func UpdateSystemLog(tx *gorm.DB, obj *SystemLog) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemLog.UpdateSystemLog Error ==> %v", err)
	}
	return err
}

func DeleteSystemLog(tx *gorm.DB, id int64) error {
	err := tx.Model(NewSystemLog()).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemLog.DeleteSystemLog Error ==> %v", err)
	}
	return err
}

func GetSystemLogInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemLog, error) {
	result := NewSystemLog()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemLog.GetSystemLogInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemLog.GetSystemLogInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemLogList(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*SystemLog, int64, error) {
	qs := tx.Model(NewSystemLog()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemLog.GetSystemLogList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewSystemLogList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemLog.GetSystemLogList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemLog(tx *gorm.DB) ([]*SystemLog, error) {
	result := NewSystemLogList()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemLog.GetAllSystemLog Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}
