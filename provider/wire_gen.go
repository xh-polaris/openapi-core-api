// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xh-polaris/openapi-core-api/biz/application/service"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_charge"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
)

// Injectors from wire.go:

func NewProvider() (*Provider, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	client := openapi_user.NewOpenapiUser(configConfig)
	openapiUser := &openapi_user.OpenapiUser{
		Client: client,
	}
	chargeClient := openapi_charge.NewOpenapiCharge(configConfig)
	openapiCharge := &openapi_charge.OpenapiCharge{
		Client: chargeClient,
	}
	callService := service.CallService{
		Config:       configConfig,
		UserClient:   openapiUser,
		ChargeClient: openapiCharge,
	}
	chargeService := service.ChargeService{
		ChargeClient: openapiCharge,
		UserClient:   openapiUser,
	}
	keyService := service.KeyService{
		KeyClient: openapiUser,
	}
	userService := service.UserService{
		Config:     configConfig,
		UserClient: openapiUser,
	}
	providerProvider := &Provider{
		Config:        configConfig,
		CallService:   callService,
		ChargeService: chargeService,
		KeyService:    keyService,
		UserService:   userService,
	}
	return providerProvider, nil
}
