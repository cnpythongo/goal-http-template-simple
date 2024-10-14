package systemrolemenu

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemRoleMenuService interface {
	List(req *ReqSystemRoleMenuList) ([]*RespSystemRoleMenuItem, int64, int, error)
	Detail(req *ReqSystemRoleMenuDetail) (*RespSystemRoleMenuItem, int, error)
	Create(payload *ReqSystemRoleMenuCreate) (int, error)
	Update(payload *ReqSystemRoleMenuUpdate) (*RespSystemRoleMenuItem, int, error)
	Delete(payload *ReqSystemRoleMenuDelete) (int, error)
	GetAllSystemRoleMenu() ([]*model.SystemRoleMenu, int, error)
}

// systemRoleMenuService 角色菜单关联服务实现类
type systemRoleMenuService struct{}

// NewSystemRoleMenuService 初始化
func NewSystemRoleMenuService() ISystemRoleMenuService {
	return &systemRoleMenuService{}
}

// List 角色菜单关联列表
func (s *systemRoleMenuService) List(req *ReqSystemRoleMenuList) ([]*RespSystemRoleMenuItem, int64, int, error) {
	// 查询
	query := model.GetDB().Model(&model.SystemRoleMenu{})
	if req.OrgId > 0 {
		query = query.Where("org_id = ?", req.OrgId)
	}
	if req.RoleId >= 0 {
		query = query.Where("role_id = ?", req.RoleId)
	}
	if req.MenuId > 0 {
		query = query.Where("menu_id = ?", req.MenuId)
	}
	// 总数
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemRoleMenu
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

	res := make([]*RespSystemRoleMenuItem, 0)
	err = copier.Copy(&res, objs)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.DBAttributesCopyError, err
	}
	return res, total, render.OK, nil
}

// Detail 角色菜单关联详情
func (s *systemRoleMenuService) Detail(req *ReqSystemRoleMenuDetail) (*RespSystemRoleMenuItem, int, error) {
	obj, err := model.GetSystemRoleMenuInstance(
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

	res := &RespSystemRoleMenuItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Create 角色菜单关联创建
func (s *systemRoleMenuService) Create(payload *ReqSystemRoleMenuCreate) (int, error) {
	err := model.CreateSystemRoleMenuByRoleId(model.GetDB(), payload.OrgId, payload.RoleId, payload.MenuIds)
	if err != nil {
		return render.CreateError, err
	}
	return render.OK, nil
}

// Update 角色菜单关联更新
func (s *systemRoleMenuService) Update(payload *ReqSystemRoleMenuUpdate) (*RespSystemRoleMenuItem, int, error) {
	obj, err := model.GetSystemRoleMenuInstance(
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
	err = model.UpdateSystemRoleMenu(model.GetDB(), obj)
	if err != nil {
		return nil, render.UpdateError, err
	}

	res := &RespSystemRoleMenuItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Delete 角色菜单关联删除
func (s *systemRoleMenuService) Delete(payload *ReqSystemRoleMenuDelete) (int, error) {
	// 删除
	err := model.DeleteSystemRoleMenu(model.GetDB(), payload.IDs)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemRoleMenu 角色菜单关联获取所有有效数据
func (s *systemRoleMenuService) GetAllSystemRoleMenu() ([]*model.SystemRoleMenu, int, error) {
	result, err := model.GetAllSystemRoleMenu(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return result, render.OK, err
}
