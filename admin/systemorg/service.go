package systemorg

import (
	"goal-app/model"
)

type ISystemOrgService interface {
	CreateOrg(payload ReqSystemOrgCreate) error
	GetOrg(id uint64) (*model.SystemOrg, error)
	GetAllOrgs() ([]*model.SystemOrg, int, error)
	UpdateOrg(payload *ReqSystemOrgUpdate) error
	DeleteOrg(ids []uint64) error
	BuildOrgTree() (*RespSystemOrgTree, error)
}

type systemOrgService struct {
}

func NewSystemOrgService() ISystemOrgService {
	return &systemOrgService{}
}

func (s *systemOrgService) CreateOrg(payload ReqSystemOrgCreate) error {
	org := model.SystemOrg{
		ParentID: payload.ParentID,
		Name:     payload.Name,
		Manager:  payload.Manager,
		Phone:    payload.Phone,
	}
	return model.CreateOrg(model.GetDB(), org)
}

func (s *systemOrgService) GetOrg(id uint64) (*model.SystemOrg, error) {
	return model.GetOrg(model.GetDB(), id)
}

func (s *systemOrgService) GetAllOrgs() ([]*model.SystemOrg, int, error) {
	orgs, err := model.GetAllOrgs(model.GetDB())
	if err != nil {
		return nil, 0, err
	}
	return orgs, len(orgs), err
}

func (s *systemOrgService) UpdateOrg(payload *ReqSystemOrgUpdate) error {
	org, err := model.GetOrg(model.GetDB(), payload.ID)
	if err != nil {
		return err
	}
	org.ParentID = payload.ParentID
	org.Name = payload.Name
	org.Manager = payload.Manager
	org.Phone = payload.Phone
	return model.UpdateOrg(model.GetDB(), org)
}

func (s *systemOrgService) DeleteOrg(ids []uint64) error {
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

func (s *systemOrgService) BuildOrgTree() (*RespSystemOrgTree, error) {
	orgs, _, err := s.GetAllOrgs()
	if err != nil {
		return nil, err
	}
	root := model.BuildOrgTree(orgs)
	result := s.ConvertOrgTreeToJSON(root, nil)
	return result, nil
}
