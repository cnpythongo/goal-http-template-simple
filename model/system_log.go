package model

type SystemLog struct {
	BaseModel
	UserUUID     string `json:"user_uuid" gorm:"column:user_uuid;type:varchar(128);not null;default:'';comment:'用户ID'"`
	Cellphone    string `json:"cellphone" gorm:"column:cellphone;type:varchar(32);not null;default:'';comment:'手机号'"`
	ClientIP     string `json:"client_ip" gorm:"column:client_ip;type:varchar(128);not null;default:'';comment:'客户端IP'"`
	Path         string `json:"path" gorm:"column:path;type:varchar(1024);not null;default:'';comment:'请求路径'"`
	Body         string `json:"body" gorm:"column:body;type:text;not null;default:'';comment:'请求Body'"`
	Method       string `json:"method" gorm:"column:body;type:varchar(64);not null;default:'';comment:'请求方法'"`
	Status       int    `json:"status" gorm:"column:status;type:int(11);not null;default:0;comment:'状态码'"`
	UserAgent    string `json:"user_agent" gorm:"column:user_agent;type:varchar(1024);not null;default:'';comment:'请求UA'"`
	Referer      string `json:"referer" gorm:"column:referer;type:varchar(1024);not null;default:'';comment:'请求携带的Referer'"`
	StartTime    int64  `json:"start_time" gorm:"column:start_time;type:int(11);not null;default:0;comment:'请求开始时间'"`
	EndTime      int64  `json:"end_time" gorm:"column:end_time;type:int(11);not null;default:0;comment:'请求结束时间'"`
	LatencyTime  string `json:"latency_time" gorm:"column:latency_time;type:varchar(128);not null;default:'';comment:'请求耗时'"`
	MemberName   string `json:"member_name" gorm:"column:member_name;type:varchar(64);not null;default:'';comment:'用户名'"`
	OperateTitle string `json:"operate_title" gorm:"column:operate_title;type:varchar(64);not null;default:'';comment:'操作标题'"`
	Page         string `json:"page" gorm:"column:page;type:varchar(64);not null;default:'';comment:'请求触发页面'"`
}

func (s *SystemLog) TableName() string {
	return "system_logs"
}
