package model

import (
	"errors"
	"goal-app/pkg/log"
	"goal-app/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	BaseModel
	UUID          string         `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;unique;comment:唯一ID"`
	Phone         string         `json:"phone" gorm:"column:phone;type:varchar(32);not null;comment:登录手机号码"`
	Password      string         `json:"-" gorm:"column:password;type:varchar(128);not null;comment:密码"`
	Salt          string         `json:"salt" gorm:"column:salt;type:varchar(24);not null;comment:密码加盐"`
	Nickname      string         `json:"nickname" gorm:"column:nickname;type:varchar(128);not null;comment:用户昵称"`
	Email         string         `json:"email" gorm:"column:email;type:varchar(128);not null;default:'';comment:邮箱"`
	Avatar        string         `json:"avatar" gorm:"column:avatar;type:varchar(255);not null;default:'';comment:用户头像"`
	Gender        UserGender     `json:"gender" gorm:"column:gender;type:int(11);not null;default:3;comment:性别:3-保密,1-男,2-女"`
	Signature     string         `json:"signature" gorm:"column:signature;type:varchar(255);not null;default:'';comment:个性化签名"`
	Status        UserStatusType `json:"status" gorm:"column:status;type:varchar(20);not null;default:'INACTIVE';comment:用户状态"`
	IsAdmin       bool           `json:"is_admin" gorm:"column:is_admin;type:tinyint(1);not null;default:0;comment:是否admin账号,默认否"`
	LastLoginTime int64          `json:"last_login_time" gorm:"column:last_login_time;default:0;comment:最后登录时间"`
}

func (u *User) TableName() string {
	return "account_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = utils.UUID()
	hashPwd, salt := utils.GeneratePassword(u.Password)
	u.Password = hashPwd
	u.Salt = salt
	return nil
}

func (u *User) PhoneMask() string {
	if len(u.Phone) == 11 {
		return u.Phone[0:4] + "****" + u.Phone[7:11]
	}
	return u.Phone
}

func NewUser() *User {
	return &User{}
}

func NewUserList() []*User {
	return make([]*User, 0)
}

func GetUserByConditions(db *gorm.DB, conditions interface{}) (*User, error) {
	result := NewUser()
	err := db.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.account_user.GetUserByConditions conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.account_user.GetUserByConditions Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetUserByPhone(db *gorm.DB, phone string) (*User, error) {
	return GetUserByConditions(db, map[string]interface{}{"phone": phone})
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	return GetUserByConditions(db, map[string]interface{}{"email": email})
}

func GetUserByUUID(db *gorm.DB, uuid string) (*User, error) {
	return GetUserByConditions(db, map[string]interface{}{"uuid": uuid})
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	err := db.Create(&user).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user.CreateUser Error ==> %v", err)
		return nil, err
	}
	return user, nil
}

func GetUserList(db *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*User, int64, error) {
	qs := db.Model(NewUser()).Where("status != ?", UserStatusDelete)
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user.GetUserList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewUserList()
	err = qs.Order("create_time desc").Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user.GetUserList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

// UpdateUserLastLoginAt 更新用户最近登录时间
func UpdateUserLastLoginAt(db *gorm.DB, uuid string) error {
	return UpdateUser(db, uuid, map[string]interface{}{
		"status":          UserStatusActive,
		"last_login_time": time.Now().Unix(),
		"update_time":     time.Now().Unix(),
	})
}

func UpdateUser(db *gorm.DB, uuid string, data map[string]interface{}) error {
	err := db.Model(NewUser()).Where("uuid = ?", uuid).UpdateColumns(data).Error
	if err != nil {
		log.GetLogger().Errorf("model.account_user.UpdateUser Error ==> %v", err)
	}
	return err
}

func DeleteUser(db *gorm.DB, uuid string) error {
	return UpdateUser(db, uuid, map[string]interface{}{
		"status":      UserStatusDelete,
		"delete_time": time.Now().Unix(),
	})
}
