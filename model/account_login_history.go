package model

import "github.com/cnpythongo/goal/pkg/basic"

type AccountLoginHistory struct {
	basic.BaseModel
	UserID int64 `json:"user_id" gorm:"index:idx_account_login_history_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
}

func (h *AccountLoginHistory) TableName() string {
	return "account_login_history"
}

func NewAccountLoginHistory() *AccountLoginHistory {
	return &AccountLoginHistory{}
}

func NewAccountLoginHistoryList() []*AccountLoginHistory {
	return make([]*AccountLoginHistory, 0)
}

func GetLoginHistoryObject(id int) (*AccountLoginHistory, error) {
	panic("implement me")
}

func GetUserLoginHistoryQueryset(userId, page, size int) ([]*AccountLoginHistory, error) {
	panic("implement me")
}

func GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*AccountLoginHistory, error) {
	panic("implement me")
}
