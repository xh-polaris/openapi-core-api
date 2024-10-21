package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/basic"
	core_api "github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/rpc/openapi_user"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IKeyService interface {
	GenerateKey(ctx context.Context, req *core_api.GenerateKeyReq) (*core_api.Response, error)
	GetKeys(ctx context.Context, req *core_api.GetKeysReq) (*core_api.GetKeysResp, error)
	UpdateKey(ctx context.Context, req *core_api.UpdateKeyReq) (*core_api.Response, error)
	RefreshKey(ctx context.Context, req *core_api.RefreshKeyReq) (*core_api.Response, error)
	UpdateHosts(ctx context.Context, req *core_api.UpdateHostReq) (*core_api.Response, error)
	DeleteKey(ctx context.Context, req *core_api.DeleteKeyReq) (*core_api.Response, error)
}

type KeyService struct {
	KeyClient openapi_user.IOpenapiUser
}

var KeyServiceSet = wire.NewSet(
	wire.Struct(new(KeyService), "*"),
	wire.Bind(new(IKeyService), new(*KeyService)),
)

func (s *KeyService) GenerateKey(ctx context.Context, req *core_api.GenerateKeyReq) (*core_api.Response, error) {
	// 提取用户信息
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	// rpc创建密钥
	resp, err := s.KeyClient.CreateKey(ctx, &genuser.CreateKeyReq{
		UserId: userMeta.UserId,
		Name:   req.Name,
		Hosts:  req.Hosts,
	})
	// 若创建失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "创建失败"), err
	}
	// 若创建成功
	return util.SuccessResponse(resp)

}

func (s *KeyService) GetKeys(ctx context.Context, req *core_api.GetKeysReq) (*core_api.GetKeysResp, error) {
	// 获取用户请求
	userMeta := adaptor.ExtractUserMeta(ctx)
	if userMeta.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	// paginationOptions 类型转换
	if req.PaginationOptions == nil {
		req.PaginationOptions = &basic.PaginationOptions{}
	}
	// rpc获取keys
	resp, err := s.KeyClient.GetKey(ctx, &genuser.GetKeysReq{
		UserId: userMeta.UserId,
		PaginationOptions: &genbasic.PaginationOptions{
			Page:      req.PaginationOptions.Page,
			Limit:     req.PaginationOptions.Limit,
			LastToken: req.PaginationOptions.LastToken,
			Backward:  req.PaginationOptions.Backward,
			Offset:    req.PaginationOptions.Offset,
		},
	})
	if err != nil {
		return nil, err
	}
	// keys类型转换
	var keys = make([]*user.Key, 0)
	err = copier.Copy(&keys, resp.Keys)
	if err != nil {
		return nil, err
	}
	// 处理status类型不同无法复制
	for i := 0; i < len(keys); i++ {
		keys[i].Status = user.KeyStatus(resp.Keys[i].Status)
	}

	return &core_api.GetKeysResp{
		Keys:  keys,
		Total: resp.Total,
	}, nil
}

func (s *KeyService) UpdateKey(ctx context.Context, req *core_api.UpdateKeyReq) (*core_api.Response, error) {
	// rpc更新密钥
	resp, err := s.KeyClient.UpdateKey(ctx, &genuser.UpdateKeyReq{
		Id:         req.Id,
		Name:       req.Name,
		Status:     (*genuser.KeyStatus)(req.Status),
		Timestamp:  req.Timestamp,
		ExpireTime: req.ExpireTime,
	})
	// 若更新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}
	// 若更新成功
	return util.SuccessResponse(resp)
}

func (s *KeyService) RefreshKey(ctx context.Context, req *core_api.RefreshKeyReq) (*core_api.Response, error) {
	// rpc刷新
	resp, err := s.KeyClient.RefreshKey(ctx, &genuser.RefreshKeyReq{
		Id: req.Id,
	})
	// 刷新失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "刷新失败"), err
	}
	// 刷新成功
	return util.SuccessResponse(resp)
}

func (s *KeyService) UpdateHosts(ctx context.Context, req *core_api.UpdateHostReq) (*core_api.Response, error) {
	// rpc更新
	resp, err := s.KeyClient.UpdateHosts(ctx, &genuser.UpdateHostsReq{
		Id:    req.Id,
		Hosts: req.Hosts,
	})
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "更新失败"), err
	}
	return util.SuccessResponse(resp)
}

func (s *KeyService) DeleteKey(ctx context.Context, req *core_api.DeleteKeyReq) (*core_api.Response, error) {
	// rpc删除
	resp, err := s.KeyClient.DeleteKey(ctx, &genuser.DeleteKeyReq{
		Id: req.Id,
	})
	// 若删除失败
	if err != nil || resp.Done == false {
		return util.FailResponse(resp, "删除失败"), err
	}
	// 若删除成功
	return util.SuccessResponse(resp)
}
