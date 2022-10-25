package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/cnpythongo/goal/pkg/common/config"
)

var (
	logger             *logrus.Logger
	mu                 sync.Mutex
	defaultLogFileName = "daily.log"
	defaultLevel       = "debug"
)

func GetLogger() *logrus.Logger {
	if logger == nil {
		panic("logger not inited")
	}
	return logger
}

func defaultConfig() *config.LoggerConfig {
	return &config.LoggerConfig{
		Level:          defaultLevel,
		Formatter:      "text",
		DisableConsole: false,
		Write:          false,
		Path:           os.TempDir(),
		FileName:       defaultLogFileName,
		MaxAge:         time.Duration(24*7) * time.Hour,
		RotationTime:   time.Duration(24) * time.Hour,
		Debug:          false,
	}
}

func Init(conf *config.LoggerConfig, app string) *logrus.Logger {
	mu.Lock()
	defer mu.Unlock()

	if conf == nil {
		conf = defaultConfig()
	}

	if logger != nil {
		return logger
	}

	_log := logrus.New()

	// get logLevel
	level := conf.Level
	if level == "" {
		level = defaultLevel
	}
	logLevel := GetLogLevel(level)

	logDir := filepath.Join(conf.Path, app)
	if logDir == "" {
		logDir = os.TempDir()
	}

	logFileName := conf.FileName
	if logFileName == "" {
		logFileName = defaultLogFileName
	}

	printLog := !conf.DisableConsole

	maxAge := conf.MaxAge

	rotationTime := conf.RotationTime

	_log.SetLevel(logLevel)
	if conf.ReportCaller {
		_log.SetReportCaller(true)
	}
	if conf.Write {
		storeLogDir := logDir

		err := os.MkdirAll(storeLogDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("creating log file failed: %s", err.Error()))
		}

		path := filepath.Join(storeLogDir, logFileName)
		writer, err := rotatelogs.New(
			path+"-%Y-%m-%d.log",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
			rotatelogs.WithClock(rotatelogs.Local),
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs log failed: %s", err.Error()))
		}

		var formatter logrus.Formatter

		formatter = &logrus.TextFormatter{}
		if conf.Formatter == "json" {
			formatter = &logrus.JSONFormatter{}
		}
		if conf.Debug {
			_log.AddHook(lfshook.NewHook(
				lfshook.WriterMap{
					logrus.DebugLevel: writer,
					logrus.InfoLevel:  writer,
					logrus.WarnLevel:  writer,
					logrus.ErrorLevel: writer,
					logrus.FatalLevel: writer,
				},
				formatter,
			))

			defaultLogFilePrefix := logFileName + "."
			pathMap := lfshook.PathMap{
				logrus.DebugLevel: fmt.Sprintf("%s/%sdebug", storeLogDir, defaultLogFilePrefix),
				logrus.InfoLevel:  fmt.Sprintf("%s/%sinfo", storeLogDir, defaultLogFilePrefix),
				logrus.WarnLevel:  fmt.Sprintf("%s/%swarn", storeLogDir, defaultLogFilePrefix),
				logrus.ErrorLevel: fmt.Sprintf("%s/%serror", storeLogDir, defaultLogFilePrefix),
				logrus.FatalLevel: fmt.Sprintf("%s/%sfatal", storeLogDir, defaultLogFilePrefix),
			}
			_log.AddHook(lfshook.NewHook(
				pathMap,
				formatter,
			))
		} else {
			_log.Out = writer
			_log.Formatter = formatter
		}

	} else {
		if printLog {
			_log.Out = os.Stdout
		}
	}
	logger = _log
	return logger
}
