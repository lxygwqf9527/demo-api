package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lxygwqf9527/demo-api/apps"
	"github.com/lxygwqf9527/demo-api/conf"
)

// HttpService构造函数
func NewHttpService() *HttpService {
	// new Gin Router实例， 并没有加载路由
	r := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		l:      zap.L().Named("HTTP Service"),
		r:      r,
	}
}

type HttpService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

func (s *HttpService) Start() error {
	// 加载Handler
	apps.InitGin(s.r)
	apps := apps.LoadedGinApps()
	s.l.Infof("loaded gin apps :%v", apps)

	// 该操作是阻塞的，监听端口，等待请求
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped success")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

func (s *HttpService) Stop() {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shut down http service error, %s", err)
	}
}
