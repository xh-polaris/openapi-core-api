package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/charge"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_charge"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	gencharge "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type IChargeService interface {
	CreateBaseInterface(ctx context.Context, req *core_api.CreateBaseInterfaceReq) (*core_api.Response, error)
	UpdateBaseInterface(ctx context.Context, req *core_api.UpdateBaseInterfaceReq) (*core_api.Response, error)
	DeleteBaseInterface(ctx context.Context, req *core_api.DeleteBaseInterfaceReq) (*core_api.Response, error)
	GetBaseInterfaces(ctx context.Context, req *core_api.GetBaseInterfacesReq) (*core_api.GetBaseInterfacesResp, error)
	CreateFullInterface(ctx context.Context, req *core_api.CreateFullInterfaceReq) (*core_api.Response, error)
	UpdateFullInterface(ctx context.Context, req *core_api.UpdateFullInterfaceReq) (*core_api.Response, error)
	UpdateMargin(ctx context.Context, req *core_api.UpdateMarginReq) (*core_api.Response, error)
	DeleteFullInterface(ctx context.Context, req *core_api.DeleteFullInterfaceReq) (*core_api.Response, error)
	GetFullInterfaces(ctx context.Context, req *core_api.GetFullInterfacesReq) (*core_api.GetFullInterfacesResp, error)
	CreateGradient(ctx context.Context, req *core_api.CreateGradientReq) (*core_api.Response, error)
	UpdateGradient(ctx context.Context, req *core_api.UpdateGradientReq) (*core_api.Response, error)
	GetGradient(ctx context.Context, req *core_api.GetGradientReq) (*core_api.GetGradientResp, error)
}

type ChargeService struct {
	ChargeClient openapi_charge.IOpenapiCharge
}

var ChargeServiceSet = wire.NewSet(
	wire.Struct(new(ChargeService), "*"),
	wire.Bind(new(IChargeService), new(*ChargeService)),
)

func (s *ChargeService) CreateBaseInterface(ctx context.Context, req *core_api.CreateBaseInterfaceReq) (*core_api.Response, error) {
	// rpc创建基础接口
	resp, err := s.ChargeClient.CreateBaseInterface(ctx, &gencharge.CreateBaseInterfaceReq{
		Name:    req.Name,
		Host:    req.Host,
		Path:    req.Path,
		Method:  gencharge.MethodType(req.Method),
		PassWay: gencharge.PassWayType(req.PassWay),
		Params:  changeParamsToGen(req.Params),
		Content: req.Content,
	})
	// 创建失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "创建基础接口失败"), err
	}
	// 创建成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) UpdateBaseInterface(ctx context.Context, req *core_api.UpdateBaseInterfaceReq) (*core_api.Response, error) {
	// rpc更新
	resp, err := s.ChargeClient.UpdateBaseInterface(ctx, &gencharge.UpdateBaseInterfaceReq{
		Id:      req.Id,
		Name:    req.Name,
		Host:    req.Host,
		Path:    req.Path,
		Method:  gencharge.MethodType(req.Method),
		PassWay: gencharge.PassWayType(req.PassWay),
		Params:  changeParamsToGen(req.Params),
		Content: req.Content,
		Status:  gencharge.InterfaceStatus(req.Status),
	})
	// 更新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}
	// 更新成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) DeleteBaseInterface(ctx context.Context, req *core_api.DeleteBaseInterfaceReq) (*core_api.Response, error) {
	// rpc删除
	resp, err := s.ChargeClient.DeleteBaseInterface(ctx, &gencharge.DeleteBaseInterfaceReq{
		Id: req.Id,
	})
	// 删除失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "删除失败"), err
	}
	// 删除成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) GetBaseInterfaces(ctx context.Context, req *core_api.GetBaseInterfacesReq) (*core_api.GetBaseInterfacesResp, error) {
	if req.PaginationOptions == nil {
		req.PaginationOptions = &basic.PaginationOptions{}
	}
	// rpc获取
	resp, err := s.ChargeClient.GetBaseInterface(ctx, &gencharge.GetBaseInterfaceReq{
		PaginationOptions: &genbasic.PaginationOptions{
			Page:      req.PaginationOptions.Page,
			Limit:     req.PaginationOptions.Limit,
			LastToken: req.PaginationOptions.LastToken,
			Backward:  req.PaginationOptions.Backward,
			Offset:    req.PaginationOptions.Offset,
		},
	})
	if err != nil {
		return nil, err
	}
	// 类型转换
	var infs = make([]*charge.BaseInterface, 0)
	err = copier.Copy(&infs, resp.BaseInterfaces)
	if err != nil {
		return nil, err
	}
	// 处理无法复制的类型
	for i := 0; i < len(infs); i++ {
		infs[i].Method = charge.MethodType(resp.BaseInterfaces[i].Method)
		infs[i].PassWay = charge.PassWayType(resp.BaseInterfaces[i].PassWay)
		infs[i].Params = changeParamsFromGen(resp.BaseInterfaces[i].Params)
		infs[i].Status = charge.InterfaceStatus(resp.BaseInterfaces[i].Status)
	}
	// 成功查询
	return &core_api.GetBaseInterfacesResp{
		BaseInterfaces: infs,
		Total:          resp.Total,
	}, nil
}

func (s *ChargeService) CreateFullInterface(ctx context.Context, req *core_api.CreateFullInterfaceReq) (*core_api.Response, error) {
	// rpc创建
	resp, err := s.ChargeClient.CreateFullInterface(ctx, &gencharge.CreateFullInterfaceReq{
		BaseInterfaceId: req.BaseInterfaceId,
		UserId:          req.UserId,
		ChargeType:      gencharge.ChargeType(req.ChargeType),
		Price:           req.Price,
		Margin:          req.Margin,
	})
	// 创建失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "创建完整请求失败"), err
	}
	// 创建成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) UpdateFullInterface(ctx context.Context, req *core_api.UpdateFullInterfaceReq) (*core_api.Response, error) {
	// rpc更新
	resp, err := s.ChargeClient.UpdateFullInterface(ctx, &gencharge.UpdateFullInterfaceReq{
		Id:         req.Id,
		ChargeType: gencharge.ChargeType(req.ChargeType),
		Price:      req.Price,
		Status:     gencharge.InterfaceStatus(req.Status),
	})
	// 更新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}
	// 更新成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) UpdateMargin(ctx context.Context, req *core_api.UpdateMarginReq) (*core_api.Response, error) {
	// rpc更新
	resp, err := s.ChargeClient.UpdateMargin(ctx, &gencharge.UpdateMarginReq{
		Id:        req.Id,
		Increment: req.Increment,
	})
	// 更新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "余额更新失败"), err
	}
	// 更新成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) DeleteFullInterface(ctx context.Context, req *core_api.DeleteFullInterfaceReq) (*core_api.Response, error) {
	// rpc删除
	resp, err := s.ChargeClient.DeleteFullInterface(ctx, &gencharge.DeleteFullInterfaceReq{
		Id: req.Id,
	})
	// 删除失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "删除失败"), err
	}
	// 删除成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) GetFullInterfaces(ctx context.Context, req *core_api.GetFullInterfacesReq) (*core_api.GetFullInterfacesResp, error) {
	if req.PaginationOptions == nil {
		req.PaginationOptions = &basic.PaginationOptions{}
	}
	// rpc获取
	resp, err := s.ChargeClient.GetFullInterface(ctx, &gencharge.GetFullInterfaceReq{
		UserId: req.UserId,
		PaginationOptions: &genbasic.PaginationOptions{
			Page:      req.PaginationOptions.Page,
			Limit:     req.PaginationOptions.Limit,
			LastToken: req.PaginationOptions.LastToken,
			Backward:  req.PaginationOptions.Backward,
			Offset:    req.PaginationOptions.Offset,
		},
	})
	// 获取失败
	if err != nil {
		return nil, err
	}
	// 类型转换
	var infs = make([]*charge.FullInterface, 0)
	err = copier.Copy(&infs, resp.FullInterfaces)
	if err != nil {
		return nil, err
	}
	// 处理无法复制字段
	for i := 0; i < len(infs); i++ {
		infs[i].ChargeType = charge.ChargeType(resp.FullInterfaces[i].ChargeType)
		infs[i].Status = charge.InterfaceStatus(resp.FullInterfaces[i].Status)
	}

	return &core_api.GetFullInterfacesResp{
		FullInterfaces: infs,
		Total:          resp.Total,
	}, nil
}

func (s *ChargeService) CreateGradient(ctx context.Context, req *core_api.CreateGradientReq) (*core_api.Response, error) {
	// rpc创建
	resp, err := s.ChargeClient.CreateGradient(ctx, &gencharge.CreateGradientReq{
		BaseInterfaceId: req.BaseInterfaceId,
		Discounts:       changeDiscountsToGen(req.Discounts),
	})
	// 创建失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "梯度折扣创建失败"), err
	}
	// 创建成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) UpdateGradient(ctx context.Context, req *core_api.UpdateGradientReq) (*core_api.Response, error) {
	// rpc更新
	resp, err := s.ChargeClient.UpdateGradient(ctx, &gencharge.UpdateGradientReq{
		Id:        req.Id,
		Discounts: changeDiscountsToGen(req.Discounts),
		Status:    gencharge.InterfaceStatus(req.Status),
	})
	// 更新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}
	// 更新成功
	return util.SuccessResponse(resp)
}

func (s *ChargeService) GetGradient(ctx context.Context, req *core_api.GetGradientReq) (*core_api.GetGradientResp, error) {
	// rpc获取
	resp, err := s.ChargeClient.GetGradient(ctx, &gencharge.GetGradientReq{
		BaseInterfaceId: req.BaseInterfaceId,
	})
	if err != nil {
		return nil, err
	}
	var gradient = &charge.Gradient{
		Id:              resp.Gradient.Id,
		BaseInterfaceId: resp.Gradient.BaseInterfaceId,
		Discounts:       changeDiscountsFromGen(resp.Gradient.Discounts),
		Status:          charge.InterfaceStatus(resp.Gradient.Status),
		CreateTime:      resp.Gradient.CreateTime,
		UpdateTime:      resp.Gradient.UpdateTime,
	}
	return &core_api.GetGradientResp{
		Gradient: gradient,
	}, nil
}

// 类型转换用工具
func changeParamsToGen(parameter []*charge.Parameter) []*gencharge.Parameter {
	var params = make([]*gencharge.Parameter, 0)
	err := copier.Copy(&params, parameter)
	if err != nil {
		return nil
	}
	return params
}

func changeParamsFromGen(parameter []*gencharge.Parameter) []*charge.Parameter {
	var params = make([]*charge.Parameter, 0)
	err := copier.Copy(&params, parameter)
	if err != nil {
		return nil
	}
	return params
}

func changeDiscountsToGen(discount []*charge.Discount) []*gencharge.Discount {
	var discounts = make([]*gencharge.Discount, 0)
	err := copier.Copy(&discounts, discount)
	if err != nil {
		return nil
	}
	return discounts
}

func changeDiscountsFromGen(discount []*gencharge.Discount) []*charge.Discount {
	var discounts = make([]*charge.Discount, 0)
	err := copier.Copy(&discounts, discount)
	if err != nil {
		return nil
	}
	return discounts
}