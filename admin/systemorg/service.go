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
	GetAllSystemOrgs() ([]*model.SystemOrg, int, error)
	Update(payload *ReqSystemOrgUpdate) (int, error)
	Delete(ids []int64) error
	Tree() ([]*RespSystemOrgTree, int, error)
	ConvertOrgTreeToJSON(nodes []*model.SystemOrg, parent *model.SystemOrg) ([]*RespSystemOrgTree, int, error)
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

func (s *systemOrgService) GetAllSystemOrgs() ([]*model.SystemOrg, int, error) {
	orgs, err := model.GetAllOrgs(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return orgs, render.OK, err
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

func (s *systemOrgService) ConvertOrgTreeToJSON(nodes []*model.SystemOrg, parent *model.SystemOrg) ([]*RespSystemOrgTree, int, error) {
	result := make([]*RespSystemOrgTree, 0)
	if len(nodes) == 0 {
		return result, render.OK, nil
	}
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	for _, node := range nodes {
		children, code, err := s.ConvertOrgTreeToJSON(node.Children, node)
		if err != nil {
			return nil, code, err
		}
		item := &RespSystemOrgTree{
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

func (s *systemOrgService) Tree() ([]*RespSystemOrgTree, int, error) {
	rows, code, err := s.GetAllSystemOrgs()
	if err != nil {
		return nil, code, err
	}
	roots := model.BuildOrgTree(rows)
	result, code, err := s.ConvertOrgTreeToJSON(roots, nil)
	return result, code, err
}
