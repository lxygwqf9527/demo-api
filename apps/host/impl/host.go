package impl

import (
	"context"

	"github.com/lxygwqf9527/demo-api/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {

	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 默认值填充
	ins.InjectDefault()

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
