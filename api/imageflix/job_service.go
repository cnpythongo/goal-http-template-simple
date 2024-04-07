package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"gorm.io/gorm"
)

type IImageFlixJobService interface {
	Create(userId, usePoint int64, src string) error
	Start(jobId, userId int64) error
}

type imageFlixJobService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewImageFlixJobService() IImageFlixJobService {
	db := model.GetDB()
	return &imageFlixJobService{
		db: db,
	}
}

func (s *imageFlixJobService) Create(userId, usePoint int64, src string) error {
	job := model.ImageFlixJob{
		UserId:   userId,
		UsePoint: usePoint,
		JobType:  model.ImageFlixJobTypeRestore,
		Status:   model.ImageFlixJobStatusWait,
		Src:      src,
	}
	err := job.Create()
	return err
}

func (s *imageFlixJobService) Start(jobId, userId int64) error {
	return model.UpdateImageFlixJobStatus(
		s.db, jobId, userId, model.ImageFlixJobStatusMaking,
	)
}
