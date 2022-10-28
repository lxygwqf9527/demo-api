package apps

import (
	"fmt"

	"github.com/lxygwqf9527/demo-api/apps/host"
)

var (
	HostService host.Service

	svcs = map[string]Service{}
)

func Registry(svc Service) {
	if _, ok := svcs[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	svcs[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

//用户初始化 注册到Ioc容器里面的所有服务
func Init() {
	for _, v := range svcs {
		v.Config()
	}
}

type Service interface {
	Config()      // service初始化
	Name() string // 返回service的名字
}
