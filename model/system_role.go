package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type SystemRole struct {
	BaseModel
	OrgID     int64  `json:"org_id" gorm:"column:org_id;type:int(11);comment:组织机构ID"`                            // 组织机构ID
	Name      string `json:"name" gorm:"column:name;type:varchar(200);comment:角色名称"`                             // 角色名称
	Desc      string `json:"desc" gorm:"column:desc;type:varchar(200);default:'';comment:角色描述"`                  // 角色描述
	Status    int    `json:"status" gorm:"column:status;type:int(11);default:1;comment:角色状态, 0-禁用, 1-启用"`        // 角色状态, 0-禁用, 1-启用
	IsDeleted int    `json:"is_deleted" gorm:"column:is_deleted;type:int(11);default:0;comment:是否被删除, 0-否, 1-是"` // 是否被删除, 0-否, 1-是
}

func (t *SystemRole) TableName() string {
	return "system_roles"
}

func NewSystemRole() *SystemRole {
	return &SystemRole{}
}

func NewSystemRoleList() []*SystemRole {
	return make([]*SystemRole, 0)
}

func (m *SystemRole) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func CreateSystemRole(tx *gorm.DB, obj *SystemRole) (*SystemRole, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRole.CreateSystemRole Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func UpdateSystemRole(tx *gorm.DB, obj *SystemRole) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRole.UpdateSystemRole Error ==> %v", err)
	}
	return err
}

func DeleteSystemRole(tx *gorm.DB, ids []int64) error {
	err := tx.Model(NewSystemRole()).Where("id in ?", ids).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRole.DeleteSystemRole Error ==> %v", err)
	}
	return err
}

func GetSystemRoleInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemRole, error) {
	result := NewSystemRole()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemRole.GetSystemRoleInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemRole.GetSystemRoleInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemRoleList(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*SystemRole, int64, error) {
	qs := tx.Model(NewSystemRole()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRole.GetSystemRoleList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewSystemRoleList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRole.GetSystemRoleList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemRole(tx *gorm.DB) ([]*SystemRole, error) {
	result := NewSystemRoleList()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemRole.GetAllSystemRole Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}
