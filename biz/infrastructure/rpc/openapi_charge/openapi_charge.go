package openapi_charge

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge/charge"
)

type IOpenapiCharge interface {
	charge.Client
}

type OpenapiCharge struct {
	charge.Client
}

var OpenapiChargeSet = wire.NewSet(
	NewOpenapiCharge,
	wire.Struct(new(OpenapiCharge), "*"),
	wire.Bind(new(IOpenapiCharge), new(*OpenapiCharge)),
)

func NewOpenapiCharge(config *config.Config) charge.Client {
	return client.NewClient(config.Name, "openapi.charge", charge.NewClient)
}
