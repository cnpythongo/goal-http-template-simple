package systemmenu

import (
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemMenuService interface {
	Create(payload *ReqSystemMenuCreate) (*model.SystemMenu, int, error)
	GetInstance(id uint64) (*model.SystemMenu, error)
	GetAllInstances() ([]*model.SystemMenu, int, error)
	Update(payload *ReqSystemMenuUpdate) error
	Delete(ids []uint64) error
	BuildTree() (*RespSystemMenuTree, error)
}

type systemMenuService struct {
}

func NewSystemMenuService() ISystemMenuService {
	return &systemMenuService{}
}

func (s *systemMenuService) Create(payload *ReqSystemMenuCreate) (*model.SystemMenu, int, error) {
	menu := model.NewSystemMenu()
	err := copier.Copy(menu, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	err = model.CreateSystemMenu(model.GetDB(), menu)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.CreateError, err
	}
	return menu, render.OK, err
}

func (s *systemMenuService) GetInstance(id uint64) (*model.SystemMenu, error) {
	return model.GetSystemMenuById(model.GetDB(), id)
}

func (s *systemMenuService) GetAllInstances() ([]*model.SystemMenu, int, error) {
	result, err := model.GetAllSystemMenus(model.GetDB())
	if err != nil {
		return nil, 0, err
	}
	return result, len(result), err
}

func (s *systemMenuService) Update(payload *ReqSystemMenuUpdate) error {
	obj, err := model.GetSystemMenuById(model.GetDB(), payload.ID)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"name":      payload.Name,
		"parent_id": payload.ParentID,
	}
	return model.UpdateSystemMenu(model.GetDB(), obj.ID, data)
}

func (s *systemMenuService) Delete(ids []uint64) error {
	return model.DeleteOrgs(model.GetDB(), ids)
}

func (s *systemMenuService) ConvertTreeToJSON(menu *model.SystemMenu, parent *model.SystemMenu) *RespSystemMenuTree {
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	result := &RespSystemMenuTree{
		ID:         menu.ID,
		ParentID:   menu.ParentID,
		ParentName: pName,
		Children:   make([]*RespSystemMenuTree, len(menu.Children)),
	}

	for i, child := range menu.Children {
		result.Children[i] = s.ConvertTreeToJSON(child, menu)
	}
	return result
}

func (s *systemMenuService) BuildTree() (*RespSystemMenuTree, error) {
	orgs, _, err := s.GetAllInstances()
	if err != nil {
		return nil, err
	}
	root := model.BuildSystemMenuTree(orgs)
	result := s.ConvertTreeToJSON(root, nil)
	return result, nil
}
