package systemorg

import (
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IService interface {
	Create(payload *ReqSystemOrgCreate) (*model.SystemOrg, int, error)
	GetInstance(id uint64) (*model.SystemOrg, error)
	GetAllInstances() ([]*model.SystemOrg, int, error)
	Update(payload *ReqSystemOrgUpdate) (int, error)
	Delete(ids []uint64) error
	GetTreeData() (*RespSystemOrgTree, error)
}

type service struct {
}

func NewService() IService {
	return &service{}
}

func (s *service) Create(payload *ReqSystemOrgCreate) (*model.SystemOrg, int, error) {
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

func (s *service) GetInstance(id uint64) (*model.SystemOrg, error) {
	return model.GetOrg(model.GetDB(), id)
}

func (s *service) GetAllInstances() ([]*model.SystemOrg, int, error) {
	orgs, err := model.GetAllOrgs(model.GetDB())
	if err != nil {
		return nil, 0, err
	}
	return orgs, len(orgs), err
}

func (s *service) Update(payload *ReqSystemOrgUpdate) (int, error) {
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

func (s *service) Delete(ids []uint64) error {
	return model.DeleteOrgs(model.GetDB(), ids)
}

func (s *service) ConvertOrgTreeToJSON(org *model.SystemOrg, parent *model.SystemOrg) *RespSystemOrgTree {
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

func (s *service) GetTreeData() (*RespSystemOrgTree, error) {
	orgs, _, err := s.GetAllInstances()
	if err != nil {
		return nil, err
	}
	root := model.BuildOrgTree(orgs)
	result := s.ConvertOrgTreeToJSON(root, nil)
	return result, nil
}
