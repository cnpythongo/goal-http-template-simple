package systemrolemenu

import "goal-app/pkg/render"

// ReqSystemRoleMenuList 角色菜单关联列表请求参数
type ReqSystemRoleMenuList struct {
	render.Pagination
	OrgId  int64 `form:"org_id"`  // 组织机构ID
	RoleId int64 `form:"role_id"` // 角色ID
	MenuId int64 `form:"menu_id"` // 菜单ID
}

// ReqSystemRoleMenuTree 角色菜单关联树结构请求参数
type ReqSystemRoleMenuTree struct {
	OrgId  int64 `form:"org_id"`  // 组织机构ID
	RoleId int64 `form:"role_id"` // 角色ID
	MenuId int64 `form:"menu_id"` // 菜单ID
}

// ReqSystemRoleMenuDetail 角色菜单关联详情请求参数
type ReqSystemRoleMenuDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemRoleMenuCreate 角色菜单关联创建请求参数
type ReqSystemRoleMenuCreate struct {
	OrgId  int64 `json:"org_id" form:"org_id"`   // 组织机构ID
	RoleId int64 `json:"role_id" form:"role_id"` // 角色ID
	MenuId int64 `json:"menu_id" form:"menu_id"` // 菜单ID
}

// ReqSystemRoleMenuUpdate 角色菜单关联更新请求参数
type ReqSystemRoleMenuUpdate struct {
	ID     int64 `json:"id" form:"id"`           // 流水ID
	OrgId  int64 `json:"org_id" form:"org_id"`   // 组织机构ID
	RoleId int64 `json:"role_id" form:"role_id"` // 角色ID
	MenuId int64 `json:"menu_id" form:"menu_id"` // 菜单ID
}

// ReqSystemRoleMenuDelete 角色菜单关联删除请求参数
type ReqSystemRoleMenuDelete struct {
	IDs []int64 `json:"ids" form:"ids"` // 流水ID
}

// RespSystemRoleMenuItem 角色菜单关联单条详情
type RespSystemRoleMenuItem struct {
	Id         int64 `json:"id" structs:"id"`                   // 流水ID
	CreateTime int64 `json:"create_time" structs:"create_time"` // 数据创建时间
	UpdateTime int64 `json:"update_time" structs:"update_time"` // 数据更新时间
	OrgId      int64 `json:"org_id" structs:"org_id"`           // 组织机构ID
	RoleId     int64 `json:"role_id" structs:"role_id"`         // 角色ID
	MenuId     int64 `json:"menu_id" structs:"menu_id"`         // 菜单ID
}
