package model

import (
	"gorm.io/gorm"
	"time"
)

type ImageFlixCredit struct {
	BaseModel
	UserId int64 `json:"user_id" gorm:"column:user_id;type:int(11);comment:用户ID"`
	Usable int64 `json:"usable" gorm:"column:usable;type:int(11);comment:可用点数"`
}

func (i *ImageFlixCredit) TableName() string {
	return "imageflix_credits"
}

func NewImageFlixCredit() *ImageFlixCredit {
	return &ImageFlixCredit{}
}

func GetImageFlixCreditByUser(db *gorm.DB, userId int64) (*ImageFlixCredit, error) {
	result := NewImageFlixCredit()
	err := db.Model(NewImageFlixCredit()).Where(
		"deleted_at = null and user_id = ?", userId,
	).First(&result).Error
	return result, err
}

func UpdateImageFlixCreditByUser(db *gorm.DB, userId, point int64) error {
	err := db.Model(NewImageFlixCredit()).Where(
		"deleted_at = null and user_id = ?", userId,
	).UpdateColumns(map[string]interface{}{
		"usable":     gorm.Expr("usable + ?", point),
		"updated_at": time.Now(),
	}).Error
	return err
}
