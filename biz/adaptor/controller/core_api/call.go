// Code generated by hertz generator.

package core_api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xh-polaris/openapi-core-api/biz/adaptor"
	core_api "github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"github.com/xh-polaris/openapi-core-api/provider"
)

// CallInterface .
// @router /call/ [POST]
func CallInterface(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CallInterfaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CallService.CallInterface(ctx, &req, string(c.GetHeader("signature")), getSignature(c), getHostOrIP(c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

func getHostOrIP(c *app.RequestContext) string {
	// 暂时先从origin中获取host，与key的host进行校验
	// 之后可以根据需求转换成真实ip校验
	return string(c.GetHeader("origin"))
}

func getSignature(c *app.RequestContext) string {
	// 获取请求的 method、url、header 和 body
	method := c.Request.Method()
	url := c.Request.Path() // 获取请求的路径
	headers := c.Request.Header.Header()
	body := c.Request.Body()

	data := make([]byte, 0)
	data = append(data, method...)
	data = append(data, url...)
	data = append(data, headers...)
	data = append(data, body...)

	// 计算 SHA256 签名
	hash := sha256.New()
	hash.Write(data)
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature
}
