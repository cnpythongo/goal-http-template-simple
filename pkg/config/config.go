package config

import (
	"fmt"
	"path"
	"runtime"
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
		RunMode  string `json:"run_mode" default:"local"` // 运行模式，也可以理解为运行环境：local-本地开发, test-测试, prod-生产
		RootPath string `json:"root_path"`                // 运行目录
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

	StorageConfig struct {
		Driver          string `json:"driver"`            // local(本地存储) or cos(腾讯云COS)
		UploadDirectory string `json:"upload_directory"`  // 文件上传根目录
		UploadImageSize int64  `json:"upload_image_size"` // 1024 * 1024 * 10  = 10m
		PublicPrefix    string `json:"public_prefix"`     // 文件访问路径前缀
	}

	// StorageCOS 腾讯云COS
	StorageCOS struct {
		SecretID  string `json:"secret_id"`
		SecretKey string `json:"secret_key"`
		Bucket    string `json:"bucket"`
		Region    string `json:"region"`
	}

	Configuration struct {
		App        AppConfig     `json:"app"`
		Http       HttpConfig    `json:"http"`
		Mysql      MysqlConfig   `json:"mysql"`
		MysqlWrite MysqlConfig   `json:"mysql_write"`
		Logger     LoggerConfig  `json:"logger"`
		Redis      RedisConfig   `json:"redis"`
		Storage    StorageConfig `json:"storage"`
		StorageCOS StorageCOS    `json:"storage_cos"` // 腾讯云COS
	}
)

func Load(file *string) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()

	err := configor.Load(&Cfg, *file)
	if err != nil {
		return Configuration{}, err
	}
	// 设置app的运行根目录
	var rootPath string
	if _, filename, _, ok := runtime.Caller(0); ok {
		rootPath = path.Dir(path.Dir(path.Dir(filename)))
	}
	Cfg.App.RootPath = rootPath
	fmt.Println("Cfg.App.RootPath:", Cfg.App.RootPath)
	return Cfg, err
}

func GetConfig() Configuration {
	mu.Lock()
	defer mu.Unlock()
	return Cfg
}
