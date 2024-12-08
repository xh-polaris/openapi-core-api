package adaptor

import (
	"context"
	"fmt"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

var rm *RightsManager

func CheckRole(ctx context.Context, role string) bool {
	if rm == nil {
		rm = &RightsManager{
			UserClient: openapi_user.NewOpenapiUser(config.GetConfig()),
		}
	}
	switch role {
	case consts.Admin:
		return rm.AdminOnly(ctx)
	default:
		return true
	}

}

type RightsManager struct {
	UserClient openapi_user.IOpenapiUser
}

func (m *RightsManager) AdminOnly(ctx context.Context) bool {
	userMeta := ExtractUserMeta(ctx)
	if userMeta == nil || userMeta.GetUserId() == "" {
		return false
	}
	resp, err := m.UserClient.GetUserInfo(ctx, &genuser.GetUserInfoReq{
		UserId: userMeta.UserId,
	})
	if err != nil || resp == nil || resp.Role != genuser.Role_ADMIN {
		return false
	}
	fmt.Printf(resp.String())
	return true
}
