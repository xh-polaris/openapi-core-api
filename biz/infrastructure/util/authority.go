package util

import (
	"context"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

var rm *RightsManager

func CheckRole(ctx context.Context, role string) bool {
	if rm == nil {
		rm = &RightsManager{}
	}
	switch role {
	case "admin":
		return rm.AdminOnly(ctx)
	default:
		return true
	}

}

type RightsManager struct {
	UserClient openapi_user.IOpenapiUser
}

func (m *RightsManager) AdminOnly(ctx context.Context) bool {
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return false
	}
	resp, err := m.UserClient.GetUserInfo(ctx, &genuser.GetUserInfoReq{
		UserId: userMeta.UserId,
	})
	if err != nil || resp.Role != genuser.Role_ADMIN {
		return false
	}
	return true
}
