package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/judwhite/go-svc"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goal-app/admin"
	"goal-app/model"
	"goal-app/model/migrate"
	"goal-app/model/redis"
	"goal-app/pkg/config"
	"goal-app/pkg/log"
	"goal-app/pkg/status"
	"goal-app/pkg/wrapper"
	"net/http"
	"path"
	"runtime"
	"syscall"
	"time"
)

type Application struct {
	wrapper    wrapper.Wrapper
	ginEngine  *gin.Engine
	httpServer *http.Server
	cron       *cron.Cron
}

var CfgFile *string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	server(.exe) start -c config.json
	start the api`,
	Run: func(cmd *cobra.Command, args []string) {
		app := &Application{}
		if err := svc.Run(app, syscall.SIGINT, syscall.SIGTERM); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	CfgFile = startCmd.Flags().StringP("config", "c", "", "api config file (required)")
	_ = startCmd.MarkFlagRequired("config")
}

func (app *Application) GetGinEngine() *gin.Engine {
	return app.ginEngine
}

func (app *Application) Init(env svc.Environment) error {
	cfg, err := config.Load(CfgFile)
	if err != nil {
		return err
	}
	// 设置app的运行根目录
	var rootPath string
	if _, filename, _, ok := runtime.Caller(0); ok {
		rootPath = path.Dir(path.Dir(path.Dir(path.Dir(filename))))
	}
	cfg.App.RootPath = rootPath
	fmt.Println("cfg.App.RootPath:", cfg.App.RootPath)

	logger := log.Init(&cfg.Logger, "admin")
	logger.Info(cfg)

	if err = model.Init(&cfg.Mysql); err != nil {
		logger.Error("Init Mysql Err:", err.Error())
		return err
	}
	//if err = model.InitWrite(&cfg.MysqlWrite); err != nil {
	//	logger.Error("Init MysqlWrite Err:", err.Error())
	//	return err
	//}
	if config.GetConfig().Redis.Enable {
		client := redis.Init()
		if client == nil {
			logger.Error("Init Redis Err:", err.Error())
			return err
		}
	}
	migrate.MigrateTables(&cfg)

	// cron task sample
	//app.cron = cron.New()
	//_, err = app.cron.AddFunc("5 0 * * ?", crontab.StatisticalNFTCollect)
	//if err != nil {
	//	return err
	//}
	//app.cron.Start()

	app.ginEngine = admin.InitAdminRouters(&cfg)

	return nil
}

func (app *Application) Start() error {
	app.wrapper.Wrap(func() {
		cfg := config.GetConfig().Http
		app.httpServer = &http.Server{
			Handler:        app.ginEngine,
			Addr:           cfg.AdminListenAddr,
			ReadTimeout:    cfg.ReadTimeout * time.Second,
			WriteTimeout:   cfg.WriteTimeout * time.Second,
			IdleTimeout:    cfg.IdleTimeout * time.Second,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		}
		log.GetLogger().Info("GoalApp server started. Listen on ", cfg.AdminListenAddr)
		if err := app.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	})
	return nil
}

func (app *Application) Stop() error {
	if app.httpServer != nil {
		if err := app.httpServer.Shutdown(context.Background()); err != nil {
			log.GetLogger().Info("HttpServer shutdown error:%v\n", err)
		}
		log.GetLogger().Info("GoalApp server shutdown")
	}
	app.wrapper.Wait()
	status.Shutdown()
	status.WaitGroup()

	_ = model.Close()
	if config.GetConfig().Redis.Enable && redis.GetRedis() != nil {
		_ = redis.GetRedis().Close()
	}
	log.GetLogger().Info("GoalApp server shutdown done")
	return nil
}
