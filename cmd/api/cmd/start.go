package cmd

import (
	"context"
	"fmt"
	"github.com/cnpythongo/goal/api"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/model/migrate"
	"github.com/cnpythongo/goal/model/redis"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/status"
	"github.com/cnpythongo/goal/pkg/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/judwhite/go-svc"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"net/http"
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
	startCmd.MarkFlagRequired("config")
}

func (app *Application) GetGinEngine() *gin.Engine {
	return app.ginEngine
}

func (app *Application) Init(_ svc.Environment) error {
	cfg, err := config.Load(CfgFile)
	if err != nil {
		return err
	}
	logger := log.Init(&cfg.Logger, "api")
	logger.Info(cfg)

	if err := model.Init(&cfg.Mysql); err != nil {
		logger.Error("Init Mysql Err:", err.Error())
		return err
	}
	if config.GetConfig().Redis.Enable {
		if err = redis.Init(&cfg.Redis); err != nil {
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

	app.ginEngine = api.InitAPIRouters(&cfg)

	return nil
}

func (app *Application) Start() error {
	cfg := config.GetConfig().Http
	app.wrapper.Wrap(func() {
		app.httpServer = &http.Server{
			Handler:        app.ginEngine,
			Addr:           cfg.ListenAddr,
			ReadTimeout:    cfg.ReadTimeout * time.Second,
			WriteTimeout:   cfg.WriteTimeout * time.Second,
			IdleTimeout:    cfg.IdleTimeout * time.Second,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		}
		if err := app.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	})
	log.GetLogger().Info("Api Server Started, Listen on ", cfg.ListenAddr)
	return nil
}

func (app *Application) Stop() error {
	if app.httpServer != nil {
		if err := app.httpServer.Shutdown(context.Background()); err != nil {
			fmt.Printf("Api Server shutdown error:%v\n", err)
		}
		fmt.Println("Api Server shutdown")
	}
	app.wrapper.Wait()
	status.Shutdown()
	status.WaitGroup()

	_ = model.Close()
	if config.GetConfig().Redis.Enable {
		_ = redis.Close()
	}
	fmt.Println("Shutdown end")
	return nil
}
