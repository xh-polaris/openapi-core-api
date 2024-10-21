package util

import (
	"github.com/xh-polaris/openapi-core-api/biz/application/dto/openapi/core_api"
	"strconv"
	"sync"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/cloudwego/hertz/pkg/common/json"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

func JSONF(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Error("JSONF fail, v=%v, err=%v", v, err)
	}
	return string(data)
}

func ParseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func ParallelRun(fns ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(fns))
	for _, fn := range fns {
		fn := fn
		gopool.Go(func() {
			defer wg.Done()
			fn()
		})
	}
	wg.Wait()
}

func GetMsg(resp interface{ GetMsg() string }, msg string) string {
	if resp == nil {
		return msg
	}
	return resp.GetMsg()
}

func FailResponse(resp interface{ GetMsg() string }, msg string) *core_api.Response {
	return &core_api.Response{
		Done: false,
		Msg:  GetMsg(resp, msg),
	}
}

func SuccessResponse(resp interface {
	GetMsg() string
	GetDone() bool
}) (*core_api.Response, error) {
	return &core_api.Response{
		Done: resp.GetDone(),
		Msg:  resp.GetMsg(),
	}, nil
}
