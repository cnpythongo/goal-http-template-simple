package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type History struct {
	BaseModel
	UserID int64                `json:"user_id" gorm:"index:idx_account_history_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
	Device AccountHistoryDevice `json:"device" gorm:"column:device;type:varchar(50);not null;default:'web';comment:登录设备"`
	IP     string               `json:"ip" gorm:"column:ip;type:varchar(50);not null;default:'';comment:登录IP"`
	Locate string               `json:"locate" gorm:"column:locate;type:varchar(100);not null;default:'';comment:归属地"`
}

func NewHistory() *History {
	return &History{}
}

func NewHistoryList() []*History {
	return make([]*History, 0)
}

func (h *History) TableName() string {
	return "account_history"
}

// GetHistoryById 获取单条历史记录
func GetHistoryById(db *gorm.DB, id int) (*History, error) {
	result := NewHistory()
	err := db.Model(NewHistory()).Where("id = ?", id).Limit(1).First(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.account_history.GetHistoryById Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

// GetHistoryList 获取账号登录历史列表
func GetHistoryList(db *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*History, int, error) {
	qs := db.Model(NewHistory()).Where("deleted_at = null")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var count int64
	err := qs.Count(&count).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_history.GetHistoryList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewHistoryList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_history.GetHistoryList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, int(count), nil
}

// CreateHistory 创建登录历史记录
func CreateHistory(db *gorm.DB, h *History) (*History, error) {
	err := db.Create(&h).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_history.CreateHistory Error ==> %v", err)
		return nil, err
	}
	return h, nil
}

func DeleteHistoryByIds(db *gorm.DB, ids []int) error {
	return db.Model(NewHistory()).Where("id in ?", ids).Update("deleted_at", time.Now()).Error
}

func DeleteHistoryByUserIds(db *gorm.DB, userIds []int) error {
	return db.Model(NewHistory()).Where("user_id in ?", userIds).Update("deleted_at", time.Now()).Error
}
