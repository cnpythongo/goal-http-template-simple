package model

import (
	"goal-app/pkg/log"
	"gorm.io/gorm"
)

type UserProfile struct {
	BaseModel
	UserID   int64  `json:"user_id" gorm:"index:idx_account_user_profile_user_id;column:user_id;type:int(11);not null;comment:用户ID"`
	RealName string `json:"real_name" gorm:"column:real_name;type:varchar(50);not null;comment:真实姓名"`
	IDNumber string `json:"id_number" gorm:"column:id_number;type:varchar(50);not null;comment:身份证号"`
}

func (p *UserProfile) TableName() string {
	return "account_user_profile"
}

// RealNameMask 真实姓名打码
func (p *UserProfile) RealNameMask() string {
	if len(p.RealName) == 3 {
		return p.RealName[0:1] + "*" + p.RealName[2:3]
	} else if len(p.RealName) == 4 {
		return p.RealName[0:1] + "*" + p.RealName[3:4]
	} else {
		return p.RealName[0:1] + "*"
	}
}

// IDNumberMask 身份证号打码
func (p *UserProfile) IDNumberMask() string {
	if p.IDNumber == "" {
		return p.IDNumber
	}
	num := p.IDNumber
	if len(num) == 15 {
		return num[0:6] + "****" + num[14:15]
	} else if len(num) == 18 {
		return num[0:6] + "****" + num[16:18]
	}
	return num
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfileList() []*UserProfile {
	return make([]*UserProfile, 0)
}

// GetUserProfileByUserId 获取单个用户的个人资料
func GetUserProfileByUserId(db *gorm.DB, userId uint64) (*UserProfile, error) {
	pf := NewUserProfile()
	err := db.Model(NewUserProfile()).Where("user_id = ?", userId).Limit(1).First(&pf).Error
	return pf, err
}

// GetUserProfileList 获取个人资料列表
func GetUserProfileList(db *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*UserProfile, int, error) {
	qs := db.Model(NewUserProfile()).Where("delete_time = 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var count int64
	err := qs.Count(&count).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user_profile.GetUserProfileList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewUserProfileList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user_profile.GetUserProfileList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, int(count), nil
}

func CreateUserProfile(db *gorm.DB, pf *UserProfile) (*UserProfile, error) {
	err := db.Save(&pf).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user_profile.CreateUserProfile Error ==> %v", err)
	}
	return pf, err
}

// UpdateUserProfileByUserId 根据用户ID更新用户资料
func UpdateUserProfileByUserId(db *gorm.DB, userId uint64, data map[string]interface{}) error {
	err := db.Where("user_id = ?", userId).UpdateColumns(data).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user_profile.UpdateUserProfileByUserId Error ==> %v", err)
	}
	return err
}

// UpdateUserProfile 更新用户资料
func UpdateUserProfile(db *gorm.DB, pf *UserProfile) (*UserProfile, error) {
	err := db.Updates(&pf).Error
	return pf, err
}

// DeleteUserProfileByUserId 根据用户ID删除用户资料
func DeleteUserProfileByUserId(db *gorm.DB, userId uint64) error {
	err := db.Where("user_id = ?", userId).Delete(NewUserProfile()).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user_profile.DeleteUserProfileByUserId Error ==> %v", err)
	}
	return err
}
