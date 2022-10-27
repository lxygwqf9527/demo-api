package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/lxygwqf9527/demo-api/apps/host/http"
	"github.com/lxygwqf9527/demo-api/apps/host/impl"
	"github.com/lxygwqf9527/demo-api/conf"
	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
	confETCD string
)

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
		// 加载Host Service的实体类
		service := impl.NewHostServiceImpl()
		// 通过Host API Handler对外提供HTTP RestFul接口
		api := http.NewHostHTTPHandler(service)

		g := gin.Default()
		api.Registry(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
