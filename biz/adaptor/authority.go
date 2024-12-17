package adaptor

import (
	"context"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util/log"
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
		info := userMeta.UserId + "尝试进行管理员操作失败:"
		if resp != nil {
			info += resp.String()
		}
		log.Info(info)

		return false
	}
	return true
}
