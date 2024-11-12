package mq

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/bytedance/sonic"
	logx "github.com/xh-polaris/gopkg/util/log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-core-api/biz/infrastructure/util"
)

type IProducer interface {
	SendBuyMsg(txId string, amount int64, rate int64, increment int64, price int64) error
}

type Producer struct {
	conn *amqp.Connection
}

type Consumer struct {
	config *config.Config
	conn   *amqp.Connection
	finish chan struct{}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewProducer(config *config.Config) (*Producer, error) {
	conn, err := newConnection(config)
	if err != nil {
		return nil, err
	}
	return &Producer{
		conn: conn,
	}, nil
}

func NewConsumer(config *config.Config) *Consumer {
	conn, err := newConnection(config)
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	return &Consumer{
		config: config,
		conn:   conn,
		finish: make(chan struct{}),
	}
}

func newConnection(config *config.Config) (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMqConf.URL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (p *Producer) SendBuyMsg(ctx context.Context, txId string, amount int64, rate int64, increment int64, price int64) error {
	ch, err := p.conn.Channel()
	if err != nil {
		logx.CtxError(ctx, "[SendBuyMsg] open channel err:%v", err)
		return err
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"buy_msg", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logx.CtxError(ctx, "[SendBuyMsg] declare exchange err:%v", err)
		return err
	}

	// 定义消息内容
	var msgBody map[string]interface{}
	msgBody = make(map[string]interface{})
	msgBody["txId"] = txId           // 事务id
	msgBody["rate"] = rate           // 折扣
	msgBody["increment"] = increment // 接口余量增加量
	msgBody["amount"] = amount       // 用户余额扣减
	msgBody["price"] = price         // 购买时单价

	msgBodyBytes, err := sonic.Marshal(msgBody)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"buy_msg", // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: 2, // 持久化
			ContentType:  "application/json",
			Body:         msgBodyBytes,
		})
	return nil
}

func (c *Consumer) Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	c.subscribe(ctx, c.buyMsgHandler)
	c.subscribe(ctx, c.osSignalHandler)
	<-c.finish
	logx.CtxInfo(ctx, "finish")
}

func (c *Consumer) subscribe(ctx context.Context, fn func(context.Context)) {
	gopool.CtxGo(ctx, func() {
		fn(ctx)
		c.finish <- struct{}{}
	})
}

func (c *Consumer) osSignalHandler(ctx context.Context) {
	logx.CtxInfo(ctx, "[osSignalHandler] start")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	osSignal := <-ch
	logx.CtxInfo(ctx, "[osSignalHandler] receive signal:[%v]", osSignal)
}

func (c *Consumer) buyMsgHandler(ctx context.Context) {
	ch, err := c.conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		c.config.Name,       // name
		c.config.Name != "", // durable
		c.config.Name == "", // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
	}
	err = ch.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		"buy_msg", // exchange
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Failed to bind a queue")
	}
	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name, // queue
		fmt.Sprintf("%s@%s", os.Getenv("MY_POD_NAME"), os.Getenv("MY_POD_IP")), // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	for msg := range msgs {
		logx.CtxInfo(ctx, "[buyMsgHandler] receive msg:[%s]", util.JSONF(string(msg.Body)))
		msg.Ack(false)
	}
}
