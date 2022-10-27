package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/lxygwqf9527/demo-api/apps/host"
	"github.com/lxygwqf9527/demo-api/apps/host/impl"
	"github.com/lxygwqf9527/demo-api/conf"
	"github.com/stretchr/testify/assert"
)

var (
	service host.Service
)

func TestCreat(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048

	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}

}

func init() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	service = impl.NewHostServiceImpl()
}
