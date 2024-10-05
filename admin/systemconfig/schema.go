package systemconfig

import "goal-app/pkg/render"

type (
	ReqSystemConfigList struct {
		render.Pagination
		Scope   string `json:"scope"`   // 作用域,global-全局,admin-管理后台,app-前台应用
		Name    string `json:"name"`    // 配置名称
		Value   string `json:"value"`   // 配置值
		Desc    string `json:"desc"`    // 配置说明
		Enabled bool   `json:"enabled"` // 是否启用,false-否,true-是
	}

	ReqCreateSystemConfig struct {
		Scope   string `json:"scope" binding:"required" example:"global"` // 作用域,global-全局,admin-管理后台,app-前台应用
		Name    string `json:"name" binding:"required" example:"test"`    // 配置名称
		Value   string `json:"value" binding:"required" example:"test"`   // 配置值
		Desc    string `json:"desc" example:"test"`                       // 配置说明
		Enabled bool   `json:"enabled" binding:"required" example:"true"` // 是否启用,false-否,true-是
	}

	ReqUpdateSystemConfig struct {
		ID      uint64 `json:"id" binding:"required" example:"1"`         // ID
		Scope   string `json:"scope" binding:"required" example:"global"` // 作用域,global-全局,admin-管理后台,app-前台应用
		Name    string `json:"name" binding:"required" example:"test"`    // 配置名称
		Value   string `json:"value" binding:"required" example:"test"`   // 配置值
		Desc    string `json:"desc" example:"test"`                       // 配置说明
		Enabled bool   `json:"enabled" binding:"required" example:"true"` // 是否启用,false-否,true-是
	}

	ReqDeleteSystemConfig struct {
		IDs []uint64 `json:"ids" binding:"required"` // ID列表
	}
)
