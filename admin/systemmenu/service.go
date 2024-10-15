package systemmenu

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"time"
)

type ISystemMenuService interface {
	List(req *ReqSystemMenuList) ([]*RespSystemMenuItem, int64, int, error)
	Detail(req *ReqSystemMenuDetail) (res *RespSystemMenuItem, code int, err error)
	Create(payload *ReqSystemMenuCreate) (*RespSystemMenuItem, int, error)
	Update(payload *ReqSystemMenuUpdate) (*RespSystemMenuItem, int, error)
	Delete(payload *ReqSystemMenuDelete) (code int, e error)
	GetAllSystemMenu(conditions map[string]interface{}) ([]*model.SystemMenu, int, error)
	Tree(req *ReqSystemMenuTree) ([]*RespSystemMenuTree, int, error)
	ConvertSystemMenuTreeToJSON(nodes []*model.SystemMenu, parent *model.SystemMenu) ([]*RespSystemMenuTree, int, error)
}

// systemMenuService 菜单管理服务实现类
type systemMenuService struct{}

// NewSystemMenuService 初始化
func NewSystemMenuService() ISystemMenuService {
	return &systemMenuService{}
}

// List 菜单管理列表
func (s *systemMenuService) List(req *ReqSystemMenuList) ([]*RespSystemMenuItem, int64, int, error) {
	// 分页信息
	limit := req.Page
	offset := req.Limit * (req.Page - 1)
	// 查询
	query := model.GetDB().Model(&model.SystemMenu{})
	if req.ParentID >= 0 {
		query = query.Where("parent_id = ?", req.ParentID)
	}
	if req.Kind != "" {
		query = query.Where("kind = ?", req.Kind)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Icon != "" {
		query = query.Where("icon = ?", req.Icon)
	}
	if req.Sort >= 0 {
		query = query.Where("sort = ?", req.Sort)
	}
	if req.AuthTag != "" {
		query = query.Where("auth_tag = ?", req.AuthTag)
	}
	if req.Route != "" {
		query = query.Where("route = ?", req.Route)
	}
	if req.Component != "" {
		query = query.Where("component = ?", req.Component)
	}
	if req.Params != "" {
		query = query.Where("params = ?", req.Params)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}
	// 总数
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	// 数据
	var objs []*model.SystemMenu
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}

	res := make([]*RespSystemMenuItem, 0)
	err = copier.Copy(&res, objs)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, total, render.DBAttributesCopyError, err
	}
	return res, total, render.OK, nil
}

// Detail 菜单管理详情
func (s *systemMenuService) Detail(req *ReqSystemMenuDetail) (res *RespSystemMenuItem, code int, err error) {
	// var obj *model.SystemMenu
	obj, err := model.GetSystemMenuInstance(
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

// Create 菜单管理创建
func (s *systemMenuService) Create(payload *ReqSystemMenuCreate) (*RespSystemMenuItem, int, error) {
	obj := model.NewSystemMenu()
	err := copier.Copy(&obj, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	obj, err = model.CreateSystemMenu(model.GetDB(), obj)
	if err != nil {
		return nil, render.CreateError, err
	}
	res := &RespSystemMenuItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Update 菜单管理更新
func (s *systemMenuService) Update(payload *ReqSystemMenuUpdate) (*RespSystemMenuItem, int, error) {
	obj, err := model.GetSystemMenuInstance(
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
	err = model.UpdateSystemMenu(model.GetDB(), obj)
	if err != nil {
		return nil, render.UpdateError, err
	}
	res := &RespSystemMenuItem{}
	err = copier.Copy(&res, &obj)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

// Delete 菜单管理删除
func (s *systemMenuService) Delete(payload *ReqSystemMenuDelete) (int, error) {
	// 删除
	_, total, err := model.GetSystemMenuList(model.GetDB(), 0, 0, "parent_id in ?", []interface{}{payload.IDs})
	if err != nil {
		return render.QueryError, err
	}
	if total > 0 {
		return render.AssociatedDataExistsError, errors.New("存在子菜单，无法删除")
	}

	err = model.DeleteSystemMenu(model.GetDB(), payload.IDs)
	if err != nil {
		return render.DeleteError, err
	}
	return render.OK, nil
}

// GetAllSystemMenu 菜单管理获取所有有效数据
func (s *systemMenuService) GetAllSystemMenu(conditions map[string]interface{}) ([]*model.SystemMenu, int, error) {
	result, err := model.GetAllSystemMenu(model.GetDB(), conditions)
	if err != nil {
		return nil, render.QueryError, err
	}
	return result, render.OK, err
}

// Tree 菜单管理树
func (s *systemMenuService) Tree(req *ReqSystemMenuTree) ([]*RespSystemMenuTree, int, error) {
	conditions := make(map[string]interface{})
	if req.Status != "" {
		conditions["status"] = req.Status
	}
	rows, code, err := s.GetAllSystemMenu(conditions)
	if err != nil {
		return nil, code, err
	}
	roots := model.BuildSystemMenuTree(rows)
	result, code, err := s.ConvertSystemMenuTreeToJSON(roots, nil)
	return result, code, err
}

// ConvertSystemMenuTreeToJSON 行模型数据转成JSON树结构
func (s *systemMenuService) ConvertSystemMenuTreeToJSON(nodes []*model.SystemMenu, parent *model.SystemMenu) ([]*RespSystemMenuTree, int, error) {
	result := make([]*RespSystemMenuTree, 0)
	if len(nodes) == 0 {
		return result, render.OK, nil
	}
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	for _, node := range nodes {
		children, code, err := s.ConvertSystemMenuTreeToJSON(node.Children, node)
		if err != nil {
			return nil, code, err
		}
		item := &RespSystemMenuTree{
			ParentName: pName,
			Children:   children,
		}
		err = copier.Copy(&item, &node)
		if err != nil {
			log.GetLogger().Error(err)
			return nil, render.DBAttributesCopyError, err
		}
		result = append(result, item)
	}
	return result, render.OK, nil
}
