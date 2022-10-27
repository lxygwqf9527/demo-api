package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

func NewHostHTTPHandler(svc host.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// 通过写一个实体类，把内部的接口通过HTTP协议暴露出去
// 所以需要以来内部接口的实现
// 该实体类，会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
