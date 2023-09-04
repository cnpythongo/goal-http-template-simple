package account

import "github.com/cnpythongo/goal/model/account"

func NewUserProfile() *account.UserProfile {
	return &account.UserProfile{}
}

func NewUserProfileList() []*account.UserProfile {
	return make([]*account.UserProfile, 0)
}

func GetUserProfileObjectByUserId(userId int) (*account.UserProfile, error) {
	panic("implement me")
}
