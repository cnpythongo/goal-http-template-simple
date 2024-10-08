package systemconfig

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemConfigService interface {
	List(req *ReqSystemConfigList) (res []*RespSystemConfigItem, total int64, code int, err error)
	Detail(req *ReqSystemConfigDetail) (res *RespSystemConfigItem, code int, err error)
	Create(payload *ReqSystemConfigCreate) (res *RespSystemConfigItem, code int, err error)
	Update(payload *ReqSystemConfigUpdate) (res *RespSystemConfigItem, code int, err error)
	Delete(payload *ReqSystemConfigDelete) (code int, e error)
	GetAllSystemConfig() ([]*model.SystemConfig, error)
}

// systemConfigService 系统配置项服务实现类
type systemConfigService struct{}

// NewSystemConfigService 初始化
func NewSystemConfigService() ISystemConfigService {
	return &systemConfigService{}
}

// List 系统配置项列表
func (s *systemConfigService) List(req *ReqSystemConfigList) (res []*RespSystemConfigItem, total int64, code int, err error) {
	// 分页信息
	limit := req.Page
	offset := req.Limit * (req.Page - 1)
	// 查询
	query := model.GetDB().Model(&model.SystemConfig{})
	if req.Scope != "" {
		query = query.Where("scope = ?", req.Scope)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Value != "" {
		query = query.Where("value = ?", req.Value)
	}
	if req.Desc != "" {
		query = query.Where("desc = ?", req.Desc)
	}
	if req.Enabled >= 0 {
		query = query.Where("enabled = ?", req.Enabled)
	}
	// 总数
	err = query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemConfig
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

// Detail 系统配置项详情
func (s *systemConfigService) Detail(req *ReqSystemConfigDetail) (res *RespSystemConfigItem, code int, err error) {
	// var obj *model.SystemConfig
	obj, err := model.GetSystemConfigInstance(
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

// Create 系统配置项创建
func (s *systemConfigService) Create(payload *ReqSystemConfigCreate) (res *RespSystemConfigItem, code int, err error) {
	obj := model.NewSystemConfig()
	err = copier.Copy(&obj, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	obj, err = model.CreateSystemConfig(model.GetDB(), obj)
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

// Update 系统配置项更新
func (s *systemConfigService) Update(payload *ReqSystemConfigUpdate) (res *RespSystemConfigItem, code int, err error) {
	obj, err := model.GetSystemConfigInstance(
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
	err = model.UpdateSystemConfig(model.GetDB(), obj)
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

// Delete 系统配置项删除
func (s *systemConfigService) Delete(payload *ReqSystemConfigDelete) (code int, e error) {
	_, err := model.GetSystemConfigInstance(
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
	err = model.DeleteSystemConfig(model.GetDB(), payload.ID)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemConfig 系统配置项获取所有有效数据
func (s *systemConfigService) GetAllSystemConfig() ([]*model.SystemConfig, error) {
	result, err := model.GetAllSystemConfig(model.GetDB())
	if err != nil {
		return nil, err
	}
	return result, err
}
