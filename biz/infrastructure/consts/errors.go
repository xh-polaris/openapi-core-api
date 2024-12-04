package consts

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Errno struct {
	err  error
	code codes.Code
}

// GRPCStatus 实现 GRPCStatus 方法
func (en *Errno) GRPCStatus() *status.Status {
	return status.New(en.code, en.err.Error())
}

// 实现 Error 方法
func (en *Errno) Error() string {
	return en.err.Error()
}

// NewErrno 创建自定义错误
func NewErrno(code codes.Code, err error) *Errno {
	return &Errno{
		err:  err,
		code: code,
	}
}

// 定义常量错误
var (
	ErrNotAuthentication    = NewErrno(codes.Unauthenticated, errors.New("not authentication"))
	ErrTimeout              = NewErrno(codes.DeadlineExceeded, errors.New("calling time exceeded the threshold"))
	ErrInvalidKey           = NewErrno(codes.InvalidArgument, errors.New("invalid key"))
	ErrUnInterface          = NewErrno(codes.InvalidArgument, errors.New("invalid interface"))
	ErrUnavailableInterface = NewErrno(codes.Unavailable, errors.New("unavailable interface"))
	ErrCall                 = NewErrno(codes.Code(1001), errors.New("call failed"))
	ErrSignature            = NewErrno(codes.InvalidArgument, errors.New("signature invalid"))
	ErrRole                 = NewErrno(codes.PermissionDenied, errors.New("没有权限"))
	ErrMargin               = NewErrno(codes.Code(1002), errors.New("获取余额失败,请重试"))
	ErrGradient             = NewErrno(codes.Code(1003), errors.New("获取价格失败，请重试"))
)
