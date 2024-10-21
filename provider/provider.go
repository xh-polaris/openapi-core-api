package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-core-api/biz/application/service"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_charge"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config        *config.Config
	CallService   service.CallService
	ChargeService service.ChargeService
	KeyService    service.KeyService
	UserService   service.UserService
}

func Get() *Provider {
	return provider
}

var RPCSet = wire.NewSet(
	openapi_charge.OpenapiChargeSet,
	openapi_user.OpenapiUserSet)

var ApplicationSet = wire.NewSet(
	service.CallServiceSet,
	service.ChargeServiceSet,
	service.KeyServiceSet,
	service.UserServiceSet,
)

var DomainSet = wire.NewSet()

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	DomainSet,
	InfrastructureSet)
