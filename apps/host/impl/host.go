package impl

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("create host")
	i.l.Info("create host %s", ins.Name)
	i.l.Warn("create host %s", ins.Name)
	i.l.Named("Create").Error("create host %s", ins.Name)
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")
	return nil, nil
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
