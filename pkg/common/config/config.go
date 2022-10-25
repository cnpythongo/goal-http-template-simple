package config

import (
	"sync"
	"time"

	"github.com/jinzhu/configor"
)

var (
	Cfg Configuration
	mu  sync.RWMutex
)

type (
	AppConfig struct {
		Language string `json:"language" default:"zh_cn"`
		Secret   string `json:"secret" default:"secret."`
		Debug    bool   `json:"debug"`
	}

	HttpConfig struct {
		ListenAddr         string        `json:"listen_addr"`
		AdminListenAddr    string        `json:"admin_listen_addr"`
		LimitConnection    int           `json:"limit_connection"`
		ReadTimeout        time.Duration `json:"read_timeout"`
		WriteTimeout       time.Duration `json:"write_timeout"`
		IdleTimeout        time.Duration `json:"idle_timeout"`
		MaxHeaderBytes     int           `json:"max_header_bytes"`
		MaxMultipartMemory int64         `json:"max_multipart_memory"`
	}

	LoggerConfig struct {
		Level          string        `json:"level"`
		Formatter      string        `json:"formatter"`
		DisableConsole bool          `json:"disable_console"`
		Write          bool          `json:"write"`
		Path           string        `json:"path"`
		FileName       string        `json:"file_name"`
		MaxAge         time.Duration `json:"max_age"`
		RotationTime   time.Duration `json:"rotation_time"`
		Debug          bool          `json:"debug"`
		ReportCaller   bool          `json:"report_caller"`
	}

	MysqlConfig struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"db_name"`
		DbParams string `json:"db_params"`
	}

	RedisConfig struct {
		Enable    bool   `json:"enable"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		Auth      string `json:"auth"`
		MaxIdle   int    `json:"max_idle"`
		MaxActive int    `json:"max_active"`
		Db        int    `json:"db"`
	}

	Configuration struct {
		App    AppConfig    `json:"app"`
		Http   HttpConfig   `json:"http"`
		Mysql  MysqlConfig  `json:"mysql"`
		Logger LoggerConfig `json:"logger"`
		Redis  RedisConfig  `json:"redis"`
	}
)

func Load(file *string) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()

	err := configor.Load(&Cfg, *file)
	if err != nil {
		return Configuration{}, err
	}
	return Cfg, err
}

func GetConfig() Configuration {
	mu.Lock()
	defer mu.Unlock()
	return Cfg
}
