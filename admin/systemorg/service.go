package systemorg

import (
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemOrgService interface {
	Create(payload *ReqSystemOrgCreate) (*model.SystemOrg, int, error)
	GetInstance(id int64) (*model.SystemOrg, error)
	GetAllInstances() ([]*model.SystemOrg, int, error)
	Update(payload *ReqSystemOrgUpdate) (int, error)
	Delete(ids []int64) error
	GetTreeData() (*RespSystemOrgTree, error)
}

type systemOrgService struct {
}

func NewService() ISystemOrgService {
	return &systemOrgService{}
}

func (s *systemOrgService) Create(payload *ReqSystemOrgCreate) (*model.SystemOrg, int, error) {
	org := model.NewSystemOrg()
	err := copier.Copy(&org, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}

	err = model.CreateOrg(model.GetDB(), org)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.CreateError, err
	}
	return org, render.OK, err
}

func (s *systemOrgService) GetInstance(id int64) (*model.SystemOrg, error) {
	return model.GetOrg(model.GetDB(), id)
}

func (s *systemOrgService) GetAllInstances() ([]*model.SystemOrg, int, error) {
	orgs, err := model.GetAllOrgs(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return orgs, len(orgs), err
}

func (s *systemOrgService) Update(payload *ReqSystemOrgUpdate) (int, error) {
	org, err := model.GetOrg(model.GetDB(), payload.ID)
	if err != nil {
		return render.DataNotExistError, err
	}
	err = copier.Copy(&org, &payload)
	if err != nil {
		log.GetLogger().Error(err)
		return render.DBAttributesCopyError, err
	}
	err = model.UpdateOrg(model.GetDB(), org)
	if err != nil {
		log.GetLogger().Error(err)
		return render.UpdateError, err
	}
	return render.OK, err
}

func (s *systemOrgService) Delete(ids []int64) error {
	return model.DeleteOrgs(model.GetDB(), ids)
}

func (s *systemOrgService) ConvertOrgTreeToJSON(org *model.SystemOrg, parent *model.SystemOrg) *RespSystemOrgTree {
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	result := &RespSystemOrgTree{
		ID:         org.ID,
		ParentID:   org.ParentID,
		ParentName: pName,
		Name:       org.Name,
		Manager:    org.Manager,
		Phone:      org.Phone,
		Children:   make([]*RespSystemOrgTree, len(org.Children)),
	}

	for i, child := range org.Children {
		result.Children[i] = s.ConvertOrgTreeToJSON(child, org)
	}
	return result
}

func (s *systemOrgService) GetTreeData() (*RespSystemOrgTree, error) {
	orgs, _, err := s.GetAllInstances()
	if err != nil {
		return nil, err
	}
	root := model.BuildOrgTree(orgs)
	result := s.ConvertOrgTreeToJSON(root, nil)
	return result, nil
}
