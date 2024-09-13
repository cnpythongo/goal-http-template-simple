package model

import "gorm.io/gorm"

type SystemOrg struct {
	BaseModel
	ParentID uint64 `json:"parent_id" gorm:"column:parent_id;type:int(11);null;default:null;comment:父节点ID"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:部门名称"`

	Children []*SystemOrg `gorm:"foreignKey:parent_id;references:id" json:"children,omitempty"`
}

func NewSystemOrg() *SystemOrg {
	return &SystemOrg{}
}

func GetOrg(db *gorm.DB, id uint64) (*SystemOrg, error) {
	var org SystemOrg
	err := db.First(&org, id).Error
	if err != nil {
		return nil, err
	}
	return &org, nil
}

func CreateOrg(db *gorm.DB, org SystemOrg) error {
	return db.Create(&org).Error
}

func GetAllOrgs(db *gorm.DB) ([]*SystemOrg, error) {
	var result []*SystemOrg
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateOrg(db *gorm.DB, org *SystemOrg) error {
	return db.Save(org).Error
}

func DeleteOrg(db *gorm.DB, id uint64) error {
	var org SystemOrg
	return db.Delete(&org, id).Error
}

func BuildOrgTree(orgs []*SystemOrg) *SystemOrg {
	rootNode := NewSystemOrg()

	orgMap := make(map[uint64]*SystemOrg)
	for _, org := range orgs {
		org.Children = make([]*SystemOrg, 0)
		orgMap[org.ID] = org
	}

	for _, org := range orgs {
		if org.ParentID == 0 {
			rootNode = org
		} else {
			parent, ok := orgMap[org.ParentID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, org)
			}
		}
	}

	return rootNode
}
