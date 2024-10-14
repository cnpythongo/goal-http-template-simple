package systemroleuser

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemRoleUserService interface {
	List(req *ReqSystemRoleUserList) ([]*RespSystemRoleUserItem, int64, int, error)
	Detail(req *ReqSystemRoleUserDetail) (*RespSystemRoleUserItem, int, error)
	Create(payload *ReqSystemRoleUserCreate) (int, error)
	Update(payload *ReqSystemRoleUserUpdate) (*RespSystemRoleUserItem, int, error)
	Delete(payload *ReqSystemRoleUserDelete) (int, error)
	GetAllSystemRoleUser() ([]*model.SystemRoleUser, int, error)
}

// systemRoleUserService 角色用户关联服务实现类
type systemRoleUserService struct{}

// NewSystemRoleUserService 初始化
func NewSystemRoleUserService() ISystemRoleUserService {
	return &systemRoleUserService{}
}

// List 角色用户关联列表
func (s *systemRoleUserService) List(req *ReqSystemRoleUserList) ([]*RespSystemRoleUserItem, int64, int, error) {
	// 查询
	query := model.GetDB().Model(&model.SystemRoleUser{})
	if req.OrgId > 0 {
		query = query.Where("org_id = ?", req.OrgId)
	}
	if req.RoleId > 0 {
		query = query.Where("role_id = ?", req.RoleId)
	}
	if req.UserId > 0 {
		query = query.Where("user_id = ?", req.UserId)
	}
	// 总数
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemRoleUser
	if req.Page > 0 && req.Limit > 0 {
		// 分页信息
		limit := req.Page
		offset := req.Limit * (req.Page - 1)
		query = query.Limit(limit).Offset(offset)
	}
	err = query.Order("id desc").Find(&objs).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	res := make([]*RespSystemRoleUserItem, 0)
	err = copier.Copy(&res, objs)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.DBAttributesCopyError, err
	}
	return res, total, render.OK, nil
}

// Detail 角色用户关联详情
func (s *systemRoleUserService) Detail(req *ReqSystemRoleUserDetail) (*RespSystemRoleUserItem, int, error) {
	obj, err := model.GetSystemRoleUserInstance(
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

	res := &RespSystemRoleUserItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Create 角色用户关联创建
func (s *systemRoleUserService) Create(payload *ReqSystemRoleUserCreate) (int, error) {
	err := model.CreateSystemRoleUsersByRoleId(model.GetDB(), payload.OrgId, payload.RoleId, payload.UserIds)
	if err != nil {
		return render.CreateError, err
	}
	return render.OK, nil
}

// Update 角色用户关联更新
func (s *systemRoleUserService) Update(payload *ReqSystemRoleUserUpdate) (*RespSystemRoleUserItem, int, error) {
	obj, err := model.GetSystemRoleUserInstance(
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
	err = model.UpdateSystemRoleUser(model.GetDB(), obj)
	if err != nil {
		return nil, render.UpdateError, err
	}

	res := &RespSystemRoleUserItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Delete 角色用户关联删除
func (s *systemRoleUserService) Delete(payload *ReqSystemRoleUserDelete) (int, error) {
	// 删除
	err := model.DeleteSystemRoleUser(model.GetDB(), payload.IDs)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemRoleUser 角色用户关联获取所有有效数据
func (s *systemRoleUserService) GetAllSystemRoleUser() ([]*model.SystemRoleUser, int, error) {
	result, err := model.GetAllSystemRoleUser(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return result, render.OK, err
}
