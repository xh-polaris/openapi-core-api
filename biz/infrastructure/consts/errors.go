package consts

import "errors"

var ErrNotAuthentication = errors.New("not authentication")
var ErrTimeout = errors.New("calling time exceeded the threshold")
var ErrInvalidKey = errors.New("invalid key")
var ErrUnInterface = errors.New("invalid interface")
var ErrUnavailableInterface = errors.New("unavailable interface")
var ErrCall = errors.New("call failed")
var ErrSignature = errors.New("signature invalid")
var ErrRole = errors.New("没有权限")
