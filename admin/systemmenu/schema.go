package systemmenu

import "goal-app/pkg/render"

// ReqSystemMenuList 菜单管理列表请求参数
type ReqSystemMenuList struct {
	render.Pagination
	ParentID  int64  `form:"parent_id"` // '上级菜单ID'
	Kind      string `form:"kind"`      // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name      string `form:"name"`      // '菜单名称'
	Icon      string `form:"icon"`      // '菜单图标'
	Sort      int64  `form:"sort"`      // '菜单排序'
	AuthTag   string `form:"auth_tag"`  // '权限标识'
	Route     string `form:"route"`     // '路由地址'
	Component string `form:"component"` // '前端组件'
	Params    string `form:"params"`    // '路由参数'
	Selected  string `form:"selected"`  // 选中菜单
	Status    string `form:"status"`    // '状态: disable=停用, enable=启用'
}

// ReqSystemMenuTree 菜单管理树结构请求参数
type ReqSystemMenuTree struct {
	ParentID  int64  `form:"parent_id"` // '上级菜单ID'
	Kind      string `form:"kind"`      // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name      string `form:"name"`      // '菜单名称'
	Icon      string `form:"icon"`      // '菜单图标'
	Sort      int64  `form:"sort"`      // '菜单排序'
	AuthTag   string `form:"auth_tag"`  // '权限标识'
	Route     string `form:"route"`     // '路由地址'
	Component string `form:"component"` // '前端组件'
	Params    string `form:"params"`    // '路由参数'
	Selected  string `form:"selected"`  // 选中菜单
	Status    string `form:"status"`    // '状态: disable=停用, enable=启用'
}

// ReqSystemMenuDetail 菜单管理详情请求参数
type ReqSystemMenuDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemMenuCreate 菜单管理创建请求参数
type ReqSystemMenuCreate struct {
	ParentID  int64  `json:"parent_id" form:"parent_id"` // '上级菜单ID'
	Kind      string `json:"kind" form:"kind"`           // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name      string `json:"name" form:"name"`           // '菜单名称'
	Icon      string `json:"icon" form:"icon"`           // '菜单图标'
	Sort      int64  `json:"sort" form:"sort"`           // '菜单排序'
	AuthTag   string `json:"auth_tag" form:"auth_tag"`   // '权限标识'
	Route     string `json:"route" form:"route"`         // '路由地址'
	Component string `json:"component" form:"component"` // '前端组件'
	Params    string `json:"params" form:"params"`       // '路由参数'
	Selected  string `json:"selected" form:"selected"`   // 选中菜单
	Status    string `json:"status" form:"status"`       // '状态: disable=停用, enable=启用'
}

// ReqSystemMenuUpdate 菜单管理更新请求参数
type ReqSystemMenuUpdate struct {
	ID        int64  `json:"id" form:"id"`               // 流水ID
	ParentID  int64  `json:"parent_id" form:"parent_id"` // '上级菜单ID'
	Kind      string `json:"kind" form:"kind"`           // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name      string `json:"name" form:"name"`           // '菜单名称'
	Icon      string `json:"icon" form:"icon"`           // '菜单图标'
	Sort      int64  `json:"sort" form:"sort"`           // '菜单排序'
	AuthTag   string `json:"auth_tag" form:"auth_tag"`   // '权限标识'
	Route     string `json:"route" form:"route"`         // '路由地址'
	Component string `json:"component" form:"component"` // '前端组件'
	Params    string `json:"params" form:"params"`       // '路由参数'
	Selected  string `json:"selected" form:"selected"`   // 选中菜单
	Status    string `json:"status" form:"status"`       // '状态: disable=停用, enable=启用'
}

// ReqSystemMenuDelete 菜单管理删除请求参数
type ReqSystemMenuDelete struct {
	IDs []int64 `json:"ids" binding:"required"`
}

// RespSystemMenuItem 菜单管理单条详情
type RespSystemMenuItem struct {
	ID         int64  `json:"id" structs:"id"`                   // 流水ID
	CreateTime int64  `json:"create_time" structs:"create_time"` // 数据创建时间
	UpdateTime int64  `json:"update_time" structs:"update_time"` // 数据更新时间
	ParentID   int64  `json:"parent_id" structs:"parent_id"`     // '上级菜单ID'
	Kind       string `json:"kind" structs:"kind"`               // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name       string `json:"name" structs:"name"`               // '菜单名称'
	Icon       string `json:"icon" structs:"icon"`               // '菜单图标'
	Sort       int64  `json:"sort" structs:"sort"`               // '菜单排序'
	AuthTag    string `json:"auth_tag" structs:"auth_tag"`       // '权限标识'
	Route      string `json:"route" structs:"route"`             // '路由地址'
	Component  string `json:"component" structs:"component"`     // '前端组件'
	Params     string `json:"params" structs:"params"`           // '路由参数'
	Selected   string `json:"selected" form:"selected"`          // 选中菜单
	Status     string `json:"status" structs:"status"`           // '状态: disable=停用, enable=启用'
}

// RespSystemMenuTree 菜单管理树结构数据
type RespSystemMenuTree struct {
	ID         int64                 `json:"id" structs:"id"`                   // 流水ID
	CreateTime int64                 `json:"create_time" structs:"create_time"` // 数据创建时间
	UpdateTime int64                 `json:"update_time" structs:"update_time"` // 数据更新时间
	ParentID   int64                 `json:"parent_id" structs:"parent_id"`     // '上级菜单ID'
	Kind       string                `json:"kind" structs:"kind"`               // '权限类型: dir=目录，menu=菜单，button=按钮''
	Name       string                `json:"name" structs:"name"`               // '菜单名称'
	Icon       string                `json:"icon" structs:"icon"`               // '菜单图标'
	Sort       int64                 `json:"sort" structs:"sort"`               // '菜单排序'
	AuthTag    string                `json:"auth_tag" structs:"auth_tag"`       // '权限标识'
	Route      string                `json:"route" structs:"route"`             // '路由地址'
	Component  string                `json:"component" structs:"component"`     // '前端组件'
	Params     string                `json:"params" structs:"params"`           // '路由参数'
	Selected   string                `json:"selected" form:"selected"`          // 选中菜单
	Status     string                `json:"status" structs:"status"`           // '状态: disable=停用, enable=启用'
	ParentName string                `json:"parent_name" structs:"parent_name"` // '父级名称'
	Children   []*RespSystemMenuTree `json:"children"`                          // 子节点
}
