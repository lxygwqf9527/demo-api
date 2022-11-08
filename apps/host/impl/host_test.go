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

func TestQuery(t *testing.T) {
	should := assert.New(t)
	req := host.NewQueryHostRequest()
	req.Keywords = "接口测试"
	set, err := service.QueryHost(context.Background(), req)
	if should.NoError(err) {
		for i := range set.Items {
			fmt.Println(set.Items[i].Id)
		}
	}
}

func TestDescribe(t *testing.T) {
	should := assert.New(t)
	req := host.NewDescribeHostRequestWithId("ins-09")
	ins, err := service.DescribeHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func TestUpdate(t *testing.T) {
	should := assert.New(t)
	req := host.NewPutUpdateHostRequest("ins-09")
	req.Name = "更新测试01"
	req.Region = "rg 02"
	req.Type = "small"
	req.CPU = 1
	req.Memory = 2048
	req.Description = "嘻嘻"
	ins, err := service.UpdateHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func TestPatch(t *testing.T) {
	should := assert.New(t)
	req := host.NewPatchUpdateHostRequest("ins-09")
	req.Name = "更新测试05"
	ins, err := service.UpdateHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
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
