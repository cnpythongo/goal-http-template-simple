package generator

import (
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IService interface {
	GetDbTableList(req *ReqDbTables) (*RespDbTableList, int, error)
}

type service struct {
}

func NewService() IService {
	return &service{}
}

func (s *service) GetDbTableList(req *ReqDbTables) (*RespDbTableList, int, error) {
	query := model.GetDbTables(model.GetDB(), req.TableName, req.TableComment)
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.GetDbTables Count Error ==> %v", err)
		return nil, render.QueryError, err
	}
	result := make([]*RespDbTableItem, 0)
	var resp = &RespDbTableList{
		Page:   req.Page,
		Limit:  req.Limit,
		Total:  total,
		Result: result,
	}
	if total == 0 {
		return resp, render.OK, nil
	}
	err = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Find(&result).Error
	resp.Result = result
	return resp, render.OK, nil
}
