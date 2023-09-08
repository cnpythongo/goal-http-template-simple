package model

import "gorm.io/gorm"

type UserProfile struct {
	BaseModel
	UserID   int64  `json:"user_id" gorm:"index:idx_account_user_profile_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
	RealName string `json:"real_name" gorm:"column:real_name;type:varchar(50);not null;comment:真实姓名"`
	IDNumber string `json:"id_number" gorm:"column:id_number;type:varchar(50);not null;comment:身份证号"`
}

func (p *UserProfile) TableName() string {
	return "account_user_profile"
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfileList() []*UserProfile {
	return make([]*UserProfile, 0)
}

func GetUserProfileByUserId(db *gorm.DB, userId int) (*UserProfile, error) {
	panic("implement me")
}