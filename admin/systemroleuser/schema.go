package systemroleuser

import "goal-app/pkg/render"

// ReqSystemRoleUserList 角色用户关联列表请求参数
type ReqSystemRoleUserList struct {
	render.Pagination
	OrgId  int64 `form:"org_id"`  // 组织机构ID
	RoleId int64 `form:"role_id"` // 角色ID
	UserId int64 `form:"user_id"` // 用户ID
}

// ReqSystemRoleUserTree 角色用户关联树结构请求参数
type ReqSystemRoleUserTree struct {
	OrgId  int64 `form:"org_id"`  // 组织机构ID
	RoleId int64 `form:"role_id"` // 角色ID
	UserId int64 `form:"user_id"` // 用户ID
}

// ReqSystemRoleUserDetail 角色用户关联详情请求参数
type ReqSystemRoleUserDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemRoleUserCreate 角色用户关联创建请求参数
type ReqSystemRoleUserCreate struct {
	OrgId   int64   `json:"org_id" form:"org_id"`     // 组织机构ID
	RoleId  int64   `json:"role_id" form:"role_id"`   // 角色ID
	UserIds []int64 `json:"user_ids" form:"user_ids"` // 用户ID
}

// ReqSystemRoleUserUpdate 角色用户关联更新请求参数
type ReqSystemRoleUserUpdate struct {
	ID     int64 `json:"id" form:"id"`           // 流水ID
	OrgId  int64 `json:"org_id" form:"org_id"`   // 组织机构ID
	RoleId int64 `json:"role_id" form:"role_id"` // 角色ID
	UserId int64 `json:"user_id" form:"user_id"` // 用户ID
}

// ReqSystemRoleUserDelete 角色用户关联删除请求参数
type ReqSystemRoleUserDelete struct {
	IDs []int64 `json:"ids" form:"ids"` // 流水ID
}

// RespSystemRoleUserItem 角色用户关联单条详情
type RespSystemRoleUserItem struct {
	ID         int64 `json:"id"`          // 流水ID
	CreateTime int64 `json:"create_time"` // 数据创建时间
	UpdateTime int64 `json:"update_time"` // 数据更新时间
	OrgId      int64 `json:"org_id"`      // 组织机构ID
	RoleId     int64 `json:"role_id"`     // 角色ID
	UserId     int64 `json:"user_id"`     // 用户ID
}
