package account

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	model.BaseModel
	UUID        string           `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;unique;comment:唯一ID"`
	Phone       string           `json:"phone" gorm:"column:phone;type:varchar(32);not null;comment:登录手机号码"`
	Password    string           `json:"-" gorm:"column:password;type:varchar(128);not null;comment:密码"`
	Salt        string           `json:"salt" gorm:"column:salt;type:varchar(24);not null;comment:密码加盐"`
	Nickname    string           `json:"nickname" gorm:"column:nickname;type:varchar(128);not null;comment:用户昵称"`
	Email       string           `json:"email" gorm:"column:email;type:varchar(128);default:'';comment:邮箱"`
	Avatar      string           `json:"avatar" gorm:"column:avatar;type:varchar(255);default:'';comment:用户头像"`
	Gender      int64            `json:"gender" gorm:"column:gender;type:int(11);default:0;comment:性别:0-保密,1-男,2-女"`
	Signature   string           `json:"signature" gorm:"column:signature;type:varchar(255);default:'';comment:个性化签名"`
	Status      userStatusType   `json:"status" gorm:"column:status;type:varchar(20);default:'INACTIVE';comment:用户状态"`
	LastLoginAt *model.LocalTime `json:"last_login_at" gorm:"column:last_login_at;default:null;comment:最后登录时间"`
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
