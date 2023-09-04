package account

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/pkg/log"
	"gorm.io/gorm"
)

func NewUser() *account.User {
	return &account.User{}
}

func NewUserList() []*account.User {
	return make([]*account.User, 0)
}

func GetUserByConditions(conditions interface{}) (*account.User, error) {
	result := NewUser()
	err := model.GetDB().Where(conditions).First(result).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.GetLogger().Infof("conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.account.user.GetUserByConditions Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetUserByPhone(phone string) (*account.User, error) {
	conditions := map[string]interface{}{"phone": phone}
	result, err := GetUserByConditions(conditions)
	return result, err
}

func GetUserByEmail(email string) (*account.User, error) {
	conditions := map[string]interface{}{"email": email}
	result, err := GetUserByConditions(conditions)
	return result, err
}

func CreateUser(user *account.User) (*account.User, error) {
	err := model.GetDB().Create(user).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.CreateUser Error ==> %v", err)
		return nil, err
	}
	return user, nil
}

func GetUserByUUID(uuid string) (*account.User, error) {
	conditions := map[string]interface{}{"uuid": uuid}
	result, err := GetUserByConditions(conditions)
	return result, err
}

func GetUserQueryset(page, size int, conditions interface{}) ([]*account.User, int, error) {
	qs := model.GetDB().Model(NewUser())
	if conditions != nil {
		qs = qs.Where(conditions)
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	var count int64
	err := qs.Count(&count).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.GetUserQueryset Count Error ==> ", err)
		return nil, 0, err
	}
	result := NewUserList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.GetUserQueryset Query Error ==> ", err)
		return nil, 0, err
	}
	return result, int(count), nil
}
