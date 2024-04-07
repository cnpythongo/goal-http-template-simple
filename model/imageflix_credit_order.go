package model

import "gorm.io/gorm"

type ImageFlixCreditOrder struct {
	BaseModel
	UserId    int64  `json:"user_id" gorm:"column:user_id;type:int(11);comment:用户ID"`
	PackageId int64  `json:"package_id" gorm:"column:package_id;type:int(11);comment:资源包ID"`
	TradeNum  string `json:"trade_num" gorm:"column:trade_num;type:varchar(128);comment:交易订单号"`
	Payment   int64  `json:"payment" gorm:"column:payment;type:int(11);comment:实付金额(单位:分)"`
	Platform  string `json:"platform" gorm:"column:platform;type:varchar(64);comment:支付平台"`
}

func (i *ImageFlixCreditOrder) TableName() string {
	return "imageflix_credit_orders"
}

func NewImageFlixCreditOrder() *ImageFlixCreditOrder {
	return &ImageFlixCreditOrder{}
}

func (i *ImageFlixCreditOrder) Create() error {
	return db.Model(&ImageFlixCreditOrder{}).Create(&i).Error
}

func (i *ImageFlixCreditOrder) Save() error {
	return db.Model(&ImageFlixCreditOrder{}).Save(&i).Error
}

func GetImageFlixCreditOrderByUserId(db *gorm.DB, page, limit int, userId int64) ([]*ImageFlixCreditOrder, int64, error) {
	query := db.Model(NewImageFlixCreditOrder()).Where(
		"deleted_at = null and user_id = ?", userId,
	)
	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit)
	}

	total := int64(0)
	result := make([]*ImageFlixCreditOrder, 0)
	err := query.Count(&total).Error
	if err != nil {
		return result, total, err
	}
	err = query.Find(&result).Error
	return result, total, err
}
