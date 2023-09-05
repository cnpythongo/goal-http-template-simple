package model

import "gorm.io/gorm"

type History struct {
	BaseModel
	UserID int64 `json:"user_id" gorm:"index:idx_account_history_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
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

func GetHistory(db *gorm.DB, id int) (*History, error) {
	panic("implement me")
}

func GetHistoryList(db *gorm.DB, page, size int, conditions map[string]interface{}) ([]*History, error) {
	panic("implement me")
}
