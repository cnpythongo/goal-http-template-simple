package model

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	BaseModel
	UUID        string         `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;unique;comment:唯一ID"`
	Phone       string         `json:"phone" gorm:"column:phone;type:varchar(32);not null;comment:登录手机号码"`
	Password    string         `json:"-" gorm:"column:password;type:varchar(128);not null;comment:密码"`
	Salt        string         `json:"salt" gorm:"column:salt;type:varchar(24);not null;comment:密码加盐"`
	Nickname    string         `json:"nickname" gorm:"column:nickname;type:varchar(128);not null;comment:用户昵称"`
	Email       string         `json:"email" gorm:"column:email;type:varchar(128);default:'';comment:邮箱"`
	Avatar      string         `json:"avatar" gorm:"column:avatar;type:varchar(255);default:'';comment:用户头像"`
	Gender      int64          `json:"gender" gorm:"column:gender;type:int(11);default:0;comment:性别:0-保密,1-男,2-女"`
	Signature   string         `json:"signature" gorm:"column:signature;type:varchar(255);default:'';comment:个性化签名"`
	Status      userStatusType `json:"status" gorm:"column:status;type:varchar(20);default:'INACTIVE';comment:用户状态"`
	LastLoginAt *LocalTime     `json:"last_login_at" gorm:"column:last_login_at;default:null;comment:最后登录时间"`
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

func NewUser() *User {
	return &User{}
}

func NewUserList() []*User {
	return make([]*User, 0)
}

func GetUserByConditions(db *gorm.DB, conditions interface{}) (*User, error) {
	result := NewUser()
	err := db.Where(conditions).First(result).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.GetLogger().Infof("conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.account.user.GetUserByConditions Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

//	func (u *User) GetUserByPhone(phone string) (*User, error) {
//		conditions := map[string]interface{}{"phone": phone}
//		result, err := GetUserByConditions(conditions)
//		return result, err
//	}
//
//	func (u *User) GetUserByEmail(email string) (*User, error) {
//		conditions := map[string]interface{}{"email": email}
//		result, err := GetUserByConditions(conditions)
//		return result, err
//	}
//
//	func (u *User) CreateUser(user *User) (*User, error) {
//		err := model.GetDB().Create(user).Error
//		if err != nil {
//			log.GetLogger().Errorf("model.account.user.CreateUser Error ==> %v", err)
//			return nil, err
//		}
//		return user, nil
//	}
//
//	func (u *User) GetUserByUUID(uuid string) (*User, error) {
//		conditions := map[string]interface{}{"uuid": uuid}
//		result, err := GetUserByConditions(conditions)
//		return result, err
//	}
func GetUserList(db *gorm.DB, page, size int, conditions interface{}) ([]*User, int, error) {
	qs := db.Model(NewUser())
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
