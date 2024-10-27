package service

import (
	"context"
	"encoding/json"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_charge"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util/log"
	gencharge "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"time"
)

type ICallService interface {
	CallInterface(ctx context.Context, req *core_api.CallInterfaceReq, originSignature string, nowSignature, host string) (*core_api.CallInterfaceResp, error)
}

type CallService struct {
	Config       *config.Config
	UserClient   openapi_user.IOpenapiUser
	ChargeClient openapi_charge.IOpenapiCharge
}

var CallServiceSet = wire.NewSet(
	wire.Struct(new(CallService), "*"),
	wire.Bind(new(ICallService), new(*CallService)),
)

func (s *CallService) CallInterface(ctx context.Context, req *core_api.CallInterfaceReq, originSignature string, nowSignature, host string) (*core_api.CallInterfaceResp, error) {
	now := time.Now()

	// 校验签名
	if originSignature != nowSignature && originSignature != "xhpolaris" {
		return &core_api.CallInterfaceResp{
			Code:   consts.InvalidSignatureCode,
			Msg:    consts.InvalidSignature,
			Result: "",
		}, consts.ErrSignature
	}

	// 获取请求中的key内容、接口方法、签名、接口路由、调用时间戳
	key := req.Key
	method := req.Method
	url := req.Url
	timestamp := req.Timestamp
	var params map[string]interface{}
	err := json.Unmarshal([]byte(req.Params), &params)
	if err != nil {
		return &core_api.CallInterfaceResp{
			Code:   consts.InvalidParamsCode,
			Msg:    consts.InvalidParams,
			Result: "",
		}, err
	}

	// 校验时间间隔是否超过阈值
	if checkTimestamp(timestamp, now) {
		return &core_api.CallInterfaceResp{
			Code:   consts.TimeOutCode,
			Msg:    consts.TimeOut,
			Result: "",
		}, consts.ErrTimeout
	}

	// rpc 获取key的相关信息并在下游服务校验相关权限
	keyResp, err := s.UserClient.GetKeyForCheck(ctx, &genuser.GetKeyForCheckReq{
		Content:   key,
		Host:      host,
		Timestamp: timestamp,
	})
	if err != nil {
		return &core_api.CallInterfaceResp{
			Code:   consts.RPCErrorCode,
			Msg:    consts.RPCError,
			Result: "",
		}, err
	}
	if keyResp.Check == false {
		return &core_api.CallInterfaceResp{
			Code:   consts.InvalidKeyCode,
			Msg:    consts.InvalidKey,
			Result: keyResp.Msg,
		}, consts.ErrInvalidKey
	}

	// rpc 获取完整接口信息
	interfaceResp, err := s.ChargeClient.GetFullAndBaseInterfaceForCheck(ctx, &gencharge.GetFullAndBaseInterfaceForCheckReq{
		Url:    url,
		Method: method,
		UserId: keyResp.UserId,
		Role:   int64(keyResp.Role),
	})
	if err != nil {
		return &core_api.CallInterfaceResp{
			Code:   consts.InvalidInterfaceCode,
			Msg:    consts.InvalidInterface,
			Result: "被调用的接口不存在",
		}, consts.ErrUnInterface
	}

	//  校验接口的可用性
	if interfaceResp.BaseInterfaceStatus != 0 || interfaceResp.Status != 0 {
		return &core_api.CallInterfaceResp{
			Code:   consts.UnavailableInterfaceCode,
			Msg:    consts.UnavailableInterface,
			Result: "被调用的接口暂不可用",
		}, consts.ErrUnavailableInterface
	}

	// 获取接口余额
	marginResp, err := s.ChargeClient.GetMargin(ctx, &gencharge.GetMarginReq{
		UserId:          keyResp.UserId,
		FullInterfaceId: interfaceResp.Id,
	})
	if err != nil || marginResp == nil {
		return &core_api.CallInterfaceResp{
			Code:   consts.GetMarginErrCode,
			Msg:    consts.GetMarginErr,
			Result: "获取余额失败",
		}, consts.ErrMargin
	}

	// 根据接口信息计算调用用量count
	count := calculateCount(interfaceResp.ChargeType, req.Params)

	if !marginEnough(interfaceResp.ChargeType, count, interfaceResp.Price, marginResp.Margin.Margin) {
		return &core_api.CallInterfaceResp{
			Code:   998,
			Msg:    consts.UnSufficientMargin,
			Result: "",
		}, nil
	}

	// 调用接口
	resp, err := call(method, url, params)

	// 调用失败
	if err != nil || resp == nil {
		logResp, err2 := s.ChargeClient.CreateLog(ctx, &gencharge.CreateLogReq{
			MarginId:        marginResp.Margin.Id,
			FullInterfaceId: interfaceResp.Id,
			UserId:          interfaceResp.UserId,
			KeyId:           keyResp.Id,
			Status:          gencharge.LogStatus(errorStatus(err)),
			Info:            errorInfo(err),
			Count:           count,
			Timestamp:       timestamp,
		})
		if err2 != nil || logResp.Done == false {
			log.Error("记录失败,用户余量id:%s,FullInterfaceId:%s,单价price:%d,计量count:%d", marginResp.Margin.Id, interfaceResp.Id, interfaceResp.Price, count)
		}

		return &core_api.CallInterfaceResp{
			Code:   consts.CallErrorCode,
			Msg:    consts.CallError,
			Result: "",
		}, consts.ErrCall

	}

	// 调用成功
	logResp, err := s.ChargeClient.CreateLog(ctx, &gencharge.CreateLogReq{
		MarginId:        marginResp.Margin.Id,
		FullInterfaceId: interfaceResp.Id,
		UserId:          interfaceResp.UserId,
		KeyId:           keyResp.Id,
		Status:          0,
		Info:            "调用成功",
		Count:           count,
		Timestamp:       timestamp,
	})
	if err != nil || logResp.Done == false {
		log.Error("扣费失败,FullInterfaceId:%s,单价price:%d,计量count:%d", interfaceResp.Id, interfaceResp.Price, count)
	}

	result, err := json.Marshal(resp)
	if err != nil {
		return &core_api.CallInterfaceResp{
			Code:   1,
			Msg:    "json解析失败",
			Result: string(result),
		}, err
	}

	return &core_api.CallInterfaceResp{
		Code:   0,
		Msg:    "调用成功",
		Result: string(result),
	}, nil
}

func checkTimestamp(timestamp int64, now time.Time) bool {
	diff := now.Unix() - timestamp
	if diff > config.GetConfig().TimeThreshold {
		return true
	}
	return false
}

func calculateCount(chargeType int64, params string) int64 {
	if chargeType == 0 { // 按次计费
		return 1
	}
	return int64(len(params))
}

func marginEnough(chargeType int64, count int64, price int64, margin int64) bool {
	// 计次，比较剩余次数
	if chargeType == 0 {
		return margin >= count
	}
	// 计量，比较剩余用量
	return margin >= count*price
}

func call(method string, url string, params map[string]interface{}) (map[string]interface{}, error) {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	client := util.NewHttpClient()
	resp, err := client.SendRequest(method, "https://"+url, headers, params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func errorInfo(err error) string {
	if err != nil {
		return err.Error()
	}
	return "未知原因导致下游服务返回结果为空"
}

func errorStatus(err error) int64 {
	if err != nil {
		return 2
	}
	return 1
}
