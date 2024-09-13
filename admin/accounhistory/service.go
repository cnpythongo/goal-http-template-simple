package accounhistory

import (
	"github.com/gin-gonic/gin"
)

type IHistoryService interface {
	GetHistoryList(c *gin.Context)
}

type historyService struct {
}

func NewHistoryService() IHistoryService {
	return &historyService{}
}

func (h *historyService) GetHistoryList(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
