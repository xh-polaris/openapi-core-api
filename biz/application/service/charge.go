package service

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/charge"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/mq"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_charge"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	gencharge "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	BuyFullInterface(ctx context.Context, req *core_api.BuyFullInterfaceReq) (*core_api.Response, error)
	GetFullInterfaces(ctx context.Context, req *core_api.GetFullInterfacesReq) (*core_api.GetFullInterfacesResp, error)
	CreateGradient(ctx context.Context, req *core_api.CreateGradientReq) (*core_api.Response, error)
	UpdateGradient(ctx context.Context, req *core_api.UpdateGradientReq) (*core_api.Response, error)
	GetGradient(ctx context.Context, req *core_api.GetGradientReq) (*core_api.GetGradientResp, error)
}

type ChargeService struct {
	ChargeClient openapi_charge.IOpenapiCharge
	UserClient   openapi_user.IOpenapiUser
	Producer     *mq.Producer
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

func (s *ChargeService) BuyFullInterface(ctx context.Context, req *core_api.BuyFullInterfaceReq) (*core_api.Response, error) {
	// 获取用户信息
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	userId := userMeta.GetUserId()
	increment := req.Increment
	infId := req.FullInterfaceId
	isDiscount := req.Discount

	// 根据id获取完整接口
	getResp, getErr := s.ChargeClient.GetOneFullInterface(ctx, &gencharge.GetOneFullInterfaceReq{
		Id: infId,
	})
	if getErr != nil || getResp == nil || getResp.Inf == nil {
		return util.FailResponse(nil, "未获取到模板，购买失败，请重试"), getErr
	}
	fmt.Printf("getResp  " + getResp.String())
	fullInf := getResp.Inf
	fullInfId := fullInf.Id

	var marginId string

	// 根据完整接口和用户id获取接口余量信息
	marginResp, marginErr := s.ChargeClient.GetMargin(ctx, &gencharge.GetMarginReq{
		UserId:          userId,
		FullInterfaceId: fullInfId,
	})

	// 之前没有购买过，则创建用户的接口余量
	if marginErr != nil || marginResp == nil || marginResp.Margin == nil {
		createMarginResp, createMarginErr := s.ChargeClient.CreateMargin(ctx, &gencharge.CreateMarginReq{
			UserId:          userId,
			FullInterfaceId: fullInfId,
			Margin:          0,
		})
		if createMarginErr != nil || createMarginResp == nil {
			return util.FailResponse(createMarginResp, "创建接口余量失败，购买失败，请重试"), createMarginErr
		}
		fmt.Printf("createMarginResp: %+v\n", createMarginResp.String())
		marginId = createMarginResp.MarginId
	} else {
		fmt.Println(marginResp.String())
		marginId = marginResp.Margin.Id
	}

	// 计算总额
	var amount int64
	var rate int64
	amount = 0
	rate = 100
	amount = increment * fullInf.Price

	// 判断是否折扣
	if isDiscount {
		// 获取梯度折扣
		gradientsResp, gradientErr := s.ChargeClient.GetGradient(ctx, &gencharge.GetGradientReq{
			BaseInterfaceId: fullInf.BaseInterfaceId,
		})
		// 未获取到梯度折扣
		if gradientErr != nil || gradientsResp == nil || gradientsResp.Gradient == nil {
			return util.FailResponse(nil, "获取折扣失败，请重新购买"), gradientErr
		}
		// 选取折扣
		gradients := gradientsResp.Gradient

		// 判断折扣是否可用
		if gradients.Status != 0 {
			return util.FailResponse(nil, "折扣暂不可用，购买失败"), nil
		}

		for _, discount := range gradients.Discounts {
			if increment > discount.Low {
				rate = discount.Rate
			}
		}
		amount = amount * rate / 100
	}

	txId := primitive.NewObjectID().Hex() // 事务id

	// 给消息队列发送对账消息
	err := s.Producer.SendBuyMsg(txId, -1*amount, rate, increment, fullInf.Price)
	if err != nil {
		return util.FailResponse(nil, "消息发送失败"), err
	}

	// 扣除用户余额
	remainResp, err := s.UserClient.SetRemain(ctx, &genuser.SetRemainReq{
		UserId:    userId,
		Increment: -1 * amount,
		TxId:      &txId,
	})
	if err != nil || remainResp == nil {
		return util.FailResponse(remainResp, "余额扣除失败，请重新购买"), err
	}
	fmt.Printf("remainResp: %+v\n", remainResp.String())

	// 增加接口余量
	updateMarginResp, updateMarginErr := s.ChargeClient.UpdateMargin(ctx, &gencharge.UpdateMarginReq{
		Id:        marginId,
		Increment: increment,
		TxId:      &txId,
	})
	if updateMarginErr != nil || !updateMarginResp.Done {
		return util.FailResponse(updateMarginResp, "接口余量增加失败"), err
	}

	return &core_api.Response{
		Done: true,
		Msg:  "购买接口成功",
	}, nil
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
