package model

type Config struct {
	BaseModel
	Scope   SystemConfigScope `json:"scope" gorm:"column:scope;type:varchar(50);not null;default:'';comment:作用域,global-全局,admin-管理后台,app-前台应用"`
	Name    string            `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';comment:配置名称"`
	Value   string            `json:"value" gorm:"column:value;type:text;not null;comment:配置值"`
	Desc    string            `json:"desc" gorm:"column:desc;type:varchar(200);not null;default:'';comment:配置说明"`
	Enabled bool              `json:"enabled" gorm:"column:enabled;type:tinyint(1);not null;default:1;comment:是否启用,0-否,1-是"`
}

func (c *Config) TableName() string {
	return "system_config"
}

func NewConfig() *Config {
	return &Config{}
}
