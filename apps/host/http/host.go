package http

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"github.com/lxygwqf9527/demo-api/apps/host"
)

// 用于暴露Host Service接口
func (h *Handler) createHost(c *gin.Context) {
	// 对外的api接口
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)
}

//

func (h *Handler) queryHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewQueryHostFromHttp(c.Request)
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
	}
	response.Success(c.Writer, set)
}

func (h *Handler) describeHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewDescribeHostRequestWithId(c.Param("id"))

	set, err := h.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
	}
	response.Success(c.Writer, set)
}

func (h *Handler) putUpdateHost(c *gin.Context) {
	req := host.NewPutUpdateHostRequest(c.Param("id"))
	// 解析Body里面的数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

func (h *Handler) patchUpdateHost(c *gin.Context) {
	req := host.NewPatchUpdateHostRequest(c.Param("id"))
	// 解析Body里面的数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}
