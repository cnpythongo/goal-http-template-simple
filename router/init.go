package router

import (
	limit "github.com/aviddiviner/gin-limit"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/pkg/liveness"
	"github.com/cnpythongo/goal/router/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitDefaultRouter(cfg *config.Configuration) *gin.Engine {
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// middleware
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.CORSMiddleware())

	if cfg.Http.LimitConnection > 0 {
		r.Use(limit.MaxAllowed(cfg.Http.LimitConnection))
	}

	// 最大运行上传文件大小
	r.MaxMultipartMemory = cfg.Http.MaxMultipartMemory

	// common test api
	apiGroup := r.Group("/api/v1")
	apiGroup.GET("/ping", liveness.Ping)

	return r
}
