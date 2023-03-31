// Code generated by protoc-gen-zoo-http. DO NOT EDIT.
// versions:
// - protoc-gen-zoo-http v0.1.0
// - protoc                (unknown)
// source: http/rest/group/group.proto

package group

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	server "github.com/iobrother/zoo/core/transport/http/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = gin.New
var _ = server.NewServer

type GroupHTTPService interface {
	// Add ...
	Add(context.Context, *AddReq) (*AddRsp, error)
	// Create ...
	Create(context.Context, *CreateReq) (*CreateRsp, error)
}

func RegisterGroupHTTPService(g *gin.RouterGroup, svc GroupHTTPService) {
	r := g.Group("")
	r.POST("/zim/groups", _Group_Create0_HTTP_Handler(svc))
	r.POST("/zim/groups/:group_id/members", _Group_Add0_HTTP_Handler(svc))
}

func _Group_Create0_HTTP_Handler(svc GroupHTTPService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := &server.Context{Context: ctx}
		shouldBind := func(req *CreateReq) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			return nil
		}

		var err error
		var req CreateReq
		var rsp *CreateRsp

		if err = shouldBind(&req); err != nil {
			c.SetError(err)
			return
		}
		rsp, err = svc.Create(c.Request.Context(), &req)
		if err != nil {
			c.SetError(err)
			return
		}
		c.JSON(200, rsp)
	}
}
func _Group_Add0_HTTP_Handler(svc GroupHTTPService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := &server.Context{Context: ctx}
		shouldBind := func(req *AddReq) error {
			if err := c.ShouldBind(req); err != nil {
				return err
			}
			if err := c.ShouldBindUri(req); err != nil {
				return err
			}
			return nil
		}

		var err error
		var req AddReq
		var rsp *AddRsp

		if err = shouldBind(&req); err != nil {
			c.SetError(err)
			return
		}
		rsp, err = svc.Add(c.Request.Context(), &req)
		if err != nil {
			c.SetError(err)
			return
		}
		c.JSON(200, rsp)
	}
}