package openapi_user

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user/user"
)

type IOpenapiUser interface {
	user.Client
}

type OpenapiUser struct {
	user.Client
}

var OpenapiUserSet = wire.NewSet(
	NewOpenapiUser,
	wire.Struct(new(OpenapiUser), "*"),
	wire.Bind(new(IOpenapiUser), new(*OpenapiUser)),
)

func NewOpenapiUser(config *config.Config) user.Client {
	return client.NewClient(config.Name, "openapi.user", user.NewClient)
}
