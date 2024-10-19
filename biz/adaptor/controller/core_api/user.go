// Code generated by hertz generator.

package core_api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	core_api "github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
)

// SignUp .
// @router /user/sign_up [POST]
func SignUp(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SignUpReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.SignUpResp)

	c.JSON(consts.StatusOK, resp)
}
