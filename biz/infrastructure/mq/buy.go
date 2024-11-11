package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	mqprimitive "github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/hertz/pkg/common/json"
	logx "github.com/xh-polaris/gopkg/util/log"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/consts"
)

type IProducer interface {
	SendBuyMsg(txId string, amount int64, rate int64, increment int64, price int64) error
}

type Producer struct {
	Config *config.Config
}

func NewProducer(config *config.Config) *Producer {
	return &Producer{Config: config}
}

func (pro *Producer) SendBuyMsg(txId string, amount int64, rate int64, increment int64, price int64) error {
	// 创建一个Producer实例
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(pro.Config.RocketMQ.NameServers),
		producer.WithRetry(2),
	)
	if err != nil {
		return err
	}

	err = p.Start()
	if err != nil {
		return err
	}
	defer func() {
		err2 := p.Shutdown()
		if err2 != nil {
			logx.Error("producer关闭失败")
		}
	}()

	// 定义消息内容
	var msgBody map[string]interface{}
	msgBody["txId"] = txId           // 事务id
	msgBody["rate"] = rate           // 折扣
	msgBody["increment"] = increment // 接口余量增加量
	msgBody["amount"] = amount       // 用户余额扣减
	msgBody["price"] = price         // 购买时单价

	msgBodyBytes, err := json.Marshal(msgBody)
	if err != nil {
		return err
	}

	msg := &mqprimitive.Message{
		Topic: consts.BuyTopic,
		Body:  msgBodyBytes,
	}

	_, err = p.SendSync(context.Background(), msg)
	if err != nil {
		return err
	}

	return nil
}
