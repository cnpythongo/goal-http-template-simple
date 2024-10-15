package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

type SystemRoleUser struct {
	BaseModel
	OrgID  int64 `json:"org_id" gorm:"column:org_id;type:int(11);comment:组织机构ID"`
	RoleID int64 `json:"role_id" gorm:"column:role_id;type:int(11);comment:角色ID"`
	UserID int64 `json:"user_id" gorm:"column:user_id;type:int(11);comment:用户ID"`
}

func (m *SystemRoleUser) TableName() string {
	return "system_role_users"
}

func NewSystemRoleUser() *SystemRoleUser {
	return &SystemRoleUser{}
}

func NewSystemRoleUserList() []*SystemRoleUser {
	return make([]*SystemRoleUser, 0)
}

func (m *SystemRoleUser) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func CreateSystemRoleUser(tx *gorm.DB, obj *SystemRoleUser) (*SystemRoleUser, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleUser.CreateSystemRoleUser Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func CreateSystemRoleUsersByRoleId(db *gorm.DB, orgId int64, roleId int64, userIds []int64) error {
	tx := db.Begin()
	items := NewSystemRoleUserList()
	for _, uid := range userIds {
		items = append(items, &SystemRoleUser{
			OrgID:  orgId,
			RoleID: roleId,
			UserID: uid,
		})
	}

	err := DeleteSystemRoleUsersByRoleId(tx, roleId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(userIds) > 0 {
		err = tx.Create(&items).Error
		if err != nil {
			tx.Rollback()
			log.GetLogger().Errorf("model.SystemRoleMenu.CreateSystemRoleUsersByRoleId Error ==> %v", err)
			return err
		}
	}
	tx.Commit()
	return nil
}

func UpdateSystemRoleUser(tx *gorm.DB, obj *SystemRoleUser) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleUser.UpdateSystemRoleUser Error ==> %v", err)
	}
	return err
}

func DeleteSystemRoleUser(tx *gorm.DB, ids []int64) error {
	err := tx.Model(NewSystemRoleUser()).Where("id in ?", ids).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleUser.DeleteSystemRoleUser Error ==> %v", err)
	}
	return err
}

func DeleteSystemRoleUsersByRoleId(tx *gorm.DB, roleId int64) error {
	err := tx.Where("role_id = ?", roleId).Delete(NewSystemRoleUser()).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleMenu.DeleteSystemRoleUsersByRoleId Error ==> %v", err)
	}
	return err
}

func GetSystemRoleUserInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemRoleUser, error) {
	result := NewSystemRoleUser()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemRoleUser.GetSystemRoleUserInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemRoleUser.GetSystemRoleUserInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemRoleUserList(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*SystemRoleUser, int64, error) {
	qs := tx.Model(NewSystemRoleUser()).Where("delete_time = 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleUser.GetSystemRoleUserList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewSystemRoleUserList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemRoleUser.GetSystemRoleUserList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemRoleUser(tx *gorm.DB, conditions map[string]interface{}) ([]*SystemRoleUser, error) {
	result := NewSystemRoleUserList()
	query := tx.Where("delete_time = 0")
	if conditions != nil && len(conditions) > 0 {
		query = query.Where(conditions)
	}
	err := query.Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemRoleUser.GetAllSystemRoleUser Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}
