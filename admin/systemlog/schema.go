package systemlog

import "goal-app/pkg/render"

// ReqSystemLogList 系统日志列表请求参数
type ReqSystemLogList struct {
	render.Pagination
	Cellphone  string `form:"cellphone"`   // '手机号'
	MemberName string `form:"member_name"` // '用户名'
}

// ReqSystemLogTree 系统日志树结构请求参数
type ReqSystemLogTree struct {
	UserUuid     string `form:"user_uuid"`     // '用户ID'
	Cellphone    string `form:"cellphone"`     // '手机号'
	ClientIp     string `form:"client_ip"`     // '客户端IP'
	Path         string `form:"path"`          // '请求路径'
	Body         string `form:"body"`          // '请求Body'
	Status       int64  `form:"status"`        // '状态码'
	UserAgent    string `form:"user_agent"`    // '请求UA'
	Referer      string `form:"referer"`       // '请求携带的Referer'
	StartTime    int64  `form:"start_time"`    // '请求开始时间'
	EndTime      int64  `form:"end_time"`      // '请求结束时间'
	LatencyTime  string `form:"latency_time"`  // '请求耗时'
	MemberName   string `form:"member_name"`   // '用户名'
	OperateTitle string `form:"operate_title"` // '操作标题'
	Page         string `form:"page"`          // '请求触发页面'
}

// ReqSystemLogDetail 系统日志详情请求参数
type ReqSystemLogDetail struct {
	ID int64 `form:"id"` // 流水ID
}

// ReqSystemLogCreate 系统日志创建请求参数
type ReqSystemLogCreate struct {
	UserUuid     string `json:"user_uuid" form:"user_uuid"`         // '用户ID'
	Cellphone    string `json:"cellphone" form:"cellphone"`         // '手机号'
	ClientIp     string `json:"client_ip" form:"client_ip"`         // '客户端IP'
	Path         string `json:"path" form:"path"`                   // '请求路径'
	Body         string `json:"body" form:"body"`                   // '请求Body'
	Status       int64  `json:"status" form:"status"`               // '状态码'
	UserAgent    string `json:"user_agent" form:"user_agent"`       // '请求UA'
	Referer      string `json:"referer" form:"referer"`             // '请求携带的Referer'
	StartTime    int64  `json:"start_time" form:"start_time"`       // '请求开始时间'
	EndTime      int64  `json:"end_time" form:"end_time"`           // '请求结束时间'
	LatencyTime  string `json:"latency_time" form:"latency_time"`   // '请求耗时'
	MemberName   string `json:"member_name" form:"member_name"`     // '用户名'
	OperateTitle string `json:"operate_title" form:"operate_title"` // '操作标题'
	Page         string `json:"page" form:"page"`                   // '请求触发页面'
}

// ReqSystemLogUpdate 系统日志更新请求参数
type ReqSystemLogUpdate struct {
	ID           int64  `json:"id" form:"id"`                       // 流水ID
	UserUuid     string `json:"user_uuid" form:"user_uuid"`         // '用户ID'
	Cellphone    string `json:"cellphone" form:"cellphone"`         // '手机号'
	ClientIp     string `json:"client_ip" form:"client_ip"`         // '客户端IP'
	Path         string `json:"path" form:"path"`                   // '请求路径'
	Body         string `json:"body" form:"body"`                   // '请求Body'
	Status       int64  `json:"status" form:"status"`               // '状态码'
	UserAgent    string `json:"user_agent" form:"user_agent"`       // '请求UA'
	Referer      string `json:"referer" form:"referer"`             // '请求携带的Referer'
	StartTime    int64  `json:"start_time" form:"start_time"`       // '请求开始时间'
	EndTime      int64  `json:"end_time" form:"end_time"`           // '请求结束时间'
	LatencyTime  string `json:"latency_time" form:"latency_time"`   // '请求耗时'
	MemberName   string `json:"member_name" form:"member_name"`     // '用户名'
	OperateTitle string `json:"operate_title" form:"operate_title"` // '操作标题'
	Page         string `json:"page" form:"page"`                   // '请求触发页面'
}

// ReqSystemLogDelete 系统日志删除请求参数
type ReqSystemLogDelete struct {
	ID int64 `json:"id" form:"id"` // 流水ID
}

// RespSystemLogItem 系统日志单条详情
type RespSystemLogItem struct {
	ID           int64  `json:"id" structs:"id"`                       // 流水ID
	CreateTime   int64  `json:"create_time" structs:"create_time"`     // 数据创建时间
	UpdateTime   int64  `json:"update_time" structs:"update_time"`     // 数据更新时间
	UserUuid     string `json:"user_uuid" structs:"user_uuid"`         // '用户ID'
	Cellphone    string `json:"cellphone" structs:"cellphone"`         // '手机号'
	ClientIp     string `json:"client_ip" structs:"client_ip"`         // '客户端IP'
	Path         string `json:"path" structs:"path"`                   // '请求路径'
	Body         string `json:"body" structs:"body"`                   // '请求Body'
	Status       int64  `json:"status" structs:"status"`               // '状态码'
	UserAgent    string `json:"user_agent" structs:"user_agent"`       // '请求UA'
	Referer      string `json:"referer" structs:"referer"`             // '请求携带的Referer'
	StartTime    int64  `json:"start_time" structs:"start_time"`       // '请求开始时间'
	EndTime      int64  `json:"end_time" structs:"end_time"`           // '请求结束时间'
	LatencyTime  string `json:"latency_time" structs:"latency_time"`   // '请求耗时'
	MemberName   string `json:"member_name" structs:"member_name"`     // '用户名'
	OperateTitle string `json:"operate_title" structs:"operate_title"` // '操作标题'
	Page         string `json:"page" structs:"page"`                   // '请求触发页面'
}
