package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/lxygwqf9527/demo-api/apps/host"
	"github.com/lxygwqf9527/demo-api/apps/host/impl"
)

var (
	service host.Service
)

func TestCreat(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	zap.DevelopmentSetup()
	service = impl.NewHostServiceImp()
}
