package systemrole

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemRoleService interface {
	List(req *ReqSystemRoleList) (res []*RespSystemRoleItem, total int64, code int, err error)
	Detail(req *ReqSystemRoleDetail) (res *RespSystemRoleItem, code int, err error)
	Create(payload *ReqSystemRoleCreate) (*RespSystemRoleItem, int, error)
	Update(payload *ReqSystemRoleUpdate) (res *RespSystemRoleItem, code int, err error)
	Delete(payload *ReqSystemRoleDelete) (int, error)
	GetAllSystemRole() ([]*model.SystemRole, int, error)
}

// systemRoleService 角色管理服务实现类
type systemRoleService struct{}

// NewSystemRoleService 初始化
func NewSystemRoleService() ISystemRoleService {
	return &systemRoleService{}
}

// List 角色管理列表
func (s *systemRoleService) List(req *ReqSystemRoleList) (res []*RespSystemRoleItem, total int64, code int, err error) {
	// 分页信息
	limit := req.Limit
	offset := req.Limit * (req.Page - 1)
	// 查询
	query := model.GetDB().Model(&model.SystemRole{}).Where("delete_time = 0")
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Desc != "" {
		query = query.Where("desc like ?", "%"+req.Desc+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	// 总数
	err = query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemRole
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

// Detail 角色管理详情
func (s *systemRoleService) Detail(req *ReqSystemRoleDetail) (res *RespSystemRoleItem, code int, err error) {
	res = &RespSystemRoleItem{}
	obj, err := model.GetSystemRoleInstance(
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
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return
}

// Create 角色管理创建
func (s *systemRoleService) Create(payload *ReqSystemRoleCreate) (*RespSystemRoleItem, int, error) {
	obj := model.NewSystemRole()
	err := copier.Copy(&obj, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	obj, err = model.CreateSystemRole(model.GetDB(), obj)
	if err != nil {
		return nil, render.CreateError, err
	}
	res := &RespSystemRoleItem{}
	err = copier.Copy(&res, obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Update 角色管理更新
func (s *systemRoleService) Update(payload *ReqSystemRoleUpdate) (*RespSystemRoleItem, int, error) {
	obj, err := model.GetSystemRoleInstance(
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
	err = model.UpdateSystemRole(model.GetDB(), obj)
	if err != nil {
		return nil, render.UpdateError, err
	}
	res := &RespSystemRoleItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Delete 角色管理删除
func (s *systemRoleService) Delete(payload *ReqSystemRoleDelete) (int, error) {
	// 删除
	err := model.DeleteSystemRole(model.GetDB(), payload.IDs)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemRole 角色管理获取所有有效数据
func (s *systemRoleService) GetAllSystemRole() ([]*model.SystemRole, int, error) {
	result, err := model.GetAllSystemRole(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return result, render.OK, err
}
