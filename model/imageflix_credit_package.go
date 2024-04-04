package model

import (
	"github.com/cnpythongo/goal-tools/utils"
	"gorm.io/gorm"
)

// ImageFlixPackage 资源包，用户购买资源包后才能用
type ImageFlixPackage struct {
	BaseModel
	Usable   int64            `json:"usable" gorm:"cloumn:usable;type:int(11);comment:购买点数"`
	OldPrice int64            `json:"old_price" gorm:"column:old_price;type:int(11);comment:原价(单位:分)"`
	Price    int64            `json:"price" gorm:"column:price;type:int(11);comment:活动售价(单位:分)"`
	EndTime  *utils.LocalTime `json:"end_time" gorm:"column:end_time;type:int(11);comment:活动结束时间"`
	Sort     int64            `json:"sort" gorm:"column:sort;type:int(11);comment:排序值"`
}

func (i *ImageFlixPackage) TableName() string {
	return "imageflix_packages"
}

func NewImageFlixPackage() *ImageFlixPackage {
	return &ImageFlixPackage{}
}

func (i *ImageFlixPackage) Create() error {
	return db.Model(NewImageFlixPackage()).Create(&i).Error
}

func (i *ImageFlixPackage) Save() error {
	return db.Model(NewImageFlixPackage()).Save(&i).Error
}

func GetImageFlixPackageList(db *gorm.DB) ([]*ImageFlixPackage, error) {
	result := make([]*ImageFlixPackage, 0)
	err := db.Model(NewImageFlixPackage()).Where(
		"deleted_at = null",
	).Order("sort desc").Find(&result).Error
	return result, err
}
