package systemconfig

import "goal-app/pkg/render"

// ReqSystemConfigList 系统配置项列表请求参数
type ReqSystemConfigList struct {
	render.Pagination
	Scope   string `form:"scope"`   // 作用域,global-全局,admin-管理后台,app-前台应用
	Name    string `form:"name"`    // 配置名称
	Value   string `form:"value"`   // 配置值
	Desc    string `form:"desc"`    // 配置说明
	Enabled int64  `form:"enabled"` // 是否启用,0-否,1-是
}

// ReqSystemConfigTree 系统配置项树结构请求参数
type ReqSystemConfigTree struct {
	Scope   string `form:"scope"`   // 作用域,global-全局,admin-管理后台,app-前台应用
	Name    string `form:"name"`    // 配置名称
	Value   string `form:"value"`   // 配置值
	Desc    string `form:"desc"`    // 配置说明
	Enabled int64  `form:"enabled"` // 是否启用,0-否,1-是
}

// ReqSystemConfigDetail 系统配置项详情请求参数
type ReqSystemConfigDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemConfigCreate 系统配置项创建请求参数
type ReqSystemConfigCreate struct {
	Scope   string `json:"scope" form:"scope"`     // 作用域,global-全局,admin-管理后台,app-前台应用
	Name    string `json:"name" form:"name"`       // 配置名称
	Value   string `json:"value" form:"value"`     // 配置值
	Desc    string `json:"desc" form:"desc"`       // 配置说明
	Enabled int64  `json:"enabled" form:"enabled"` // 是否启用,0-否,1-是
}

// ReqSystemConfigUpdate 系统配置项更新请求参数
type ReqSystemConfigUpdate struct {
	ID      int64  `json:"id" form:"id"`           // 流水ID
	Scope   string `json:"scope" form:"scope"`     // 作用域,global-全局,admin-管理后台,app-前台应用
	Name    string `json:"name" form:"name"`       // 配置名称
	Value   string `json:"value" form:"value"`     // 配置值
	Desc    string `json:"desc" form:"desc"`       // 配置说明
	Enabled int64  `json:"enabled" form:"enabled"` // 是否启用,0-否,1-是
}

// ReqSystemConfigDelete 系统配置项删除请求参数
type ReqSystemConfigDelete struct {
	ID int64 `json:"id" form:"id"` // 流水ID
}

// RespSystemConfigItem 系统配置项单条详情
type RespSystemConfigItem struct {
	ID         int64  `json:"id" structs:"id"`                   // 流水ID
	CreateTime int64  `json:"create_time" structs:"create_time"` // 数据创建时间
	UpdateTime int64  `json:"update_time" structs:"update_time"` // 数据更新时间
	Scope      string `json:"scope" structs:"scope"`             // 作用域,global-全局,admin-管理后台,app-前台应用
	Name       string `json:"name" structs:"name"`               // 配置名称
	Value      string `json:"value" structs:"value"`             // 配置值
	Desc       string `json:"desc" structs:"desc"`               // 配置说明
	Enabled    int64  `json:"enabled" structs:"enabled"`         // 是否启用,0-否,1-是
}
