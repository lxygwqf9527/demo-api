package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lxygwqf9527/demo-api/apps"
	"github.com/lxygwqf9527/demo-api/protocol"

	_ "github.com/lxygwqf9527/demo-api/apps/all"
	"github.com/lxygwqf9527/demo-api/conf"
	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
	confETCD string
)

func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	// Config里面的日志配置，来配置全局logger对象
	lc := conf.C().Log

	// 解析日志Level配置
	// DebugLevel: "debug",
	// InfoLevel:  "info",
	// WarnLevel:  "warning",
	// ErrorLevel: "error",
	// FatalLevel: "fatal",
	// PanicLevel: "panic",

	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	zapConfig := zap.DefaultConfig()        // 使用默认配置初始化Logger的全局配置
	zapConfig.Level = level                 // 配置日志的Level级别
	zapConfig.Files.RotateOnStartup = false //程序每启动一次，不必都生成一个新的日志文件
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没在把日志输入输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "demo-api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	// 配置日志的输出格式
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

// 1.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}

		// 初始化全局日志Logger
		if err := loadGlobalLogger(); err != nil {
			return err
		}
		// 加载Host Service的实体类
		apps.InitImpl()

		svc := newManager()
		ch := make(chan os.Signal, 1)
		// channel是一种复合数据结构，可以当成一个容器
		defer close(ch)

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		go svc.WaitStop(ch)
		return svc.Start()
	},
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}

// 用于管理所有需要启动的服务
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("CLI"),
	}
}

func (m *manager) Start() error {
	return m.http.Start()
}

// 处理来自外部的中断信号，比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {

		default:
			m.l.Infof("received signal %s", v)
			m.http.Stop()
		}
	}
}

// 处理来自外部的中断信号，比如Terminal
func (m *manager) Stop() {

}
