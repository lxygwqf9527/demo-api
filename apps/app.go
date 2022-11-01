package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

var (
	// 使用Interface{} + 断言进行抽象
	HostService host.Service

	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// Get Impl实例：implApps
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

func RegistryGin(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc

}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// 已经加载完成的Gin App是那些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return
}

func InitGin(r gin.IRouter) {
	// 先初始化所有对象
	for _, v := range ginApps {
		v.Config()
	}
	// 再完成注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

func init() {

}

type ImplService interface {
	Config()      // service初始化
	Name() string // 返回service的名字
}

// 注册由gin编写的handler
// 比如 编写了Http服务A，只需要实现Registry方法，就能把Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config() // service初始化
	Name() string
}
