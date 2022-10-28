package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lxygwqf9527/demo-api/apps"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

// 面向接口，真正Service的实现，在服务实例化的时候传递进来
// 也就是(CLI) Start的时候
func NewHostHTTPHandler() *Handler {
	return &Handler{}
}

// 通过写一个实体类，把内部的接口通过HTTP协议暴露出去
// 所以需要以来内部接口的实现
// 该实体类，会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	if apps.HostService == nil {
		panic("dependence host service required")
	}
	h.svc = apps.HostService
}
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
