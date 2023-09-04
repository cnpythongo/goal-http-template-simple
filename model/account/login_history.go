package account

import "github.com/cnpythongo/goal/model"

type LoginHistory struct {
	model.BaseModel
	UserID int64 `json:"user_id" gorm:"index:idx_account_login_history_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
}

func (h *LoginHistory) TableName() string {
	return "account_login_history"
}
