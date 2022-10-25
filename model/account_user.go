package model

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/pkg/basic"
	"github.com/cnpythongo/goal/pkg/common/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	basic.BaseModel
	UUID        string `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;unique;comment:唯一ID"`
	Username    string `json:"username" gorm:"column:username;type:varchar(256);unique;not null;comment:用户名"`
	Password    string `json:"-" gorm:"column:password;type:varchar(200);not null;comment:密码"`
	Salt        string `json:"salt" gorm:"column:salt;type:varchar(20);not null;comment:密码加盐"`
	Email       string `json:"email" gorm:"column:email;type:varchar(200);default:'';comment:邮箱"`
	Avatar      string `json:"avatar" gorm:"column:avatar;type:varchar(200);default:'';comment:用户头像"`
	Gender      int    `json:"gender" gorm:"column:gender;type:int(11);default:0;comment:性别:0-保密,1-男,2-女"`
	Signature   string `json:"signature" gorm:"column:signature;type:varchar(512);default:'';comment:个性化签名"`
	LastLoginAt int64  `json:"last_login_at" gorm:"column:last_login_at;default:0;comment:最后登录时间"`
}

func NewUser() *User {
	return &User{}
}

func NewUsers() []*User {
	return make([]*User, 0)
}

func (u *User) TableName() string {
	return "account_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	us := uuid.New().String()
	u.UUID = strings.ReplaceAll(us, "-", "")
	hashPwd, salt := utils.GeneratePassword(u.Password)
	u.Password = hashPwd
	u.Salt = salt
	return nil
}

func GetUserByCondition(condition interface{}) (*User, error) {
	result := NewUser()
	err := db.Where(condition).First(result).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.GetLogger().Errorf("model.account.user.GetUserByCondition Error ==> %v", err)
			log.GetLogger().Infof("condition ==> %v", condition)
		}
		return nil, err
	}
	return result, nil
}

func GetUserByUsername(username string) (*User, error) {
	condition := map[string]interface{}{"username": username}
	result, err := GetUserByCondition(condition)
	return result, err
}

func GetUserByEmail(email string) (*User, error) {
	condition := map[string]interface{}{"email": email}
	result, err := GetUserByCondition(condition)
	return result, err
}

func CreateUser(user *User) (*User, error) {
	err := db.Create(user).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.CreateUser Error ==> %v", err)
		return nil, err
	}
	return user, nil
}

func GetUserByUuid(uuid string) (*User, error) {
	condition := map[string]interface{}{"uuid": uuid}
	result, err := GetUserByCondition(condition)
	return result, err
}

func GetUserQueryset(page, size int, conditions interface{}) ([]*User, int, error) {
	qs := db.Model(NewUser())
	if conditions != nil {
		qs = qs.Where(conditions)
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.GetUserQueryset Count Error ==> ", err)
		return nil, 0, err
	}
	result := NewUsers()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.account.user.GetUserQueryset Query Error ==> ", err)
		return nil, 0, err
	}
	return result, int(total), nil
}

func GetUserById(userID int) (*User, error) {
	condition := map[string]interface{}{"id": userID}
	result, err := GetUserByCondition(condition)
	return result, err
}
