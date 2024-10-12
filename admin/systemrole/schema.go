package systemrole

import "goal-app/pkg/render"

// ReqSystemRoleList 角色管理列表请求参数
type ReqSystemRoleList struct {
	render.Pagination
	Name   string `form:"name"`   // 角色名称
	Desc   string `form:"desc"`   // 角色描述
	Status *int64 `form:"status"` // 角色状态, 0-禁用, 1-启用
}

// ReqSystemRoleDetail 角色管理详情请求参数
type ReqSystemRoleDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemRoleCreate 角色管理创建请求参数
type ReqSystemRoleCreate struct {
	Name      string `json:"name" form:"name"`             // 角色名称
	Desc      string `json:"desc" form:"desc"`             // 角色描述
	Status    int64  `json:"status" form:"status"`         // 角色状态, 0-禁用, 1-启用
	IsDeleted int64  `json:"is_deleted" form:"is_deleted"` // 是否被删除, 0-否, 1-是
}

// ReqSystemRoleUpdate 角色管理更新请求参数
type ReqSystemRoleUpdate struct {
	ID     int64  `json:"id" form:"id"`         // 流水ID
	Name   string `json:"name" form:"name"`     // 角色名称
	Desc   string `json:"desc" form:"desc"`     // 角色描述
	Status int64  `json:"status" form:"status"` // 角色状态, 0-禁用, 1-启用
}

// ReqSystemRoleDelete 角色管理删除请求参数
type ReqSystemRoleDelete struct {
	IDs []int64 `json:"ids" form:"ids"` // 流水ID
}

// RespSystemRoleItem 角色管理单条详情
type RespSystemRoleItem struct {
	ID         int64  `json:"id"`          // 流水ID
	Name       string `json:"name"`        // 角色名称
	Desc       string `json:"desc"`        // 角色描述
	Status     int64  `json:"status"`      // 角色状态, 0-禁用, 1-启用
	CreateTime int64  `json:"create_time"` // 数据创建时间
	UpdateTime int64  `json:"update_time"` // 数据更新时间
}
