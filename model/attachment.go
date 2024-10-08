package model

import "gorm.io/gorm"

type Attachment struct {
	BaseModel
	UserId int64  `json:"user_id" gorm:"column:user_id;type:int(11);not null;comment:用户ID"` // 用户ID
	UUID   string `json:"uuid" gorm:"column:uuid;type:varchar(64);not null;comment:文件UUID"` // 文件UUID
	Name   string `json:"name" gorm:"column:name;type:varchar(256);not null;comment:文件名"`   // 文件名
	Ext    string `json:"ext" gorm:"column:ext;type:varchar(32);not null;comment:文件扩展名"`    // 文件扩展名
	Size   int64  `json:"size" gorm:"column:size;type:int(11);not null;comment:文件大小"`       // 文件大小
	Md5    string `json:"md5" gorm:"column:md5;type:varchar(128);not null;comment:文件MD5"`   // 文件MD5
	Path   string `json:"path" gorm:"column:path;type:varchar(512);not null;comment:文件路径"`  // 文件路径
	IP     string `json:"ip" gorm:"column:ip;type:varchar(64);not null;comment:上传附件时的IP"`   // 上传附件时的IP
}

func (a *Attachment) TableName() string {
	return "attachments"
}

func NewAttachment() *Attachment {
	return &Attachment{}
}

func GetAttachmentByUserIdAndMd5(db *gorm.DB, userId int64, md5 string) (*Attachment, error) {
	var result *Attachment
	err := db.Model(&Attachment{}).Where(
		"user_id = ? and md5 = ? and delete_time = 0", userId, md5,
	).Limit(1).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAttachmentListByUserId(db *gorm.DB, page, limit int, userId int64) ([]*Attachment, int64, error) {
	total := int64(0)
	result := make([]*Attachment, 0)
	query := db.Model(&Attachment{}).Where("user_id = ? and delete_time = 0", userId)
	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit)
	}
	err := query.Find(&result).Error
	if err != nil {
		return result, total, err
	}
	return result, total, err
}
