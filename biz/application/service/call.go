package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
)

type ICallService interface {
	CallInterface(ctx context.Context, c *core_api.CallInterfaceReq) (*core_api.CallInterfaceResp, error)
}

type CallService struct{}

var CallServiceSet = wire.NewSet(
	wire.Struct(new(CallService), "*"),
	wire.Bind(new(ICallService), new(*CallService)),
)

func (s *CallService) CallInterface(ctx context.Context, c *core_api.CallInterfaceReq) (*core_api.CallInterfaceResp, error) {

}
