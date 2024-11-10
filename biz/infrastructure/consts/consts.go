package consts

// http
const (
	Get             = "GET"
	Post            = "POST"
	PlatformUrl     = "https://api.xhpolaris.com/platform/auth/sign_in"
	ContentTypeJson = "application/json"
	CharSetUTF8     = "UTF-8"
)

// msg & code
const (
	GetMarginErrCode         = 997
	GetMarginErr             = "获取余额失败"
	UnSufficientMarginCode   = 998
	UnSufficientMargin       = "余量不足"
	CallErrorCode            = 999
	CallError                = "调用失败"
	RPCErrorCode             = 1000
	RPCError                 = "RPC调用错误"
	InvalidParamsCode        = 1001
	InvalidParams            = "参数格式错误"
	TimeOutCode              = 1002
	TimeOut                  = "调用时间超时"
	InvalidKeyCode           = 1003
	InvalidKey               = "密钥校验失败"
	InvalidInterfaceCode     = 1004
	InvalidInterface         = "被调用的接口不存在"
	UnavailableInterfaceCode = 1005
	UnavailableInterface     = "被调用的接口暂不可用"
	InvalidSignatureCode     = 1006
	InvalidSignature         = "签名不匹配"
)

// role
const (
	Admin = "admin"
	User  = "user"
)

var FilteredHeaders = map[string]bool{
	"X-Real-Ip":         true,
	"X-Forwarded-For":   true,
	"X-Forwarded-Proto": true,
	"X-Forwarded-Host":  true,
	"X-Forwarded-Port":  true,
	"Content-Length":    true,
	"X-Xh-Env":          true,
}

// mq
const (
	BuyTopic = "openapi_buy"
)
