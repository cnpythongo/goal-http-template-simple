package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type SystemRoleMenu struct {
	BaseModel
	OrgID  int64 `json:"org_id" gorm:"column:org_id;type:int(11);comment:组织机构ID"` // 组织机构ID
	RoleID int64 `json:"role_id" gorm:"column:role_id;type:int(11);comment:角色ID"` // 角色ID
	MenuID int64 `json:"menu_id" gorm:"column:menu_id;type:int(11);comment:菜单ID"` // 菜单ID
}

func (m *SystemRoleMenu) TableName() string {
	return "system_role_menus"
}

func NewSystemRoleMenu() *SystemRoleMenu {
	return &SystemRoleMenu{}
}

func NewSystemRoleMenuList() []*SystemRoleMenu {
	return make([]*SystemRoleMenu, 0)
}

func (m *SystemRoleMenu) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func CreateSystemRoleMenu(tx *gorm.DB, obj *SystemRoleMenu) (*SystemRoleMenu, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.CreateSystemRoleMenu Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func UpdateSystemRoleMenu(tx *gorm.DB, obj *SystemRoleMenu) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.UpdateSystemRoleMenu Error ==> %v", err)
	}
	return err
}

func DeleteSystemRoleMenu(tx *gorm.DB, ids []int64) error {
	err := tx.Model(NewSystemRoleMenu()).Where("id in ?", ids).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.DeleteSystemRoleMenu Error ==> %v", err)
	}
	return err
}

func GetSystemRoleMenuInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemRoleMenu, error) {
	result := NewSystemRoleMenu()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemRoleMenu.GetSystemRoleMenuInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemRoleMenu.GetSystemRoleMenuInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemRoleMenuList(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*SystemRoleMenu, int64, error) {
	qs := tx.Model(NewSystemRoleMenu()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.GetSystemRoleMenuList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewSystemRoleMenuList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.GetSystemRoleMenuList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemRoleMenu(tx *gorm.DB) ([]*SystemRoleMenu, error) {
	result := NewSystemRoleMenuList()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemRoleMenu.GetAllSystemRoleMenu Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}
