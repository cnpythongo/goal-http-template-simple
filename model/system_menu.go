package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"time"
)

// SystemMenu 系统菜单
type SystemMenu struct {
	BaseModel
	ParentID  int64         `json:"parent_id" gorm:"column:parent_id;type:int(11);not null;default:0;comment:上级菜单ID"`
	Kind      string        `json:"kind" gorm:"column:kind;type:varchar(100);not null;default:'';comment:权限类型: dir=目录，menu=菜单，button=按钮"`
	Name      string        `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:菜单名称"`
	Icon      string        `json:"icon" gorm:"column:icon;type:varchar(100);not null;default:'';comment:菜单图标"`
	Sort      int64         `json:"sort" gorm:"column:sort;type:smallint(5);not null;default:0;comment:菜单排序"`
	AuthTag   string        `json:"auth_tag" gorm:"column:auth_tag;type:varchar(200);not null;default:'';comment:权限标识"`
	Route     string        `json:"route" gorm:"column:route;type:varchar(500);not null;default:'';comment:路由地址"`
	Component string        `json:"component" gorm:"column:component;type:varchar(500);not null;default:'';comment:前端组件"`
	Params    string        `json:"params" gorm:"column:params;type:varchar(500);not null;default:'';comment:路由参数"`
	Selected  string        `json:"selected" gorm:"column:selected;type:varchar(500);not null;default:'';comment:选中菜单"`
	Status    string        `json:"status" gorm:"column:status;type:varchar(20);not null;default:'enable';comment:状态: disable=停用, enable=启用"`
	Children  []*SystemMenu `gorm:"-" json:"children,omitempty"`
}

func (m *SystemMenu) TableName() string {
	return "system_menus"
}

func NewSystemMenu() *SystemMenu {
	return &SystemMenu{}
}

func NewSystemMenuList() []*SystemMenu {
	return make([]*SystemMenu, 0)
}

func (m *SystemMenu) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = int64(now)
	m.UpdateTime = int64(now)
	return nil
}

func GetSystemMenuByName(tx *gorm.DB, name string) (*SystemMenu, error) {
	obj := NewSystemMenu()
	err := tx.Model(NewSystemMenu()).Where("name = ?", name).First(&obj).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.GetLogger().Errorf("model.SystemMenu.GetSystemMenuByName Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func CreateSystemMenu(tx *gorm.DB, obj *SystemMenu) (*SystemMenu, error) {
	err := tx.Create(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemMenu.CreateSystemMenu Error ==> %v", err)
		return nil, err
	}
	return obj, nil
}

func UpdateSystemMenu(tx *gorm.DB, obj *SystemMenu) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemMenu.UpdateSystemMenu Error ==> %v", err)
	}
	return err
}

func DeleteSystemMenu(db *gorm.DB, ids []int64) error {
	tx := db.Begin()
	err := tx.Model(NewSystemMenu()).Where("id in ?", ids).UpdateColumns(map[string]interface{}{
		"delete_time": time.Now().Unix(),
	}).Error
	if err != nil {
		tx.Rollback()
		log.GetLogger().Errorf("model.SystemMenu.DeleteSystemMenu Error ==> %v", err)
	}

	err = tx.Where("menu_id in ?", ids).Delete(NewSystemRoleMenu()).Error
	if err != nil {
		tx.Rollback()
		log.GetLogger().Errorf("model.SystemMenu.DeleteSystemMenu.DeleteSystemRoleMenu Error ==> %v", err)
	}
	tx.Commit()
	return err
}

func GetSystemMenuInstance(tx *gorm.DB, conditions map[string]interface{}) (*SystemMenu, error) {
	result := NewSystemMenu()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.SystemMenu.GetSystemMenuInstance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.SystemMenu.GetSystemMenuInstance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func GetSystemMenuList(tx *gorm.DB, page, limit int, query interface{}, args []interface{}) ([]*SystemMenu, int64, error) {
	qs := tx.Model(NewSystemMenu()).Where("delete_time = 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemMenu.GetSystemMenuList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		qs = qs.Limit(limit).Offset(offset)
	}
	result := NewSystemMenuList()
	err = qs.Order("sort desc").Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.SystemMenu.GetSystemMenuList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAllSystemMenu(tx *gorm.DB, conditions map[string]interface{}) ([]*SystemMenu, error) {
	result := NewSystemMenuList()
	query := tx.Where("delete_time = 0")
	if conditions != nil && len(conditions) > 0 {
		query = query.Where(conditions)
	}
	err := query.Order("sort desc").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.SystemMenu.GetAllSystemMenu Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func BuildSystemMenuTree(rows []*SystemMenu) []*SystemMenu {
	rootNodes := NewSystemMenuList()

	tmpMap := make(map[int64]*SystemMenu)
	for _, r := range rows {
		r.Children = make([]*SystemMenu, 0)
		tmpMap[r.ID] = r
	}

	for _, r := range rows {
		if r.ParentID == 0 {
			rootNodes = append(rootNodes, r)
		} else {
			parent, ok := tmpMap[r.ParentID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, r)
			}
		}
	}

	return rootNodes
}
