package systemlog

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemLogService interface {
	List(req *ReqSystemLogList) (res []*RespSystemLogItem, total int64, code int, err error)
	Detail(req *ReqSystemLogDetail) (res *RespSystemLogItem, code int, err error)
	Create(payload *ReqSystemLogCreate) (res *RespSystemLogItem, code int, err error)
	Update(payload *ReqSystemLogUpdate) (res *RespSystemLogItem, code int, err error)
	Delete(payload *ReqSystemLogDelete) (code int, e error)
	GetAllSystemLog() ([]*model.SystemLog, error)
}

// systemLogService 系统日志服务实现类
type systemLogService struct{}

// NewSystemLogService 初始化
func NewSystemLogService() ISystemLogService {
	return &systemLogService{}
}

// List 系统日志列表
func (s *systemLogService) List(req *ReqSystemLogList) (res []*RespSystemLogItem, total int64, code int, err error) {
	// 分页信息
	limit := req.Page
	offset := req.Limit * (req.Page - 1)
	// 查询
	query := model.GetDB().Model(&model.SystemLog{})
	if req.Cellphone != "" {
		query = query.Where("cellphone like ?", "%"+req.Cellphone+"%")
	}
	if req.MemberName != "" {
		query = query.Where("member_name like ?", "%"+req.MemberName+"%")
	}
	// 总数
	err = query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemLog
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	err = copier.Copy(&res, objs)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.DBAttributesCopyError, err
	}
	return res, total, render.OK, nil
}

// Detail 系统日志详情
func (s *systemLogService) Detail(req *ReqSystemLogDetail) (res *RespSystemLogItem, code int, err error) {
	// var obj *model.SystemLog
	obj, err := model.GetSystemLogInstance(
		model.GetDB(),
		map[string]interface{}{
			"id":          req.ID,
			"delete_time": 0,
		},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		return nil, render.QueryError, err
	}
	err = copier.Copy(&res, obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return
}

// Create 系统日志创建
func (s *systemLogService) Create(payload *ReqSystemLogCreate) (res *RespSystemLogItem, code int, err error) {
	obj := model.NewSystemLog()
	err = copier.Copy(&obj, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	obj, err = model.CreateSystemLog(model.GetDB(), obj)
	if err != nil {
		return nil, render.CreateError, err
	}
	err = copier.Copy(&res, obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Update 系统日志更新
func (s *systemLogService) Update(payload *ReqSystemLogUpdate) (res *RespSystemLogItem, code int, err error) {
	obj, err := model.GetSystemLogInstance(
		model.GetDB(),
		map[string]interface{}{
			"id":          payload.ID,
			"delete_time": 0,
		},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		return nil, render.QueryError, err
	}
	// 更新
	err = copier.Copy(&obj, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	obj.UpdateTime = time.Now().Unix()
	err = model.UpdateSystemLog(model.GetDB(), obj)
	if err != nil {
		return nil, render.UpdateError, err
	}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Delete 系统日志删除
func (s *systemLogService) Delete(payload *ReqSystemLogDelete) (code int, e error) {
	_, err := model.GetSystemLogInstance(
		model.GetDB(),
		map[string]interface{}{
			"id":          payload.ID,
			"delete_time": 0,
		},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return render.DataNotExistError, err
		}
		return render.QueryError, err
	}
	// 删除
	err = model.DeleteSystemLog(model.GetDB(), payload.ID)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemLog 系统日志获取所有有效数据
func (s *systemLogService) GetAllSystemLog() ([]*model.SystemLog, error) {
	result, err := model.GetAllSystemLog(model.GetDB())
	if err != nil {
		return nil, err
	}
	return result, err
}
