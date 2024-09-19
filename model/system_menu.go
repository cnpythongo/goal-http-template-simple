package model

import (
	"gorm.io/gorm"
)

// SystemMenu 系统菜单
type SystemMenu struct {
	BaseModel
	ParentID  uint64 `json:"parent_id" gorm:"column:parent_id;type:int(11);not null;default:0;comment:'上级菜单ID'"`
	Kind      string `json:"kind" gorm:"column:kind;type:varchar(100);not null;default:'';comment:'权限类型: dir=目录，menu=菜单，button=按钮''"`
	Name      string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:'菜单名称'"`
	Icon      string `json:"icon" gorm:"column:icon;type:varchar(100);not null;default:'';comment:'菜单图标'"`
	Sort      uint16 `json:"sort" gorm:"column:sort;type:smallint(5);not null;default:0;comment:'菜单排序'"`
	AuthTag   string `json:"auth_tag" gorm:"column:auth_tag;type:varchar(200);not null;default:'';comment:'权限标识'"`
	Route     string `json:"route" gorm:"column:route;type:varchar(500);not null;default:'';comment:'路由地址'"`
	Component string `json:"component" gorm:"column:component;type:varchar(500);not null;default:'';comment:'前端组件'"`
	Params    string `json:"params" gorm:"column:params;type:varchar(500);not null;default:'';comment:'路由参数'"`
	Status    string `json:"status" gorm:"column:status;type:varchar(20);not null;default:'enable';comment:'状态: disable=停用, enable=启用'"`

	Children []*SystemMenu `gorm:"foreignKey:parent_id;references:id" json:"children,omitempty"`
}

func NewSystemMenu() *SystemMenu {
	return &SystemMenu{}
}

func NewSystemMenuList() []*SystemMenu {
	return make([]*SystemMenu, 0)
}

func (s *SystemMenu) TableName() string {
	return "system_menus"
}

func (s *SystemMenu) BeforeCreateSystemMenu(db *gorm.DB) error {
	return nil
}

func CreateSystemMenu(db *gorm.DB, menu *SystemMenu) error {
	return db.Create(&menu).Error
}

func UpdateSystemMenu(db *gorm.DB, id uint64, data map[string]interface{}) error {
	return db.Model(&SystemMenu{}).Where("id = ?", id).Updates(data).Error
}

func DeleteSystemMenus(db *gorm.DB, ids []uint64) error {
	return db.Where("id IN (?)", ids).Delete(&SystemMenu{}).Error
}

func GetSystemMenuById(db *gorm.DB, id uint64) (*SystemMenu, error) {
	var menu SystemMenu
	return &menu, db.Where("id = ?", id).First(&menu).Error
}

func GetSystemMenusByParentId(db *gorm.DB, parentId uint64) ([]*SystemMenu, error) {
	result := NewSystemMenuList()
	return result, db.Where("parent_id = ?", parentId).Order("sort desc").Find(&result).Error
}

func GetSystemMenuByName(db *gorm.DB, name string) (*SystemMenu, error) {
	var menu SystemMenu
	return &menu, db.Where("name = ?", name).First(&menu).Error
}

func GetAllSystemMenus(db *gorm.DB) ([]*SystemMenu, error) {
	result := NewSystemMenuList()
	return result, db.Order("sort desc").Find(&result).Error
}

func BuildSystemMenuTree(rows []*SystemMenu) *SystemMenu {
	rootNode := NewSystemMenu()

	mapping := make(map[uint64]*SystemMenu)
	for _, item := range rows {
		item.Children = make([]*SystemMenu, 0)
		mapping[item.ID] = item
	}

	for _, item := range rows {
		if item.ParentID == 0 {
			rootNode = item
		} else {
			parent, ok := mapping[item.ParentID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, item)
			}
		}
	}

	return rootNode
}
