package test

import (
	"context"
	"testing"

	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/mq"
)

func TestMQ(t *testing.T) {
	con := &config.Config{
		RabbitMqConf: config.RabbitMqConf{URL: "amqp://user:DEHjiFaLtrM6cJ5d@rabbitmq.rabbitmq:5672/"},
	}
	pro, err := mq.NewProducer(con)
	if err != nil {
		t.Error(err)
		return
	}
	err = pro.SendBuyMsg(context.Background(), "test", 1, 1, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestConsumer(t *testing.T) {
	con := &config.Config{
		RabbitMqConf: config.RabbitMqConf{URL: "amqp://user:DEHjiFaLtrM6cJ5d@rabbitmq.rabbitmq:5672/"},
	}
	con.Name = "openapi.core-api"
	c := mq.NewConsumer(con)
	c.Start()
}
