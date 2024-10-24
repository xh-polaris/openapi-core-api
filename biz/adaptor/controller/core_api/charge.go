// Code generated by hertz generator.

package core_api

import (
	"context"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/openapi-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	core_api "github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	constsx "github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
)

// CreateBaseInterface .
// @router /base/create [POST]
func CreateBaseInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateBaseInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.CreateBaseInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateBaseInterface .
// @router /base/update [POST]
func UpdateBaseInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateBaseInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.UpdateBaseInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteBaseInterface .
// @router /base/delete [GET]
func DeleteBaseInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteBaseInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.DeleteBaseInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetBaseInterfaces .
// @router /base/get [GET]
func GetBaseInterfaces(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetBaseInterfacesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.GetBaseInterfaces(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CreateFullInterface .
// @router /full/create [POST]
func CreateFullInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateFullInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.CreateFullInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateFullInterface .
// @router /full/update [POST]
func UpdateFullInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateFullInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.UpdateFullInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateMargin .
// @router /full/margin [POST]
func UpdateMargin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateMarginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.UpdateMargin(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteFullInterface .
// @router /full/delete [GET]
func DeleteFullInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteFullInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.DeleteFullInterface(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetFullInterfaces .
// @router /full/get [GET]
func GetFullInterfaces(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetFullInterfacesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.GetFullInterfaces(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CreateGradient .
// @router /gradient/create [POST]
func CreateGradient(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateGradientReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.CreateGradient(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateGradient .
// @router /gradient/update [POST]
func UpdateGradient(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateGradientReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 权限校验
	if util.CheckRole(ctx, constsx.Admin) {
		adaptor.PostProcess(ctx, c, &req, nil, constsx.ErrRole)
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.UpdateGradient(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetGradient .
// @router /gradient/get [GET]
func GetGradient(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetGradientReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.ChargeService.GetGradient(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}
