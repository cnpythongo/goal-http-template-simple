package model

import (
	"gorm.io/gorm"
	"time"
)

type ImageFlixJobStatus int64

const (
	ImageFlixJobStatusQueue   = 1
	ImageFlixJobStatusWait    = 2
	ImageFlixJobStatusMaking  = 3
	ImageFlixJobStatusFail    = 4
	ImageFlixJobStatusSuccess = 5
)

type ImageFlixJobType string

const (
	ImageFlixJobTypeRestore = "restore"
)

type ImageFlixJob struct {
	BaseModel
	UserId   int64              `json:"user_id" gorm:"column:user_id;type:int(11);comment:用户ID"`
	UsePoint int64              `json:"use_point" gorm:"column:use_point;type:int(11);comment:使用点数"`
	JobType  ImageFlixJobType   `json:"job_type" gorm:"column:job_type;type:varchar(32);comment:任务类型,restore-修复"`
	Status   ImageFlixJobStatus `json:"status" gorm:"column:status;type:varchar(32);comment:状态,1-排队;2-待开始;3-处理中;4-失败;5-成功"`
	Src      string             `json:"src" gorm:"column:src;type:varchar(512);comment:源图片地址"`
	Target   string             `json:"target" gorm:"column:target;type:varchar(512);comment:修复图片地址"`
}

func (i *ImageFlixJob) TableName() string {
	return "imageflix_job"
}

func NewImageFlixJob() *ImageFlixJob {
	return &ImageFlixJob{}
}

func (i *ImageFlixJob) Create() error {
	return db.Model(NewImageFlixJob()).Create(&i).Error
}

func (i *ImageFlixJob) Save() error {
	return db.Model(NewImageFlixJob()).Save(&i).Error
}

func GetImageFlixJobByUserId(db *gorm.DB, page, limit int, userId, status int64, jobType string) ([]*ImageFlixJob, int64, error) {
	query := db.Model(NewImageFlixJob()).Where("deleted_at = null and user_id = ?", userId)
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if jobType != "" {
		query = query.Where("job_type = ?", jobType)
	}
	total := int64(0)
	result := make([]*ImageFlixJob, 0)
	err := query.Count(&total).Error
	if err != nil {
		return result, total, err
	}

	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit)
	}
	err = query.Find(&result).Error
	return result, total, err
}

func UpdateImageFlixJobStatus(db *gorm.DB, jobId, userId int64, status ImageFlixJobStatus) error {
	return db.Model(NewImageFlixJob()).Where(
		"job_id = ? and user_id = ?", jobId, userId,
	).Updates(map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}).Error
}
