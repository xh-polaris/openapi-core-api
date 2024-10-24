package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IUserService interface {
	SignUp(ctx context.Context, req *core_api.SignUpReq) (*core_api.SignUpResp, error)
}

type UserService struct {
	Config     *config.Config
	UserClient openapi_user.IOpenapiUser
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

// SignUp 用于注册一个用户，之后的登录只需要用中台获取token即可
func (s *UserService) SignUp(ctx context.Context, req *core_api.SignUpReq) (*core_api.SignUpResp, error) {
	// 在中台注册账户
	httpClient := util.NewHttpClient()
	httpResp, err := httpClient.SignUp(req.AuthType, req.AuthId, *req.VerifyCode)
	if err != nil {
		return nil, err
	}
	// 调用rpc初始化用户
	userId := httpResp["userId"].(string)
	resp, err := s.UserClient.SignUp(ctx, &genuser.SignUpReq{
		UserId: userId,
		Role:   genuser.Role(req.Role),
	})
	if err != nil {
		return nil, err
	}
	// 初始化失败
	if resp.Done == false {
		return nil, errors.New(resp.Msg)
	}
	// 正常初始化
	return &core_api.SignUpResp{
		UserId:       userId,
		AccessToken:  httpResp["accessToken"].(string),
		AccessExpire: httpResp["accessExpire"].(int64),
	}, nil
	// TODO 如果中台创建用户成功但是下游服务初始化用户失败，怎么办。给登录也包一层，来补偿这种下游服务注册失败的情况
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error) {
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}

	info, err := s.UserClient.GetUserInfo(ctx, &genuser.GetUserInfoReq{
		UserId: userMeta.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	return &core_api.GetUserInfoResp{
		Username: info.Username,
		Role:     user.Role(info.Role),
		Auth:     info.Auth,
		AuthId:   info.AuthId,
		Status:   user.UserStatus(info.Status),
	}, nil
}

func (s *UserService) SetUserInfo(ctx context.Context, req *core_api.SetUserInfoReq) (*core_api.Response, error) {
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}

	resp, err := s.UserClient.SetUserInfo(ctx, &genuser.SetUserInfoReq{
		UserId:   userMeta.GetUserId(),
		Username: req.Username,
		Status:   (*genuser.UserStatus)(req.Status),
	})
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}

	return util.SuccessResponse(resp)
}

func (s *UserService) SignIn(ctx context.Context, req *core_api.SignUpReq) (*core_api.SignInResp, error) {
	// 在中台登录账户
	httpClient := util.NewHttpClient()
	httpResp, err := httpClient.SignIn(req.AuthType, req.AuthId, *req.VerifyCode, *req.Password)
	if err != nil {
		return nil, err
	}
	userId := httpResp["userId"].(string)

	// 校验用户初始化是否成功
	info, err := s.UserClient.GetUserInfo(ctx, &genuser.GetUserInfoReq{
		UserId: userId,
	})
	// 获取失败
	if err != nil {
		return nil, err
	}
	// 先前用户初始化失败，补偿初始化
	if info == nil || info.UserId == "" {
		resp, err2 := s.UserClient.SignUp(ctx, &genuser.SignUpReq{
			UserId: userId,
			Role:   genuser.Role(req.Role),
		})
		if err2 != nil {
			return nil, err
		}
		// 初始化失败
		if resp.Done == false {
			return nil, errors.New(resp.Msg)
		}
	}

	// 返回登录信息
	return &core_api.SignInResp{
		UserId:       userId,
		AccessToken:  httpResp["accessToken"].(string),
		AccessExpire: httpResp["accessExpire"].(int64),
	}, nil
}
