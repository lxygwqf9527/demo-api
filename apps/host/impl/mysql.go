package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

//
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImp() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子Logger
		// 1.Logger全局实例
		// 2.Logger Level的动态调整，Logrus不支持Level共同调整
		// 3.加入日志轮转
		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
