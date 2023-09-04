package account

import "github.com/cnpythongo/goal/model/account"

func NewLoginHistory() *account.LoginHistory {
	return &account.LoginHistory{}
}

func NewLoginHistoryList() []*account.LoginHistory {
	return make([]*account.LoginHistory, 0)
}

func GetLoginHistoryObject(id int) (*account.LoginHistory, error) {
	panic("implement me")
}

func GetUserLoginHistoryQueryset(userId, page, size int) ([]*account.LoginHistory, error) {
	panic("implement me")
}

func GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*account.LoginHistory, error) {
	panic("implement me")
}
