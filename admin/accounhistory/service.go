package accounhistory

import (
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"gorm.io/gorm"
)

type IHistoryService interface {
	GetHistoryList(c *gin.Context)
}

type historyService struct {
	db *gorm.DB
}

func NewHistoryService() IHistoryService {
	db := model.GetDB()
	return &historyService{db: db}
}

func (h *historyService) GetHistoryList(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
