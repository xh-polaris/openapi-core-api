package test

import (
	"fmt"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/mq"
	"testing"
)

func TestMQ(t *testing.T) {
	con := &config.Config{
		RocketMQ: config.RocketMQ{
			NameServers: []string{"namesrv.rocketmq"},
		},
	}
	pro := mq.NewProducer(con)
	err2 := pro.SendBuyMsg("test", 1, 1, 1, 1)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
